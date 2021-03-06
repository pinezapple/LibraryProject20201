package doc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/dao/cache"
	"github.com/pinezapple/LibraryProject20201/portal/dao/database"
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
			docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcodes[j].DocVerID)
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
				DocID:       doc.DocID,
				Version:     docver.DocumentVersion,
				Author:      aut.AuthorName,
				Publisher:   docver.Publisher,
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
		resp, er := ser.Docmanager.SelectAllSellingBarcode(ctx, &docmanagerModel.SelectAllSellingBarcodeReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		for j := 0; j < len(resp.Barcodes); j++ {
			docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcodes[j].DocVerID)
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
				DocID:       doc.DocID,
				Version:     docver.DocumentVersion,
				Author:      aut.AuthorName,
				Publisher:   docver.Publisher,
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

func selectAllBarcodeAvail(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	db := core.GetDB()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all avail barcode from shards", Data: ""}
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
		resp, er := ser.Docmanager.SelectAllAvailableBarcode(ctx, &docmanagerModel.SelectAllAvailableBarcodeReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		for j := 0; j < len(resp.Barcodes); j++ {
			docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcodes[j].DocVerID)
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
				DocID:       doc.DocID,
				Version:     docver.DocumentVersion,
				Author:      aut.AuthorName,
				Publisher:   docver.Publisher,
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
		resp, er := ser.Docmanager.SelectAllDamageBarcode(ctx, &docmanagerModel.SelectAllDamageBarcodeReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}

		for j := 0; j < len(resp.Barcodes); j++ {
			docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcodes[j].DocVerID)
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
				DocID:       doc.DocID,
				Version:     docver.DocumentVersion,
				Author:      aut.AuthorName,
				Publisher:   docver.Publisher,
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
	db := core.GetDB()
	userDAO := database.GetUserDAO()

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*portalModel.SelectAllSaleBillResp

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
		fmt.Println(resp)
		for j := 0; j < len(resp.SaleBills); j++ {
			var price uint64
			for k := 0; k < len(resp.SaleBills[j].Price); k++ {
				price += resp.SaleBills[j].Price[k]
			}

			//fmt.Println(resp.SaleBills[j].LibrarianID)
			lib, _ := userDAO.Select(ctx, db, resp.SaleBills[j].LibrarianID)

			tmp := &portalModel.SelectAllSaleBillResp{
				SaleBillID:    resp.SaleBills[j].ID,
				LibrarianID:   resp.SaleBills[j].LibrarianID,
				LibrarianName: lib.Name,
				TotalMoney:    price,
				CreatedAt:     resp.SaleBills[j].CreatedAt,
			}

			finalResp = append(finalResp, tmp)
		}
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
	db := core.GetDB()

	userDAO := database.GetUserDAO()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*portalModel.SelectAllPaymentResp

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
		for j := 0; j < len(resp.Payments); j++ {
			var money = resp.Payments[i].Fine
			var stat uint64
			if money > 0 {
				stat = 3
			} else {
				stat = 0
			}
			lib, er := userDAO.Select(ctx, db, resp.Payments[j].LibrarianID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}

			user, er := userDAO.Select(ctx, db, resp.Payments[j].ReaderID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return

			}

			for k := 0; k < len(resp.Payments[j].Money); k++ {
				money += resp.Payments[j].Money[k]
			}

			tmp := &portalModel.SelectAllPaymentResp{
				PaymentID:     resp.Payments[j].ID,
				BorrowFormID:  resp.Payments[j].BorrowFormID,
				LibrarianID:   resp.Payments[j].LibrarianID,
				LibrarianName: lib.Name,
				Status:        stat,
				ReaderID:      resp.Payments[j].ReaderID,
				ReaderName:    user.Name,
				TotalMoney:    money,
				CreatedAt:     resp.Payments[j].CreatedAt,
			}

			finalResp = append(finalResp, tmp)
		}

	}
	data = finalResp

	return
}

func selectPaymentWithFine(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select payment with fine from shards", Data: ""}
	shardNum := core.ShardNumber
	shardService := microservice.GetDocmanagerShardServices()
	db := core.GetDB()

	userDAO := database.GetUserDAO()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*portalModel.SelectAllPaymentResp

	for i := 0; i < shardNum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectPaymentWithFine(ctx, &docmanagerModel.SelectPaymentWithFineReq{})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}
		for j := 0; j < len(resp.Payment); j++ {
			var money = resp.Payment[i].Fine
			lib, er := userDAO.Select(ctx, db, resp.Payment[j].LibrarianID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}

			user, er := userDAO.Select(ctx, db, resp.Payment[j].ReaderID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return

			}

			for k := 0; k < len(resp.Payment[j].Money); k++ {
				money += resp.Payment[j].Money[k]
			}
			tmp := &portalModel.SelectAllPaymentResp{
				PaymentID:     resp.Payment[j].ID,
				BorrowFormID:  resp.Payment[j].BorrowFormID,
				LibrarianID:   resp.Payment[j].LibrarianID,
				LibrarianName: lib.Name,
				ReaderID:      resp.Payment[j].ReaderID,
				ReaderName:    user.Name,
				TotalMoney:    money,
				CreatedAt:     resp.Payment[j].CreatedAt,
			}

			finalResp = append(finalResp, tmp)
		}

	}
	data = finalResp

	return
}

func selectAllBorrowForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all borrow form", Data: ""}
	shardNum := core.ShardNumber
	conf := core.GetConfig()

	db := core.GetDB()
	if db == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("Database connection / Config not initialized")
		return
	}

	// Select user by his username
	userDAO := database.GetUserDAO()

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*portalModel.SelectAllBorrowFormElement

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
		for j := 0; j < len(resp.BorrowForms); j++ {
			lib, er := userDAO.Select(ctx, db, resp.BorrowForms[j].LibrarianID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return
			}

			user, er := userDAO.Select(ctx, db, resp.BorrowForms[j].ReaderID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return

			}

			fine := (time.Now().Unix() - resp.BorrowForms[j].EndTime.Seconds) / 86400 * conf.FinePerDay
			if fine < 0 {
				fine = 0
			}

			tmp := &portalModel.SelectAllBorrowFormElement{
				ID:            resp.BorrowForms[j].ID,
				LibrarianID:   resp.BorrowForms[j].LibrarianID,
				LibrarianName: lib.Name,
				Status:        resp.BorrowForms[j].Status,
				ReaderID:      resp.BorrowForms[j].ReaderID,
				ReaderName:    user.Name,
				BarcodeID:     resp.BorrowForms[j].BarcodeID,
				StartTime:     resp.BorrowForms[j].StartTime,
				EndTime:       resp.BorrowForms[j].EndTime,
				Fine:          fine,
				CreatedAt:     resp.BorrowForms[j].CreatedAt,
				UpdatedAt:     resp.BorrowForms[j].UpdatedAt,
			}

			finalResp = append(finalResp, tmp)
		}

	}
	data = finalResp

	return
}

func selectAllUnreturnedBorrowForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all unreturned borrow form", Data: ""}
	shardNum := core.ShardNumber
	conf := core.GetConfig()
	db := core.GetDB()
	if db == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("Database connection / Config not initialized")
		return
	}
	userDAO := database.GetUserDAO()

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var finalResp []*portalModel.SelectAllBorrowFormElement

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
		for j := 0; j < len(resp.BorrowForms); j++ {
			lib, er := userDAO.Select(ctx, db, resp.BorrowForms[j].LibrarianID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return

			}

			user, er := userDAO.Select(ctx, db, resp.BorrowForms[j].ReaderID)
			if er != nil {
				statusCode, err = http.StatusInternalServerError, er
				return

			}

			fine := (time.Now().Unix() - resp.BorrowForms[j].EndTime.Seconds) / 86400 * conf.FinePerDay
			fmt.Println(time.Now().Unix())
			fmt.Println(resp.BorrowForms[j].EndTime.Seconds)
			fmt.Println(fine)
			if fine <= 0 {
				fine = 0
			} else {
				resp.BorrowForms[j].Status = 2
			}

			tmp := &portalModel.SelectAllBorrowFormElement{
				ID:            resp.BorrowForms[j].ID,
				LibrarianID:   resp.BorrowForms[j].LibrarianID,
				LibrarianName: lib.Name,
				Status:        resp.BorrowForms[j].Status,
				ReaderID:      resp.BorrowForms[j].ReaderID,
				ReaderName:    user.Name,
				BarcodeID:     resp.BorrowForms[j].BarcodeID,
				StartTime:     resp.BorrowForms[j].StartTime,
				EndTime:       resp.BorrowForms[j].EndTime,
				Fine:          fine,
				CreatedAt:     resp.BorrowForms[j].CreatedAt,
				UpdatedAt:     resp.BorrowForms[j].UpdatedAt,
			}

			finalResp = append(finalResp, tmp)
		}

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
	docver, er := cache.SelectDocumentVersionByID(ctx, db, resp.Barcode.DocVerID)
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
		BarcodeID: resp.Barcode.ID,
		Status:    resp.Barcode.Status,
		DocName:   doc.DocName,
		DocID:     doc.DocID,
		// Version:     docver.Version,
		Author:      aut.AuthorName,
		Publisher:   docver.Publisher,
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

func selectBorrowFormByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*portalModel.SelectBorrowFormByIDReq)
	conf := core.GetConfig()
	db := core.GetDB()
	if db == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("Database connection / Config not initialized")
		return
	}
	userDAO := database.GetUserDAO()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select barcode by id", Data: ""}

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	shardId := core.GetShardID(uint32(req.BorrowFormID))
	ser, ok := shardService[uint64(shardId)]
	if !ok {
		fmt.Println("nil shardID")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return
	}

	resp, er := ser.Docmanager.SelectBorrowFormByID(ctx, &docmanagerModel.SelectBorrowFormByIDReq{BorrowFormID: req.BorrowFormID})
	if er != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	fine := (time.Now().Unix() - resp.Borrowform.EndTime.Seconds) / 86400 * conf.FinePerDay
	if fine < 0 {
		fine = 0
	}

	lib, er := userDAO.Select(ctx, db, resp.Borrowform.LibrarianID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return

	}

	user, er := userDAO.Select(ctx, db, resp.Borrowform.ReaderID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return

	}

	var barcode []*portalModel.RespBarcodeOverview

	for i := 0; i < len(resp.Borrowform.BarcodeID); i++ {
		docver, er := cache.SelectDocVerIDFromCacheByBarcode(ctx, db, resp.Borrowform.BarcodeID[i])
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		docversion, er := cache.SelectDocumentVersionByID(ctx, db, docver)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}

		doc, er := cache.SelectDocumentByID(ctx, db, docversion.DocID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		aut, er := cache.SelectAuthorByID(ctx, db, docversion.AuthorID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return

		}

		tmp := &portalModel.RespBarcodeOverview{
			BarcodeID: resp.Borrowform.BarcodeID[i],
			DocName:   doc.DocName,
			Author:    aut.AuthorName,
		}

		barcode = append(barcode, tmp)
	}
	// TODO: filter more information
	data = &portalModel.SelectBorrowFormByIDResp{
		BorrowFormID:  resp.Borrowform.ID,
		LibrarianID:   resp.Borrowform.LibrarianID,
		LibrarianName: lib.Name,
		ReaderID:      resp.Borrowform.ReaderID,
		ReaderName:    user.Name,
		Fine:          fine,
		Status:        resp.Borrowform.Status,
		Barcodes:      barcode,
		StartTime:     resp.Borrowform.StartTime,
		EndTime:       resp.Borrowform.EndTime,
	}
	return
}

func selectPaymentByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*portalModel.SelectPaymentByIDReq)
	db := core.GetDB()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select payment by id", Data: ""}

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	shardId := core.GetShardID(uint32(req.BorrowFormID))
	ser, ok := shardService[uint64(shardId)]
	if !ok {
		fmt.Println("nil shardID")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return
	}
	resp, er := ser.Docmanager.SelectPaymentByID(ctx, &docmanagerModel.SelectPaymentByIDReq{PaymentID: req.PaymentID})
	if er != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	var barcodes []*portalModel.RespBarcodePaymentOverview
	var totalMoney = resp.Payment.Fine
	var stat uint64
	if totalMoney > 0 {
		stat = 3
	} else {
		stat = 0
	}

	for i := 0; i < len(resp.Payment.BarcodeID); i++ {
		docver, er := cache.SelectDocVerIDFromCacheByBarcode(ctx, db, resp.Payment.BarcodeID[i])
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		docversion, er := cache.SelectDocumentVersionByID(ctx, db, docver)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}

		doc, er := cache.SelectDocumentByID(ctx, db, docversion.DocID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		aut, er := cache.SelectAuthorByID(ctx, db, docversion.AuthorID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}

		tmp := &portalModel.RespBarcodePaymentOverview{
			BarcodeID:  resp.Payment.BarcodeID[i],
			Status:     resp.Payment.BarcodeStatus[i],
			DocName:    doc.DocName,
			AuthorName: aut.AuthorName,
			Money:      resp.Payment.Money[i],
		}

		totalMoney += resp.Payment.Money[i]
		barcodes = append(barcodes, tmp)
	}
	data = &portalModel.SelectPaymentByIDResp{
		PaymentID:    resp.Payment.ID,
		BorrowFormID: resp.Payment.BorrowFormID,
		ReaderID:     resp.Payment.ReaderID,
		LibrarianID:  resp.Payment.LibrarianID,
		TotalMoney:   totalMoney,
		Fine:         resp.Payment.Fine,
		Status:       stat,
		Barcodes:     barcodes,
		CreatedAt:    resp.Payment.CreatedAt,
	}

	return
}

func selectSaleBillByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*portalModel.SelectSaleBillByIDReq)
	db := core.GetDB()
	userDAO := database.GetUserDAO()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select sale bill by id", Data: ""}

	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	shardId := core.GetShardID(uint32(req.SaleBillID))
	ser, ok := shardService[uint64(shardId)]
	if !ok {
		fmt.Println("nil shardID")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return
	}
	resp, er := ser.Docmanager.SelectSaleBillByID(ctx, &docmanagerModel.SelectSaleBillByIDReq{SaleBillID: req.SaleBillID})
	if er != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	lib, er := userDAO.Select(ctx, db, resp.SaleBill.LibrarianID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return

	}
	var barcodes []*portalModel.SaleBillDetail
	var totalMoney uint64

	for i := 0; i < len(resp.SaleBill.BarcodeID); i++ {
		docver, er := cache.SelectDocVerIDFromCacheByBarcode(ctx, db, resp.SaleBill.BarcodeID[i])
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		docversion, er := cache.SelectDocumentVersionByID(ctx, db, docver)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}

		doc, er := cache.SelectDocumentByID(ctx, db, docversion.DocID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		tmp := &portalModel.SaleBillDetail{
			BarcodeID: resp.SaleBill.BarcodeID[i],
			DocName:   doc.DocName,
			Money:     resp.SaleBill.Price[i],
		}

		totalMoney += resp.SaleBill.Price[i]
		barcodes = append(barcodes, tmp)
	}

	data = &portalModel.SelectSaleBillByIDResp{
		SaleBillID:    resp.SaleBill.ID,
		LibrarianID:   resp.SaleBill.LibrarianID,
		LibrarianName: lib.Name,
		TotalMoney:    totalMoney,
		Barcodes:      barcodes,
	}

	return
}

