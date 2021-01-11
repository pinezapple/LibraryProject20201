package model

import (
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
	DocID       uint64      `json:"doc_id"`
	Version     string      `json:"version"`
	Publisher   string      `json:"publisher"`
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

type UpdateBarcodeStatus struct {
	BarcodeID uint64 `json:"barcode_id"`
	Status    uint64 `json:"status"`
}

type DeleteBarCodeByIDReq struct {
	BarcodeID uint64 `json:"barcode_id"`
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

type SelectAllBorrowFormElement struct {
	ID            uint64      `protobuf:"varint,1,opt,name=ID,proto3" json:"borrow_form_id" db:"borrow_form_id"`
	LibrarianID   uint64      `protobuf:"varint,2,opt,name=LibrarianID,proto3" json:"librarian_id" db:"librarian_id"`
	LibrarianName string      `protobuf:"varint,2,opt,name=LibrarianID,proto3" json:"librarian_name" db:"librarian_name"`
	Status        uint64      `protobuf:"varint,3,opt,name=Status,proto3" json:"status" db:"status"`
	ReaderID      uint64      `protobuf:"varint,4,opt,name=ReaderID,proto3" json:"reader_id" db:"reader_id"`
	ReaderName    string      `protobuf:"varint,4,opt,name=ReaderID,proto3" json:"reader_name" db:"reader_name"`
	BarcodeID     []uint64    `protobuf:"varint,5,rep,packed,name=BarcodeID,proto3" json:"barcode_id" db:"barcode_id"`
	StartTime     *model.Time `protobuf:"bytes,6,opt,name=StartTime,proto3" json:"start_time" db:"start_time"`
	EndTime       *model.Time `protobuf:"bytes,7,opt,name=EndTime,proto3" json:"end_time" db:"end_time"`
	Fine          int64       `json:"fine"`
	CreatedAt     *model.Time `protobuf:"bytes,8,opt,name=Created_at,json=CreatedAt,proto3" json:"created_at" db:"created_at"`
	UpdatedAt     *model.Time `protobuf:"bytes,9,opt,name=Updated_at,json=UpdatedAt,proto3" json:"updated_at" db:"updated_at"`
}

type SelectAllBorrowFormResp struct {
	BorrowForms []*SelectAllBorrowFormElement `json:"borrow_forms"`
}

type SelectBorrowFormByIDReq struct {
	BorrowFormID uint64 `json:"borrow_form_id"`
}

type SelectBorrowFormByIDResp struct {
	BorrowFormID  uint64                 `json:"borrow_form_id"`
	LibrarianID   uint64                 `json:"librarian_id"`
	LibrarianName string                 `protobuf:"varint,2,opt,name=LibrarianID,proto3" json:"librarian_name" db:"librarian_name"`
	ReaderID      uint64                 `protobuf:"varint,4,opt,name=ReaderID,proto3" json:"reader_id" db:"reader_id"`
	ReaderName    string                 `protobuf:"varint,4,opt,name=ReaderID,proto3" json:"reader_name" db:"reader_name"`
	Status        uint64                 `json:"status"`
	Fine          int64                  `json:"fine"`
	Barcodes      []*RespBarcodeOverview `json:"barcodes"`
	StartTime     *model.Time            `json:"start_time"`
	EndTime       *model.Time            `json:"end_time"`
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
	LibrarianID   uint64 `json:"librarian_id"`
	ReaderID      uint64 `json:"reader_id"`
	BarcodeUpdate []struct {
		BarcodeID     uint64 `json:"barcode_id"`
		BarcodeStatus uint64 `json:"barcode_status"`
		Fee           uint64 `json:"fee"`
	} `json:"barcode_update"`
	Fine uint64 `json:"fine"`
}

// ----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- PAYMENTS ----------------------------------------------------------

type SelectAllPaymentResp struct {
	PaymentID     uint64      `json:"payment_id"`
	BorrowFormID  uint64      `json:"borrow_form_id"`
	TotalMoney    uint64      `json:"total_money"`
	LibrarianID   uint64      `json:"librarian_id"`
	LibrarianName string      `json:"librarian_name"`
	ReaderID      uint64      `json:"reader_id"`
	ReaderName    string      `json:"reader_name"`
	CreatedAt     *model.Time `json:"created_at"`
}

type RespBarcodePaymentOverview struct {
	BarcodeID  uint64 `json:"barcode_id"`
	Status     uint64 `json:"status"`
	DocName    string `json:"doc_name"`
	AuthorName string `json:"author_name"`
	Money      uint64 `json:"money"`
}

type SelectPaymentByIDReq struct {
	PaymentID    uint64 `json:"payment_id"`
	BorrowFormID uint64 `json:"borrow_form_id"`
}

type SelectPaymentByIDResp struct {
	PaymentID    uint64                        `json:"payment_id"`
	BorrowFormID uint64                        `json:"borrow_form_id"`
	LibrarianID  uint64                        `json:"librarian_id"`
	ReaderID     uint64                        `json:"reader_id"`
	TotalMoney   uint64                        `json:"total_money"`
	Fine         uint64                        `json:"fine"`
	Barcodes     []*RespBarcodePaymentOverview `json:"barcodes"`
	CreatedAt    *model.Time                   `json:"created_at"`
}

// ------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- SALE BILLS ----------------------------------------------------------

type SelectAllSaleBillResp struct {
	SaleBillID    uint64      `json:"sale_bill_id"`
	LibrarianID   uint64      `json:"librarian_id"`
	LibrarianName string      `json:"librian_name"`
	TotalMoney    uint64      `json:"total_money"`
	CreatedAt     *model.Time `json:"created_at"`
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
	SaleBillID    uint64            `json:"sale_bill_id"`
	LibrarianID   uint64            `json:"librarian_id"`
	LibrarianName string            `json:"librian_name"`
	TotalMoney    uint64            `json:"total_money"`
	Barcodes      []*SaleBillDetail `json:"barcodes"`
	CreatedAt     *model.Time       `json:"created_at"`
}

type CreateSaleBillReq struct {
	LibrarianID uint64 `json:"librian_id"`
	SaleBarcode []*struct {
		BarcodeID uint64 `json:"barcode_id"`
		Price     uint64 `json:"price"`
	} `json:"sale_barcode"`
}

// ------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- USER ----------------------------------------------------------

type SelectAllViolatedUserReq struct {
}

type SelectAllViolatedUserResp struct {
	Users []*model.User `json:"users"`
}

type SaveDocReq struct {
	DocName     string `json:"doc_name"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Publisher   string `json:"publisher"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Number      uint64 `json:"number"`
	Price       uint64 `json:"price"`
}

type SaveDocResp struct {
	Barcodes []uint64 `json:"barcodes"`
}

// -----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- DOCUMENTS ----------------------------------------------------------

//TODO: GetAllDocReq & Resp

type UpdateDocReq struct {
	DocID    uint64 `json:"document_id"`   //
	DocName  string `json:"document_name"` //
	Category string `json:"category"`      //First or create
}

// ------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- DOCUMENT VERSION ----------------------------------------------------------

//TODO: GetAllDocVerReq & Resp

type UpdateDocVerReq struct {
	DocVerID  uint64 `json:"document_version_id"` //
	DocVer    string `json:"document_version"`    //
	Publisher string `json:"publisher"`           //First or create
	Author    string `json:"author"`              //First or create
	Price     uint64 `json:"price"`               //
}

type AddBarcodeByDocverIDReq struct {
	DocVerID         uint64 `json:"document_version_id"`
	AddBarcodeNumber uint64 `json:"add_barcode_number"`
}

type AddBarcodeByDocverIDResp struct {
	Barcodes []uint64 `json:"barcodes"`
}

type CreateDocVerReq struct {
	DocID          uint64 `json:"document_id"`
	Price          uint64 `json:"price"`
	Number         uint64 `json:"number"`
	DocVersion     string `json:"document_version"`
	DocDescription string `json:"document_description"`
	Publisher      string `json:"publisher"`
	Author         string `json:"author"`
}

// ------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- BLACK LIST----------------------------------------------------------

type SelectByUserIDReq struct {
	UserID uint64 `json:"user_id"`
}

type BlackListSelectAllElement struct {
	UserID   uint64 `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Count    int    `json:"count" db:"count"`
	Money    uint64 `json:"total_money" db:"total_money"`
}

type BlackListSelectByIDResp struct {
	UserID   uint64             `json:"user_id" db:"user_id"`
	Username string             `json:"username" db:"username"`
	Count    int                `json:"count" db:"count"`
	Money    uint64             `json:"total_money" db:"total_money"`
	Detail   []*BlackListDAOobj `json:"detail" db:"detail"`
}

// ------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- DOC ----------------------------------------------------------

type SelectDocByIDReq struct {
	DocID uint64 `json:"doc_id"`
}

type SelectDocVerByIDReq struct {
	DocVerID uint64 `json:"document_version_id"`
}

type SelectAllDocumentElement struct {
	DocID        uint64      `json:"doc_id" db:"doc_id"`
	DocName      string      `json:"doc_name" db:"doc_name"`
	CategoryName string      `json:"category_name" db:"category_name"`
	CreatedAt    *model.Time `json:"created_at" db:"created_at"`
}

type SelectDocVerByIDResp struct {
	DocID        uint64 `json:"doc_id" db:"doc_id"`
	DocName      string `json:"doc_name" db:"doc_name"`
	CategoryName string `json:"category_name" db:"category_name"`

	DocVerID   uint64 `json:"document_version_id"`
	DocVerName string `json:"document_version"`
	Publisher  string `json:"publisher"`
	AuthorName string `json:"author_name"`
	Count      uint64 `json:"count"`
	Price      uint64 `json:"price"`

	Barcode []*docmanagerModel.Barcode `json:"barcode"`

	CreatedAt *model.Time `json:"created_at" db:"created_at"`
}

type DocverOverviewElement struct {
	DocVerID   uint64 `json:"document_version_id"`
	DocVerName string `json:"document_version"`
	Publisher  string `json:"publisher"`
	AuthorName string `json:"author_name"`
	Count      uint64 `json:"count"`
	Price      uint64 `json:"price"`
}

type SelectDocumentByID struct {
	DocID        uint64                   `json:"doc_id" db:"doc_id"`
	DocName      string                   `json:"doc_name" db:"doc_name"`
	CategoryName string                   `json:"category_name" db:"category_name"`
	Docver       []*DocverOverviewElement `json:"doc_version"`
	CreatedAt    *model.Time              `json:"created_at" db:"created_at"`
}
