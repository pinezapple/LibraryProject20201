package doc

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

var (
	ErrCreatePaymentLengthNotMatch = fmt.Errorf("len barcodeID != len barcodeID != len barcodeStatus")
)

func createPayment(ctx context.Context, borrowFormID uint64, librarianID uint64, readerID uint64, barcodeID []uint64, barcodeStatus []uint64, money []uint64, fine uint64) (err error) {
	if len(barcodeID) != len(money) && len(barcodeID) != len(barcodeStatus) {
		return ErrCreatePaymentLengthNotMatch
	}

	// get shard by borrow form ID
	ser, err := getDocMangerServiceByUint64(borrowFormID)
	if err != nil {
		return err
	}

	// create payment ID
	paymentUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	rpcCreatePaymentreq := &docmanagerModel.SavePaymentReq{
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

	// rpc
	rpcCreatePaymentresp, err := ser.Docmanager.SavePayment(ctx, rpcCreatePaymentreq)
	if err != nil || rpcCreatePaymentresp.Code != 0 {
		return err
	}

	return nil
}
