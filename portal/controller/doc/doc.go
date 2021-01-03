package doc

import (
	"fmt"

	"github.com/labstack/echo"
)

type reqSelectByID struct {
	ID uint64 `json:"id"`
}

type reqSelectFormByID struct {
	FormID uint64 `json:"form_id"`
	DocID  uint64 `json:"doc_id"`
}
type reqUpdateStatus struct {
	FormID uint64 `json:"form_id"`
	DocID  uint64 `json:"doc_id"`
	Status int    `json:"status"`
}

type reqDoc struct {
	ID           uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id_doc" db:"id_doc"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"doc_name" db:"doc_name"`
	Author       string `protobuf:"bytes,3,opt,name=Author,proto3" json:"doc_author" db:"doc_author"`
	Type         string `protobuf:"bytes,4,opt,name=Type,proto3" json:"doc_type" db:"doc_type"`
	Description  string `protobuf:"bytes,5,opt,name=Description,proto3" json:"doc_description" db:"doc_description"`
	Status       uint32 `protobuf:"varint,6,opt,name=Status,proto3" json:"status" db:"status"`
	BorrowFormID uint64 `protobuf:"varint,7,opt,name=BorrowFormID,proto3" json:"id_borrow" db:"id_borrow"`
	Fee          uint32 `protobuf:"varint,8,opt,name=Fee,proto3" json:"fee" db:"fee"`
}

type reqSaveBorrowForm struct {
	ID     uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id_borrow" db:"id_borrow"`
	DocID  uint64 `protobuf:"varint,2,opt,name=DocID,proto3" json:"id_doc" db:"id_doc"`
	CusID  uint64 `protobuf:"varint,3,opt,name=CusID,proto3" json:"id_cus" db:"id_cus"`
	LibID  uint64 `protobuf:"varint,4,opt,name=LibID,proto3" json:"id_lib" db:"id_lib"`
	Status int32  `protobuf:"varint,5,opt,name=Status,proto3" json:"status" db:"status"`
	TTL    int32  `protobuf:"varint,6,opt,name=TTL,proto3" json:"ttl" db:"ttl"`
}

type reqSaveDoc struct {
	DocName     string `json:"doc_name"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Publisher   string `json:"publisher"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Number      uint64 `json:"number"`
	Price       uint64 `json:"price"`
}

func SelectAllBarcode(c echo.Context) (erro error) {
	return
}

func SelectAllAvailableBarcode(c echo.Context) (erro error) {
	return
}

func SelectAllSellingBarcode(c echo.Context) (erro error) {
	return
}

func SelectAllDamagedBarcode(c echo.Context) (erro error) {
	return
}

func SelectAllBorrowForm(c echo.Context) (erro error) {
	return
}

func SelectAllNotReturnedBorrowForm(c echo.Context) (erro error) {
	return
}

func SelectAllPayment(c echo.Context) (erro error) {
	return
}

func SelectAllSaleBill(c echo.Context) (erro error) {
	return
}

func SelectBarcodeByID(c echo.Context) (erro error) {
	return
}

func SelectBorrowFormByID(c echo.Context) (erro error) {
	return
}

func SelectPaymentByID(c echo.Context) (erro error) {
	return
}

func SelectSaleBillByID(c echo.Context) (erro error) {
	return
}

func SaveDocumentByBatch(c echo.Context) (erro error) {
	fmt.Println("In Save Document By Batch")
	return
}
