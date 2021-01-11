package doc

import (
	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/controller"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
)

// -------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- BORROW FORM ----------------------------------------------------------

func CreateBorrowForm(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.CreateBorrowFormReq{}, createBorrowForm)
}

func UpdateBorrowForm(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.UpdateBorrowFormReq{}, updateBorrowForm)
}

// -------------------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- DOCUMENT AND DOCUMENT VERSION ----------------------------------------------------------

func SaveDocumentByBatch(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SaveDocReq{}, saveDocumentByBatch)
}

func UpdateDocument(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.UpdateDocReq{}, updateDocument)
}

func UpdateDocumentVersion(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.UpdateDocVerReq{}, updateDocumentVersion)
}

func AddBarcodeUpdateDocumentVersion(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.AddBarcodeByDocverIDReq{}, addBarcodeByDocVerID)
}

func CreateDocumentVersion(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.CreateDocVerReq{}, saveDocVer)
}

// ---------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- BARCODE ----------------------------------------------------------

func UpdateBarcodeStatus(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.UpdateBarcodeStatus{}, updateBarcodeStatus)
}

func DeleteBarcode(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.DeleteBarCodeByIDReq{}, deleteBarcodeByID)
}

// -----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- SALE BILL ----------------------------------------------------------

func CreateSaleBill(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.CreateSaleBillReq{}, saveSaleBill)
}
