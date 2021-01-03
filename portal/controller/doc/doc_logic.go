package doc

import (
	"fmt"
	"net/http"

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
	/*
		authorID, err := dao.GetDocDAO.GetAuthorID(req.Author)
		if err != nil {

		}
	*/

	// search category ID
	/*
		catID, err := dao.GetDocDAO.GetCategoryID(req.Category)
		if err != nil {

		}
	*/
	// check if exist document, then create
	/*
		reqDoc := &docmanagerModel.Doc{
			Name: req.DocName,
			CategoryId: catID,
		}

		docID, err := dao.GetDocDAO().ExistOrCreateDocument(ctx, reqDoc)
		if err != nil {

		}
	*/

	// check if exist document version, then create
	/*
		reqDocVer := &docmanagerModel.DocVersion{
			DocID: docID,
			DocVer: req.Version,
			Price: req.Price,
			Publisher: req.Publisher,
		}
		docVer, err := dao.GetDocDAO.ExistOrCreateDocumentVersion(ctx, reqDocVer)
		if err != nil {

		}
	*/

	// create barcode
	/*
		var (
			saveBarcodes = make([]docmanagerModel.Barcode, req.Number)
		)
		BarcodesID := sth.CreateBarcode(req.Number)
		for i := range BarcodesID {
			saveBarcodes[i] = &docmanagerModel.Barcode{
			ID: BarcodesID[i],
			DocVer: docVer,
			Status: statusTrongKho,
		}

		for i := range saveBarcodes {
			getShard()
			send saveBarcodes rpc
		}


	*/

	return http.StatusOK, nil, lg, false, nil
}
