package doc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

var (
	ErrUpdateBarcodeLengthNotMatch = errors.New("len barcodes != len barCodestatus")
)

func updateBarcodeStatus(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	var (
		ctx     = c.Request().Context()
		httpReq = request.(*portalModel.UpdateBarcodeStatus)
	)
	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "update Barcode Status",
	}

	//rpc
	rpcUpdateBarcodeStatusReq := &docmanagerModel.UpdateBarcodeReq{
		Barcode: &docmanagerModel.Barcode{
			ID:     httpReq.BarcodeID,
			Status: httpReq.Status,
		},
	}

	ser, er := getDocMangerServiceByUint64(httpReq.BarcodeID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	rpcUpdateBarcodeStatusResp, er := ser.Docmanager.UpdateBarcode(ctx, rpcUpdateBarcodeStatusReq)
	if er != nil || rpcUpdateBarcodeStatusResp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

func updateBarcodeStatusByBatch(ctx context.Context, barcodeID []uint64, barcodeStatus []uint64) (err error) {
	if len(barcodeID) != len(barcodeStatus) {
		return ErrUpdateBarcodeLengthNotMatch
	}

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		return fmt.Errorf("nil shardService")
	}

	templateUpdateBarcodeReq := &docmanagerModel.UpdateBarcodeReq{
		Barcode: &docmanagerModel.Barcode{
			ID:     0,
			Status: 0,
		},
	}

	for barcodeIndex := range barcodeID {
		ser, ok := shardService[uint64(core.GetShardID(uint32(barcodeID[barcodeIndex])))]
		if !ok {
			fmt.Println("nil shardID")
			return fmt.Errorf("no shard id")
		}

		templateUpdateBarcodeReq.Barcode.ID = barcodeID[barcodeIndex]
		templateUpdateBarcodeReq.Barcode.Status = barcodeStatus[barcodeIndex]

		resp, er := ser.Docmanager.UpdateBarcode(ctx, templateUpdateBarcodeReq)
		if er != nil || resp.Code != 0 {
			return fmt.Errorf("grpc Error: %v", er)
		}
	}

	return nil
}
