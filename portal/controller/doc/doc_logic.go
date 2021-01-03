package doc

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/dao/cache"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"

	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
)

func selectAllBarcode(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all barcode from shard", Data: ""}
	shardNum := core.ShardNumber
	db := core.GetDB()

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*portalModel.BarcodeFrontEndResp
	for i := 0; i < shardNum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectAllBarcode(ctx, &docmanagerModel.SelectAllBarcodeReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		for j := 0; j < len(resp.Barcodes); j++ {
			docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcodes[j].DocVer)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}
			doc, er := cache.SelectDocumentByID(ctx, db, docver.DocID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}
			aut, er := cache.SelectAuthorByID(ctx, db, docver.AuthorID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return

			}
			cat, er := cache.SelectCategoriesByID(ctx, db, doc.CategoryID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}

			tmp := &portalModel.BarcodeFrontEndResp{
				BarcodeID:   resp.Barcodes[j].ID,
				Status:      resp.Barcodes[j].Status,
				DocName:     doc.DocName,
				Version:     docver.Version,
				Author:      aut.AuthorName,
				Fee:         docver.Fee,
				Price:       docver.Price,
				Category:    cat.CategoryName,
				SaleBillID:  resp.Barcodes[j].SaleBillID,
				Description: docver.DocDescription,
				CreatedAt:   resp.Barcodes[j].CreatedAt,
			}

			finalResp = append(finalResp, tmp)
		}

	}
	data = finalResp

	return
}

func selectAllBarcodeSelling(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all selling barcode from shard", Data: ""}
	db := core.GetDB()
	shardNum := core.ShardNumber

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*portalModel.BarcodeFrontEndResp
	for i := 0; i < shardNum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectAllBarcode(ctx, &docmanagerModel.SelectAllBarcodeReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		for j := 0; j < len(resp.Barcodes); j++ {
			docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcodes[j].DocVer)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}
			doc, er := cache.SelectDocumentByID(ctx, db, docver.DocID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}
			aut, er := cache.SelectAuthorByID(ctx, db, docver.AuthorID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return

			}
			cat, er := cache.SelectCategoriesByID(ctx, db, doc.CategoryID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}

			tmp := &portalModel.BarcodeFrontEndResp{
				BarcodeID:   resp.Barcodes[j].ID,
				Status:      resp.Barcodes[j].Status,
				DocName:     doc.DocName,
				Version:     docver.Version,
				Author:      aut.AuthorName,
				Fee:         docver.Fee,
				Price:       docver.Price,
				Category:    cat.CategoryName,
				SaleBillID:  resp.Barcodes[j].SaleBillID,
				Description: docver.DocDescription,
				CreatedAt:   resp.Barcodes[j].CreatedAt,
			}

			finalResp = append(finalResp, tmp)
		}

	}
	data = finalResp

	return
}

func selectAllBarcodeDamaged(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	db := core.GetDB()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all damaged barcode from shards", Data: ""}
	shardNum := core.ShardNumber

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*portalModel.BarcodeFrontEndResp
	for i := 0; i < shardNum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectAllBarcode(ctx, &docmanagerModel.SelectAllBarcodeReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		for j := 0; j < len(resp.Barcodes); j++ {
			docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcodes[j].DocVer)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}
			doc, er := cache.SelectDocumentByID(ctx, db, docver.DocID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}
			aut, er := cache.SelectAuthorByID(ctx, db, docver.AuthorID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return

			}
			cat, er := cache.SelectCategoriesByID(ctx, db, doc.CategoryID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}

			tmp := &portalModel.BarcodeFrontEndResp{
				BarcodeID:   resp.Barcodes[j].ID,
				Status:      resp.Barcodes[j].Status,
				DocName:     doc.DocName,
				Version:     docver.Version,
				Author:      aut.AuthorName,
				Fee:         docver.Fee,
				Price:       docver.Price,
				Category:    cat.CategoryName,
				SaleBillID:  resp.Barcodes[j].SaleBillID,
				Description: docver.DocDescription,
				CreatedAt:   resp.Barcodes[j].CreatedAt,
			}

			finalResp = append(finalResp, tmp)
		}

	}
	data = finalResp

	return
}

func selectAllSaleBill(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all Sale bill from shards", Data: ""}
	shardNum := core.ShardNumber

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*docmanagerModel.SaleBill

	for i := 0; i < shardNum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectAllSaleBill(ctx, &docmanagerModel.SelectAllSaleBillReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		finalResp = append(finalResp, resp.SaleBills...)
	}
	data = finalResp

	return
}

func selectAllPayment(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all payment from shards", Data: ""}
	shardNum := core.ShardNumber

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*docmanagerModel.Payment

	for i := 0; i < shardNum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectAllPayment(ctx, &docmanagerModel.SelectAllPaymentReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		finalResp = append(finalResp, resp.Payments...)
	}
	data = finalResp

	return
}

func selectAllBorrowForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all borrow form", Data: ""}
	shardNum := core.ShardNumber

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*docmanagerModel.BorrowForm

	for i := 0; i < shardNum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectAllBorrowForm(ctx, &docmanagerModel.SelectAllBorrowFormReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		finalResp = append(finalResp, resp.BorrowForms...)
	}
	data = finalResp

	return
}

func selectAllUnreturnedBorrowForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all unreturned borrow form", Data: ""}
	shardNum := core.ShardNumber

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*docmanagerModel.BorrowForm

	for i := 0; i < shardNum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectAllUnReturnBorrowForm(ctx, &docmanagerModel.SelectAllUnReturnBorrowFormReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		finalResp = append(finalResp, resp.BorrowForms...)
	}
	data = finalResp

	return
}

func selectBarcodeByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*portalModel.SelectBarcodeByIDReq)
	db := core.GetDB()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select barcode by id", Data: ""}

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	shardId := core.GetShardID(uint32(req.BarcodeID))
	ser, ok := shardService[uint64(shardId)]
	if !ok {
		fmt.Println("nil shardID")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return
	}

	resp, er := ser.Docmanager.SelectBarcodeByID(ctx, &docmanagerModel.SelectBarcodeByIDReq{BarcodeID: req.BarcodeID})
	if er != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}
	docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcode.DocVer)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	doc, er := cache.SelectDocumentByID(ctx, db, docver.DocID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	aut, er := cache.SelectAuthorByID(ctx, db, docver.AuthorID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return

	}
	cat, er := cache.SelectCategoriesByID(ctx, db, doc.CategoryID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	data = &portalModel.BarcodeFrontEndResp{
		BarcodeID:   resp.Barcode.ID,
		Status:      resp.Barcode.Status,
		DocName:     doc.DocName,
		Version:     docver.Version,
		Author:      aut.AuthorName,
		Fee:         docver.Fee,
		Price:       docver.Price,
		Category:    cat.CategoryName,
		SaleBillID:  resp.Barcode.SaleBillID,
		Description: docver.DocDescription,
		CreatedAt:   resp.Barcode.CreatedAt,
	}

	// TODO: filter more information
	return
}

func saveDocumentByBatch(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	fmt.Println("in saveDocumentByBatch")
	req := request.(*reqSaveDoc)
	ctx := c.Request().Context()
	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "saveDocumentByBatch",
		Data:   req,
	}

	// search author ID

	authorID, er := cache.SelectAuthorID(ctx, core.GetDB(), req.Author)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	// search category ID

	catID, er := cache.SelectCategoryID(ctx, core.GetDB(), req.Category)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	// check if exist document, then create
	uuidDocId, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	reqDoc := &portalModel.DocumentsDAOobj{
		DocID:      uint64(core.GetHash(uuidDocId.String())),
		DocName:    req.DocName,
		CategoryID: catID,
	}

	docID, er := cache.FirstOrCreateDocument(ctx, core.GetDB(), reqDoc)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	if docID != 0 {
		reqDoc.DocID = docID
	}

	// check if exist document version, then create
	uuidDocVersion, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	reqDocVer := &portalModel.DocumentVersionDAOobj{
		DocumentVersion: uuidDocVersion.String(),
		DocID:           reqDoc.DocID,
		Version:         req.Version,
		DocDescription:  req.Description,
		AuthorID:        authorID,
		Price:           req.Price,
	}
	docVer, err := cache.FirstOrCreateDocumentVersion(ctx, core.GetDB(), reqDocVer)
	if err != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	if docVer != "" {
		reqDocVer.DocumentVersion = docVer
	}

	// SAVE BARCODES TO DOCMANAGER
	// get Shard for barcode by documentVersion
	shardID := core.GetShardID(core.GetHash(reqDocVer.DocumentVersion))
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	ser, ok := shardService[uint64(shardID)]
	if !ok {
		fmt.Println("nil shardID")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return
	}

	// create list barcodes
	var (
		saveBarcodes = make([]*docmanagerModel.SaveBarcodeReq, req.Number)
	)
	for i := range saveBarcodes {
		uuidBarcode, er := uuid.NewUUID()
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		saveBarcodes[i] = &docmanagerModel.SaveBarcodeReq{
			Barcode: &docmanagerModel.Barcode{
				ID:     uint64(core.GetHash(uuidBarcode.String())),
				DocVer: reqDocVer.DocumentVersion,
				Status: 0, //FIXME: define barcode status
			},
		}
	}

	// send grpc request
	respBarcodeIDs := make([]uint64, req.Number)
	for i := range saveBarcodes {
		resp, er := ser.Docmanager.SaveBarcode(ctx, saveBarcodes[i])
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}
		respBarcodeIDs = append(respBarcodeIDs, saveBarcodes[i].Barcode.ID)
	}

	return http.StatusOK, respBarcodeIDs, lg, false, nil
}
