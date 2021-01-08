package doc

import (
	"context"
	"errors"
	"fmt"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

var (
	ErrUpdateBarcodeLengthNotMatch = errors.New("len barcodes != len barCodestatus")
)

func updateBarcodeStatusByBatch(ctx context.Context, barcodeID []uint64, barcodeStatus []uint64) (err error) {
	if len(barcodeID) != len(barcodeStatus) {
		return ErrUpdateBarcodeLengthNotMatch
	}

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
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
