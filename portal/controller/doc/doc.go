package doc

import (
	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/controller"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
)

func SelectAllBarcode(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBarcode)
}

func SelectAllAvailableBarcode(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBarcodeAvail)
}

func SelectAllSellingBarcode(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBarcodeSelling)
}

func SelectAllDamagedBarcode(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBarcodeDamaged)
}

func SelectAllBorrowForm(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBorrowForm)
}

func SelectAllNotReturnedBorrowForm(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllUnreturnedBorrowForm)
}

func SelectAllPayment(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllPayment)
}

func SelectPaymentWithFine(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectPaymentWithFine)
}

func SelectAllSaleBill(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllSaleBill)
}

func SelectBarcodeByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectBarcodeByIDReq{}, selectBarcodeByID)
}

func SelectBorrowFormByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectBorrowFormByIDReq{}, selectBorrowFormByID)
}

func SelectPaymentByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectPaymentByIDReq{}, selectPaymentByID)
}

func SelectSaleBillByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectSaleBillByIDReq{}, selectSaleBillByID)
}

func SelectAllBlackList(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllFromBlackList)
}
func SelectBlackListByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectByUserIDReq{}, selectBlackListByUserID)
}

func SelectAllDoc(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllDoc)
}

func SelectDocByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectDocByIDReq{}, selectDocByID)
}

func SelectDocVerByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectDocVerByIDReq{}, selectDocVerByID)
}
