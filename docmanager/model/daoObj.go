package model

import "github.com/pinezapple/LibraryProject20201/skeleton/model"

type BorrowFormDAOobj struct {
	ID          uint64      `protobuf:"varint,1,opt,name=ID,proto3" json:"borrow_form_id" db:"borrow_form_id"`
	LibrarianID uint64      `protobuf:"varint,2,opt,name=LibrarianID,proto3" json:"librarian_id" db:"librarian_id"`
	Status      uint64      `protobuf:"varint,3,opt,name=Status,proto3" json:"status" db:"status"`
	BarcodeID   []byte      `protobuf:"varint,4,rep,packed,name=BarcodeID,proto3" json:"barcode_id" db:"barcode_id"`
	StartTime   *model.Time `protobuf:"bytes,6,opt,name=StartTime,proto3" json:"start_time" db:"start_time"`
	EndTime     *model.Time `protobuf:"bytes,7,opt,name=EndTime,proto3" json:"end_time" db:"end_time"`
	CreatedAt   *model.Time `protobuf:"bytes,8,opt,name=Created_at,json=CreatedAt,proto3" json:"created_at" db:"created_at"`
	UpdatedAt   *model.Time `protobuf:"bytes,9,opt,name=Updated_at,json=UpdatedAt,proto3" json:"updated_at" db:"updated_at"`
}

type PaymentDAOobj struct {
	ID            uint64      `protobuf:"varint,1,opt,name=ID,proto3" json:"id_sale_bill" db:"id_sale_bill"`
	BorrowFormID  byte        `json:"borrow_form_id" db:"borrow_form_id"`
	BarcodeId     []byte      `protobuf:"varint,2,rep,packed,name=barcode_id,json=barcodeId,proto3" json:"barcode_id" db:"barcode_id"`
	BarcodeStatus []byte      `json:"barcode_status" db:"barcode_status"`
	Price         []byte      `protobuf:"varint,3,rep,packed,name=price,proto3" json:"price" db:"price"`
	CreatedAt     *model.Time `protobuf:"bytes,4,opt,name=Created_at,json=CreatedAt,proto3" json:"created_at" db:"created_at"`
}

type SaleBillDAOobj struct {
	ID        uint64      `protobuf:"varint,1,opt,name=ID,proto3" json:"id_sale_bill" db:"id_sale_bill"`
	BarcodeId []byte      `protobuf:"varint,2,rep,packed,name=barcode_id,json=barcodeId,proto3" json:"barcode_id" db:"barcode_id"`
	Price     []byte      `protobuf:"varint,3,rep,packed,name=price,proto3" json:"price" db:"price"`
	CreatedAt *model.Time `protobuf:"bytes,4,opt,name=Created_at,json=CreatedAt,proto3" json:"created_at" db:"created_at"`
}
