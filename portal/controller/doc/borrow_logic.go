package doc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pinezapple/LibraryProject20201/portal/dao/cache"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

var (
	hours24 = time.Hour * 24
)

func createBorrowForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*portalModel.CreateBorrowFormReq)

	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "create borrow form request",
	}

	// gen borrowform ID
	borrowformUUID, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	// create borrowform to save
	rpcBorrowFormReq := docmanagerModel.SaveBorrowFormReq{
		Borrowform: &docmanagerModel.BorrowForm{
			ID:          uint64(core.GetHash(borrowformUUID.String())),
			LibrarianID: req.LibrarianID,
			ReaderID:    req.ReaderID,
			BarcodeID:   req.Barcodes,
			StartTime:   model.NewTime(time.Now()),
			EndTime:     model.NewTime(time.Now().Add(hours24 * time.Duration(req.BorrowDays))),
		},
	}

	// get borrow form shard
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}
	shardID := core.GetShardID(uint32(rpcBorrowFormReq.Borrowform.ID))
	ser, ok := shardService[uint64(shardID)]
	if !ok {
		fmt.Println("nil shardID")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return
	}

	resp, er := ser.Docmanager.SaveBorrowForm(ctx, &rpcBorrowFormReq)
	if er != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	// update barcode status to borrowing
	barcodeUpdateBorrowStatus := make([]uint64, len(req.Barcodes))
	for i := range barcodeUpdateBorrowStatus {
		barcodeUpdateBorrowStatus[i] = model.BarcodeBorrowingStatus
	}
	if er := updateBarcodeStatusByBatch(ctx, req.Barcodes, barcodeUpdateBorrowStatus); er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

func updateBorrowForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	var (
		ctx = c.Request().Context()
		req = request.(*portalModel.UpdateBorrowFormReq)
	)

	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "update borrow form",
	}

	if req.Fine != 0 {
		req.Status = 3
	} else {
		req.Status = 0
	}

	// create request
	rpcUpdateBorrowFormReq := &docmanagerModel.UpdateBorrowFormStatusReq{
		BorrowFormID: req.BorrowFormID,
		Status:       req.Status,
	}

	// get shard by borrowform id
	ser, er := getDocMangerServiceByUint64(rpcUpdateBorrowFormReq.BorrowFormID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	resp, er := ser.Docmanager.UpdateBorrowFormStatus(ctx, rpcUpdateBorrowFormReq)
	if er != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error: %v", er)
		return
	}

	// create batch to save barcode
	var (
		barcodeID            = make([]uint64, len(req.BarcodeUpdate))
		barcodeStatus        = make([]uint64, len(req.BarcodeUpdate))
		paymentBarcodeStatus []uint64
		barcodeFee           []uint64
		totalFee             uint64
	)

	for i := range req.BarcodeUpdate {
		barcodeID[i] = req.BarcodeUpdate[i].BarcodeID
		barcodeStatus[i] = req.BarcodeUpdate[i].BarcodeStatus
		if req.BarcodeUpdate[i].Fee == 0 && req.BarcodeUpdate[i].BarcodeStatus != model.BarcodeNormalStatus {
			continue
		}
		paymentBarcodeStatus = append(paymentBarcodeStatus, req.BarcodeUpdate[i].BarcodeStatus)
		barcodeFee = append(barcodeFee, req.BarcodeUpdate[i].Fee)
		totalFee += req.BarcodeUpdate[i].Fee
	}

	if er := updateBarcodeStatusByBatch(ctx, barcodeID, barcodeStatus); er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	// create payment if necessary
	if len(paymentBarcodeStatus) > 0 {
		if er := createPayment(ctx, req.BorrowFormID, req.LibrarianID, req.ReaderID, barcodeID, paymentBarcodeStatus, barcodeFee, req.Fine); er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
	}

	// insert to blacklist
	rpcSelectBorrowFormReq := &docmanagerModel.SelectBorrowFormByIDReq{
		BorrowFormID: req.BorrowFormID,
	}
	rpcSelectBorrowFormResp, er := ser.Docmanager.SelectBorrowFormByID(ctx, rpcSelectBorrowFormReq)
	if er != nil || rpcSelectBorrowFormResp.Code != 0 || rpcSelectBorrowFormResp.Borrowform == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error: %v", er)
		return
	}

	if er := cache.InsertIntoBlackList(ctx, core.GetDB(), rpcSelectBorrowFormResp.Borrowform.ReaderID, req.BorrowFormID, totalFee); er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, nil, lg, false, nil
}
