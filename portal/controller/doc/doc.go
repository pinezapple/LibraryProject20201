package doc

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/controller"
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

func SelectAllDoc(c echo.Context) (erro error) {

	fmt.Println("in Select doc all")
	return controller.ExecHandler(c, nil, selectAllDoc)
}

func SelectAllDoc0(c echo.Context) (erro error) {

	fmt.Println("in Select doc 0 all")
	return controller.ExecHandler(c, nil, selectAllDoc0)
}

func SelectDocByID(c echo.Context) (erro error) {

	fmt.Println("in select by id")
	return controller.ExecHandler(c, &reqSelectByID{}, selectDocByID)
}

func SaveDoc(c echo.Context) (erro error) {

	fmt.Println("in save doc")
	return controller.ExecHandler(c, &reqDoc{}, saveDoc)
}

func DelDoc(c echo.Context) (erro error) {
	fmt.Println("in del doc")
	return controller.ExecHandler(c, &reqSelectByID{}, delDoc)
}

func UpdateDoc(c echo.Context) (erro error) {
	fmt.Println("in update doc")
	return controller.ExecHandler(c, &reqDoc{}, updateDoc)
}

func UpdateStatus(c echo.Context) (erro error) {
	fmt.Println("in Update Status")
	return controller.ExecHandler(c, &reqUpdateStatus{}, updateStatus)

}

func SelectAllForm(c echo.Context) (erro error) {
	fmt.Println("in Select all form")
	return controller.ExecHandler(c, nil, selectAllForm)
}

func SaveForm(c echo.Context) (erro error) {
	fmt.Println("in save form")
	return controller.ExecHandler(c, &reqSaveBorrowForm{}, saveForm)
}

func SelectFormByID(c echo.Context) (erro error) {
	fmt.Println("in form by id")
	return controller.ExecHandler(c, &reqSelectFormByID{}, selectFormByID)
}
