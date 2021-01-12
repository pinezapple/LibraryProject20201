package model

import "github.com/pinezapple/LibraryProject20201/skeleton/model"

type BorrowFormDAOobj struct {
	ID          uint64      `protobuf:"varint,1,opt,name=ID,proto3" json:"borrow_form_id" db:"borrow_form_id"`
	LibrarianID uint64      `protobuf:"varint,2,opt,name=LibrarianID,proto3" json:"librarian_id" db:"librarian_id"`
	Status      uint64      `protobuf:"varint,3,opt,name=Status,proto3" json:"status" db:"status"`
	ReaderID    uint64      `protobuf:"varint,4,opt,name=ReaderID,proto3" json:"reader_id" db:"reader_id"`
	BarcodeID   []byte      `protobuf:"varint,5,rep,packed,name=BarcodeID,proto3" json:"barcode_id" db:"barcode_id"`
	StartTime   *model.Time `protobuf:"bytes,6,opt,name=StartTime,proto3" json:"borrow_start_time" db:"borrow_start_time"`
	EndTime     *model.Time `protobuf:"bytes,7,opt,name=EndTime,proto3" json:"borrow_end_time" db:"borrow_end_time"`
	CreatedAt   *model.Time `protobuf:"bytes,8,opt,name=Created_at,json=CreatedAt,proto3" json:"created_at" db:"created_at"`
	UpdatedAt   *model.Time `protobuf:"bytes,9,opt,name=Updated_at,json=UpdatedAt,proto3" json:"updated_at" db:"updated_at"`
}

type PaymentDAOobj struct {
	ID            uint64      `protobuf:"varint,1,opt,name=ID,proto3" json:"payment_id" db:"payment_id"`
	LibrarianID   uint64      `protobuf:"varint,3,opt,name=LibrarianID,proto3" json:"librarian_id" db:"librarian_id"`
	BorrowFormID  uint64      `protobuf:"varint,4,opt,name=BorrowFormID,proto3" json:"borrow_form_id" db:"borrow_form_id"`
	Fine          uint64      `protobuf:"varint,5,opt,name=Fine,proto3" json:"fine" db:"fine"`
	ReaderID      uint64      `protobuf:"varint,6,opt,name=ReaderID,proto3" json:"reader_id" db:"reader_id"`
	BarcodeID     []byte      `protobuf:"varint,7,rep,packed,name=BarcodeID,proto3" json:"barcode_id" db:"barcode_id"`
	BarcodeStatus []byte      `protobuf:"varint,8,rep,packed,name=BarcodeStatus,proto3" json:"barcode_status" db:"barcode_status"`
	Money         []byte      `protobuf:"varint,9,rep,packed,name=Money,proto3" json:"money" db:"money"`
	CreatedAt     *model.Time `protobuf:"bytes,10,opt,name=Created_at,json=CreatedAt,proto3" json:"created_at" db:"created_at"`
	UpdatedAt     *model.Time `protobuf:"bytes,11,opt,name=Updated_at,json=UpdatedAt,proto3" json:"updated_at" db:"updated_at"`
}

type SaleBillDAOobj struct {
	ID          uint64      `protobuf:"varint,1,opt,name=ID,proto3" json:"sale_bill_id" db:"sale_bill_id"`
	LibrarianID uint64      `protobuf:"varint,2,opt,name=LibrarianID,proto3" json:"librarian_id" db:"librarian_id"`
	BarcodeId   []byte      `protobuf:"varint,2,rep,packed,name=barcode_id,json=barcodeId,proto3" json:"barcode_id" db:"barcode_id"`
	Price       []byte      `protobuf:"varint,3,rep,packed,name=sale_price,proto3" json:"sale_price" db:"sale_price"`
	CreatedAt   *model.Time `protobuf:"bytes,4,opt,name=Created_at,json=CreatedAt,proto3" json:"created_at" db:"created_at"`
}
