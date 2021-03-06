package doc

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"

	"github.com/google/uuid"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/dao/cache"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
)

// -----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- DOCUMENTS ----------------------------------------------------------

func updateDocument(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	var (
		ctx = c.Request().Context()
		req = request.(*portalModel.UpdateDocReq)
	)

	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "updateDocument",
	}

	// check category, create if not exist
	catID, er := cache.SelectCategoryID(ctx, core.GetDB(), req.Category)
	if er != nil && !errors.Is(er, sql.ErrNoRows) {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	if catID == 0 || errors.Is(er, sql.ErrNoRows) {
		catUUID, er := uuid.NewUUID()
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}

		catID = uint64(core.GetHash(catUUID.String()))
		newCategory := &portalModel.CategoriesDAOobj{
			CategoryID:   catID,
			CategoryName: req.Category,
		}

		if er = cache.SaveCategories(ctx, core.GetDB(), newCategory); er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
	}

	// update document
	updateDoc := &portalModel.DocumentsDAOobj{
		DocID:      req.DocID,
		DocName:    req.DocName,
		CategoryID: catID,
	}

	if er = cache.UpdateDocument(ctx, core.GetDB(), updateDoc); er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

// ------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- DOCUMENT VERSION ----------------------------------------------------------

func updateDocumentVersion(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	var (
		ctx = c.Request().Context()
		req = request.(*portalModel.UpdateDocVerReq)
	)

	// search author ID
	authUUID, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	authorID, er := cache.FirstOrCreateAuthor(ctx, core.GetDB(), req.Author, uint64(core.GetHash(authUUID.String())))
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	if err = cache.UpdateDocumentVersion(ctx, core.GetDB(), req.DocVerID, req.DocVer, req.Publisher, authorID, req.Price); err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

func addBarcodeByDocVerID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	var (
		ctx  = c.Request().Context()
		req  = request.(*portalModel.AddBarcodeByDocverIDReq)
		resp = &portalModel.AddBarcodeByDocverIDResp{
			Barcodes: make([]uint64, req.AddBarcodeNumber),
		}
	)

	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "Add Barcode By DocVer ID",
	}

	// create more barcode
	for i := uint64(0); i < req.AddBarcodeNumber; i++ {
		barcodeUUID, er := uuid.NewUUID()
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		newBarcodeID := uint64(core.GetHash(barcodeUUID.String()))
		rpcCreateBarcodeReq := &docmanagerModel.SaveBarcodeReq{
			Barcode: &docmanagerModel.Barcode{
				ID:       newBarcodeID,
				DocVerID: req.DocVerID,
				Status:   model.BarcodeNormalStatus,
			},
		}

		barcodeSer, er := getDocMangerServiceByUint64(newBarcodeID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		rpcCreateBarcodeResp, er := barcodeSer.Docmanager.SaveBarcode(ctx, rpcCreateBarcodeReq)
		if er != nil || rpcCreateBarcodeResp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, er
			return
		}

		// save to cache and response
		resp.Barcodes[i] = newBarcodeID
		if er := cache.SaveDocverIDToCache(ctx, core.GetDB(), newBarcodeID, req.DocVerID); er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
	}

	return http.StatusOK, resp, lg, false, nil
}

func saveDocVer(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	var (
		ctx     = c.Request().Context()
		httpReq = request.(*portalModel.CreateDocVerReq)
	)

	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "Add Barcode By DocVer ID",
	}

	// check author
	authUUID, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	authID, er := cache.FirstOrCreateAuthor(ctx, core.GetDB(), httpReq.Author, uint64(core.GetHash(authUUID.String())))
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	// save docver
	docVerUUID, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	newDocVer := &portalModel.DocumentVersionDAOobj{
		DocVerID:        uint64(core.GetHash(docVerUUID.String())),
		DocumentVersion: httpReq.DocVersion,
		DocID:           httpReq.DocID,
		Publisher:       httpReq.Publisher,
		AuthorID:        authID,
		Price:           httpReq.Price,
		DocDescription:  httpReq.DocDescription,
	}
	if er = cache.SaveDocumentVersion(ctx, core.GetDB(), newDocVer); er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	for i := uint64(0); i < httpReq.Number; i++ {
		barcodeUUID, er := uuid.NewUUID()
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		rpcSaveBarcodeReq := &docmanagerModel.SaveBarcodeReq{
			Barcode: &docmanagerModel.Barcode{
				ID:       uint64(core.GetHash(barcodeUUID.String())),
				DocVerID: newDocVer.DocVerID,
				Status:   model.BarcodeNormalStatus,
			},
		}

		ser, er := getDocMangerServiceByUint64(rpcSaveBarcodeReq.Barcode.ID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}

		rpcSaveBarcodeResp, er := ser.Docmanager.SaveBarcode(ctx, rpcSaveBarcodeReq)
		if er != nil || rpcSaveBarcodeResp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		// Save to cache: Barcode - DocVer
		if er := cache.SaveDocverIDToCache(ctx, core.GetDB(), rpcSaveBarcodeReq.Barcode.ID, rpcSaveBarcodeReq.Barcode.DocVerID); er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}

	}

	return http.StatusOK, nil, lg, false, nil
}