func saveDocumentByBatch(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	fmt.Println("in saveDocumentByBatch")
	req := request.(*portalModel.SaveDocReq)
	ctx := c.Request().Context()
	httpResp := &portalModel.SaveDocResp{
		Barcodes: make([]uint64, req.Number),
	}
	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "saveDocumentByBatch",
		Data:   req,
	}

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

	// search category ID
	catUUID, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	catID, er := cache.FirstOrCreateCategory(ctx, core.GetDB(), req.Category, uint64(core.GetHash(catUUID.String())))
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
		DocVerID:        uint64(core.GetHash(uuidDocVersion.String())),
		DocumentVersion: req.Version,
		DocID:           reqDoc.DocID,
		Publisher:       req.Publisher,
		DocDescription:  req.Description,
		AuthorID:        authorID,
		Price:           req.Price,
	}
	docVerID, er := cache.FirstOrCreateDocumentVersion(ctx, core.GetDB(), reqDocVer)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	fmt.Println(docVerID)

	fmt.Println(req.Number)

	if docVerID != 0 {
		reqDocVer.DocVerID = docVerID
	}

	// SAVE BARCODES TO DOCMANAGER
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
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
				ID:       uint64(core.GetHash(uuidBarcode.String())),
				DocVerID: docVerID,
				Status:   model.BarcodeNormalStatus,
			},
		}

		shardID := core.GetShardID(uint32(saveBarcodes[i].Barcode.ID))
		ser, ok := shardService[uint64(shardID)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SaveBarcode(ctx, saveBarcodes[i])
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}
		httpResp.Barcodes[i] = saveBarcodes[i].Barcode.ID

		// Save to cache: Barcode - DocVer
		if er := cache.SaveDocverIDToCache(ctx, core.GetDB(), saveBarcodes[i].Barcode.ID, saveBarcodes[i].Barcode.DocVerID); er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
	}

	return http.StatusOK, httpResp, lg, false, nil
}

