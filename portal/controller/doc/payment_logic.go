package doc

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

var (
	ErrCreatePaymentLengthNotMatch = fmt.Errorf("len barcodeID != len barcodeID != len barcodeStatus")
)

func createPayment(ctx context.Context, librarianID uint64, readerID uint64, borrowFormID uint64, barcodeID []uint64, barcodeStatus []uint64, money []uint64) (err error) {
	if len(barcodeID) != len(money) && len(barcodeID) != len(barcodeStatus) {
		return ErrCreatePaymentLengthNotMatch
	}

	// create payment ID
	paymentUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	reqCreatePayment := &docmanagerModel.SavePaymentReq{
		Payment: &docmanagerModel.Payment{
			ID:            uint64(core.GetHash(paymentUUID.String())),
			LibrarianID:   librarianID,
			ReaderID:      readerID,
			BorrowFormID:  borrowFormID,
			BarcodeID:     barcodeID,
			BarcodeStatus: barcodeStatus,
			Money:         money,
		},
	}

	// get shard by borrow form ID
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		return fmt.Errorf("nil shardService")
	}
	shardID := core.GetShardID(uint32(borrowFormID))
	ser, ok := shardService[uint64(shardID)]
	if !ok {
		fmt.Println("nil shardID")
		return fmt.Errorf("no shard id")
	}

	// rpc
	resp, err := ser.Docmanager.SavePayment(ctx, reqCreatePayment)
	if err != nil || resp.Code != 0 {
		return err
	}

	return nil
}
