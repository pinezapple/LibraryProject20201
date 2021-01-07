package model

import (
	"time"

	"github.com/pinezapple/LibraryProject20201/skeleton/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

// ----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- BARCODES ----------------------------------------------------------

type RespBarcodeOverview struct {
	BarcodeID uint64 `json:"barcode_id"`
	Status    uint64 `json:"status"`
	DocName   string `json:"doc_name"`
	Author    string `json:"author"`
	Fee       uint64 `json:"fee"`
	Price     uint64 `json:"price"`
}

type SelectAllAvailableBarcodeReq struct {
}

type SelectAllAvailableBarcodeResp struct {
	Barcodes []*RespBarcodeOverview `json:"barcodes"`
}

type SelectBarcodeByIDReq struct {
	BarcodeID uint64 `json:"barcode_id"`
}

type BarcodeFrontEndResp struct {
	BarcodeID   uint64      `json:"barcode_id"`
	Status      uint64      `json:"status"`
	DocName     string      `json:"doc_name"`
	Version     string      `json:"version"`
	Author      string      `json:"author"`
	Fee         uint64      `json:"fee"`
	Price       uint64      `json:"price"`
	Description string      `json:"description"`
	Category    string      `json:"category"`
	SaleBillID  uint64      `json:"sale_bill_id"`
	CreatedAt   *model.Time `json:"created_at"`
}

type SelectAllSellingBarcodesReq struct {
}

type SelectAllSellingBarcodesResp struct {
	Barcodes []*RespBarcodeOverview `json:"barcodes"`
}

type SelectAllDamageBarcodesReq struct {
}

type SelectAllDamageBarcodesResp struct {
	Barcodes []*RespBarcodeOverview `json:"barcodes"`
}

// --------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- BORROW FORMS ----------------------------------------------------------

type SelectAllUnReturnBorrowFormReq struct {
}

type SelectAllUnReturnBorrowFormResp struct {
	BorrowForms []*docmanagerModel.BorrowForm `json:"borrow_forms"`
}

type SelectAllBorrowFormReq struct {
}

type SelectAllBorrowFormResp struct {
	BorrowForms []*docmanagerModel.BorrowForm `json:"borrow_forms"`
}

type SelectBorrowFormByIDReq struct {
	BorrowFormID uint64 `json:"borrow_form_id"`
}

type SelectBorrowFormByIDResp struct {
	BorrowFormID uint64                 `json:"borrow_form_id"`
	LibrarianID  uint64                 `json:"librarian_id"`
	Status       uint64                 `json:"status"`
	Barcodes     []*RespBarcodeOverview `json:"barcodes"`
	StartTime    time.Time              `json:"start_time"`
	EndTime      time.Time              `json:"end_time"`
}

type CreateBorrowFormReq struct {
	LibrarianID uint64   `json:"librarian_id"`
	ReaderID    uint64   `json:"reader_id"`
	BorrowDays  uint64   `json:"borrow_days"`
	Barcodes    []uint64 `json:"barcodes"`
}

type UpdateBorrowFormReq struct {
	BorrowFormID  uint64 `json:"borrow_form_id"`
	Status        uint64 `json:"status"`
	BarcodeUpdate []struct {
		BarcodeID     uint64 `json:"barcode_id"`
		BarcodeStatus uint64 `json:"barcode_status"`
		Fee           uint64 `json:"fee"`
	} `json:"barcode_update"`
}

// ----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- PAYMENTS ----------------------------------------------------------

type SelectAllPaymentResp struct {
	PaymentID    uint64      `json:"payment_id"`
	BorrowFormID uint64      `json:"borrow_form_id"`
	TotalMoney   uint64      `json:"total_money"`
	LibrarianID  uint64      `json:"librarian_id"`
	CreatedAt    *model.Time `json:"created_at"`
}

type RespBarcodePaymentOverview struct {
	BarcodeID uint64 `json:"barcode_id"`
	Status    uint64 `json:"status"`
	DocName   string `json:"doc_name"`
	Money     uint64 `json:"money"`
}

type SelectPaymentByIDReq struct {
	PaymentID    uint64 `json:"payment_id"`
	BorrowFormID uint64 `json:"borrow_form_id"`
}

type SelectPaymentByIDResp struct {
	PaymentID    uint64                        `json:"payment_id"`
	BorrowFormID uint64                        `json:"borrow_form_id"`
	LibrarianID  uint64                        `json:"librarian_id"`
	TotalMoney   uint64                        `json:"total_money"`
	Barcodes     []*RespBarcodePaymentOverview `json:"barcodes"`
}

// ------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- SALE BILLS ----------------------------------------------------------
type SelectAllSaleBillResp struct {
	SaleBillID  uint64 `json:"sale_bill_id"`
	LibrarianID uint64 `json:"librarian_id"`
	TotalMoney  uint64 `json:"total_money"`
}

type SaleBillDetail struct {
	BarcodeID uint64 `json:"barcode_id"`
	DocName   string `json:"doc_name"`
	Money     uint64 `json:"money"`
}

type SelectSaleBillByIDReq struct {
	SaleBillID uint64 `json:"sale_bill_id"`
}

type SelectSaleBillByIDResp struct {
	SaleBillID  uint64            `json:"sale_bill_id"`
	LibrarianID uint64            `json:"librian_id"`
	TotalMoney  uint64            `json:"total_money"`
	Barcodes    []*SaleBillDetail `json:"barcodes"`
}

// ------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- USER ----------------------------------------------------------

type SelectAllViolatedUserReq struct {
}

type SelectAllViolatedUserResp struct {
	Users []*model.User `json:"users"`
}