func selectAllFromBlackList(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	db := core.GetDB()
	userDAO := database.GetUserDAO()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all from blacklist", Data: ""}
	var result []*portalModel.BlackListSelectAllElement

	user, er := userDAO.SelectAllUserFromCache(ctx, db)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	for i := 0; i < len(user); i++ {
		f, er := cache.SelectFromBlackListByUserID(ctx, db, user[i].ID)
		if er != nil {
			fmt.Println(er)
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		if len(f) > 0 {
			var money uint64
			for j := 0; j < len(f); j++ {
				money += f[j].Money
			}

			tmp := &portalModel.BlackListSelectAllElement{
				UserID:   user[i].ID,
				Username: user[i].Name,
				Count:    len(f),
				Money:    money,
			}
			result = append(result, tmp)
		} else {
			continue
		}
	}

	data = result
	return http.StatusOK, data, lg, false, nil
}

func selectBlackListByUserID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*portalModel.SelectByUserIDReq)
	db := core.GetDB()
	userDAO := database.GetUserDAO()

	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select from blacklist by id", Data: ""}

	f, er := cache.SelectFromBlackListByUserID(ctx, db, req.UserID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	user, er := userDAO.Select(ctx, db, req.UserID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	var money uint64
	for j := 0; j < len(f); j++ {
		money += f[j].Money
	}

	result := &portalModel.BlackListSelectByIDResp{
		UserID:   req.UserID,
		Username: user.Name,
		Count:    len(f),
		Money:    money,
		Detail:   f,
	}

	data = result
	return http.StatusOK, data, lg, false, nil
}

func selectAllDoc(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	db := core.GetDB()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all from doc", Data: ""}

	doc, er := cache.SelectAllDocument(ctx, db)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	var result []*portalModel.SelectAllDocumentElement

	for i := 0; i < len(doc); i++ {
		cat, er := cache.SelectCategoriesByID(ctx, db, doc[i].CategoryID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		tmp := &portalModel.SelectAllDocumentElement{
			DocID:        doc[i].DocID,
			DocName:      doc[i].DocName,
			CategoryName: cat.CategoryName,
			CreatedAt:    doc[i].CreatedAt,
		}
		result = append(result, tmp)
	}
	data = result

	return http.StatusOK, data, lg, false, nil
}

func selectDocByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	db := core.GetDB()
	req := request.(*portalModel.SelectDocByIDReq)
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all from doc", Data: ""}

	doc, er := cache.SelectDocumentByID(ctx, db, req.DocID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	cat, er := cache.SelectCategoriesByID(ctx, db, doc.CategoryID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	docvers, er := cache.SelectDocumentVersionByDocumentID(ctx, db, doc.DocID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	var dvs []*portalModel.DocverOverviewElement
	for i := 0; i < len(docvers); i++ {
		barcodeCount, er := cache.SelectCountBarcodeByDocverID(ctx, db, docvers[i].DocVerID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		aut, er := cache.SelectAuthorByID(ctx, db, docvers[i].AuthorID)
		if er != nil {
			statusCode, err = http.StatusInternalServerError, er
			return
		}
		tmp := &portalModel.DocverOverviewElement{
			DocVerID:   docvers[i].DocVerID,
			DocVerName: docvers[i].DocumentVersion,
			Publisher:  docvers[i].Publisher,
			AuthorName: aut.AuthorName,
			Count:      barcodeCount,
			Price:      docvers[i].Price,
		}
		dvs = append(dvs, tmp)
	}

	result := &portalModel.SelectDocumentByID{
		DocID:        req.DocID,
		DocName:      doc.DocName,
		CategoryName: cat.CategoryName,
		CreatedAt:    doc.CreatedAt,
		Docver:       dvs,
	}

	data = result
	return http.StatusOK, data, lg, false, nil
}

func selectDocVerByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	db := core.GetDB()
	shardnum := core.ShardNumber
	req := request.(*portalModel.SelectDocVerByIDReq)
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select all barcode by docver ID", Data: ""}

	docver, er := cache.SelectDocumentVersionByID(ctx, db, req.DocVerID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	doc, er := cache.SelectDocumentByID(ctx, db, docver.DocID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	cat, er := cache.SelectCategoriesByID(ctx, db, doc.CategoryID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	aut, er := cache.SelectAuthorByID(ctx, db, docver.AuthorID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	barcodeCount, er := cache.SelectCountBarcodeByDocverID(ctx, db, docver.DocVerID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	var barcodes []*docmanagerModel.Barcode
	for i := 0; i < shardnum; i++ {
		ser, ok := shardService[uint64(i)]
		if !ok {
			fmt.Println("nil shardID")
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
			return
		}
		resp, er := ser.Docmanager.SelectAllBarcodeByDocVerID(ctx, &docmanagerModel.SelectAllBarcodeByDocVerIDReq{DocVerID: req.DocVerID})
		if er != nil || resp.Code != 0 {
			statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
			return
		}
		barcodes = append(barcodes, resp.Barcode...)
	}

	result := &portalModel.SelectDocVerByIDResp{
		DocID:        doc.DocID,
		DocName:      doc.DocName,
		CategoryName: cat.CategoryName,
		DocVerID:     req.DocVerID,
		DocVerName:   docver.DocumentVersion,
		Publisher:    docver.Publisher,
		AuthorName:   aut.AuthorName,
		Count:        barcodeCount,
		Price:        docver.Price,
		Barcode:      barcodes,
		CreatedAt:    doc.CreatedAt,
	}

	data = result
	return http.StatusOK, data, lg, false, nil
}
