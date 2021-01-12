package doc

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

func saveSaleBill(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	var (
		ctx = c.Request().Context()
		req = request.(*portalModel.CreateSaleBillReq)
	)

	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "create Sale Bill request",
	}

	// parse data to []uint64 for req
	var (
		barcodeIDs    = make([]uint64, len(req.SaleBarcode))
		price         = make([]uint64, len(req.SaleBarcode))
		barcodeStatus = make([]uint64, len(req.SaleBarcode))
	)

	for i := range req.SaleBarcode {
		barcodeIDs[i] = req.SaleBarcode[i].BarcodeID
		price[i] = req.SaleBarcode[i].Price
		barcodeStatus[i] = model.BarcodeSoldStatus
	}

	// gen Sale Bill ID
	saleBillUUID, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	// rpc
	rpcCreateSaleBillReq := &docmanagerModel.SaveSaleBillReq{
		SaleBill: &docmanagerModel.SaleBill{
			ID:          uint64(core.GetHash(saleBillUUID.String())),
			LibrarianID: req.LibrarianID,
			BarcodeID:   barcodeIDs,
			Price:       price,
		},
	}

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}
	shardID := core.GetShardID(uint32(rpcCreateSaleBillReq.SaleBill.ID))
	ser, ok := shardService[uint64(shardID)]
	if !ok {
		fmt.Println("nil shardID")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return
	}

	resp, er := ser.Docmanager.SaveSaleBill(ctx, rpcCreateSaleBillReq)
	if er != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	if er := updateBarcodeStatusByBatch(ctx, barcodeIDs, barcodeStatus, rpcCreateSaleBillReq.SaleBill.ID); er != nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, nil, lg, false, nil
}
