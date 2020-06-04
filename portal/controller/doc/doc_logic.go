package doc

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	dao "github.com/pinezapple/LibraryProject20201/portal/dao/database"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

func selectAllDoc(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select Doc from cache", Data: ""}
	db := core.GetDB()
	dDAO := dao.GetDocCacheDAO()

	res, er := dDAO.SelectAllDocFromCache(ctx, db)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, res, lg, false, nil
}

func selectAllDoc0(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select Doc from cache", Data: ""}
	db := core.GetDB()
	dDAO := dao.GetDocCacheDAO()

	res, er := dDAO.SelectAllDoc0FromCache(ctx, db)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, res, lg, false, nil
}

func saveDoc(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	req := request.(*reqDoc)
	ctx := c.Request().Context()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Save Doc", Data: req}
	fmt.Println(req)
	// generate ID
	id := core.GetHash(req.Name + req.Author)
	shardID := core.GetShardID(id)

	// save Doc to cache
	db := core.GetDB()
	dDAO := dao.GetDocCacheDAO()
	docObj := &docmanagerModel.Doc{
		ID:          uint64(id),
		Name:        req.Name,
		Author:      req.Author,
		Type:        req.Type,
		Description: req.Description,
		Fee:         req.Fee,
	}
	er := dDAO.SaveDoc(ctx, db, docObj)
	if er != nil {
		fmt.Println(er)
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	// save Doc to docmanager
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

	resp, err := ser.Docmanager.SaveDoc(ctx, &docmanagerModel.SaveDocReq{Doc: docObj})
	if err != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

func delDoc(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	req := request.(*reqSelectByID)
	ctx := c.Request().Context()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Del Doc", Data: req}
	//del doc on cache
	db := core.GetDB()
	dDAO := dao.GetDocCacheDAO()
	er := dDAO.DelDoc(ctx, db, req.ID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	//del doc to docmanager
	shardID := core.GetShardID(uint32(req.ID))
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	ser, ok := shardService[uint64(shardID)]
	if !ok {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return

	}

	resp, err := ser.Docmanager.DeleteDoc(ctx, &docmanagerModel.DeleteDocReq{DocID: req.ID})
	if err != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

func updateDoc(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	req := request.(*reqDoc)
	ctx := c.Request().Context()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Update Doc", Data: req}
	shardID := core.GetShardID(uint32(req.ID))

	// update Doc to cache
	db := core.GetDB()
	dDAO := dao.GetDocCacheDAO()
	docObj := &docmanagerModel.Doc{
		ID:          req.ID,
		Name:        req.Name,
		Author:      req.Author,
		Type:        req.Type,
		Description: req.Description,
		Fee:         req.Fee,
	}

	fmt.Println(docObj)
	er := dDAO.UpdateDoc(ctx, db, docObj)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	//update doc to docmanager
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	ser, ok := shardService[uint64(shardID)]
	if !ok {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return

	}

	fmt.Println("updating 2")
	ctx1 := context.Background()
	resp, err := ser.Docmanager.UpdateDoc(ctx1, &docmanagerModel.UpdateDocReq{Doc: docObj})
	if err != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

func selectDocByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	req := request.(*reqSelectByID)
	ctx := c.Request().Context()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select Doc by ID", Data: req}
	//get shard id
	shardID := core.GetShardID(uint32(req.ID))
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	//select doc
	ser, ok := shardService[uint64(shardID)]
	if !ok {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return

	}

	resp, err := ser.Docmanager.SelectDocByID(ctx, &docmanagerModel.SelectDocByIDReq{DocID: req.ID})
	if err != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, resp.Documents, lg, false, nil
}

func selectAllForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	//select all form from cache
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select Form from cache", Data: ""}
	db := core.GetDB()
	dDAO := dao.GetDocCacheDAO()

	res, er := dDAO.SelectAllFormFromCache(ctx, db)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, res, lg, false, nil
}

func saveForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	req := request.(*reqSaveBorrowForm)
	ctx := c.Request().Context()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Save Form", Data: req}
	// generate ID
	shardID := core.GetShardID(uint32(req.DocID))
	id := core.GetHash(strconv.Itoa(int(req.DocID)) + strconv.Itoa(int(req.CusID)) + strconv.Itoa(rand.Int()))

	// save form to cache
	db := core.GetDB()
	dDAO := dao.GetDocCacheDAO()
	formObj := &docmanagerModel.BorrowForm{
		ID:      uint64(id),
		DocID:   req.DocID,
		CusID:   req.CusID,
		LibID:   req.LibID,
		Status:  req.Status,
		StartAt: model.NewTime(time.Now()),
		EndAt:   model.NewTime(time.Now().Add(time.Duration(req.TTL) * time.Hour * 24)),
	}
	er := dDAO.SaveBorrowForm(ctx, db, formObj)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	// save form to docmanager
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	ser, ok := shardService[uint64(shardID)]
	if !ok {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return

	}

	resp, err := ser.Docmanager.SaveBorrowForm(ctx, &docmanagerModel.SaveBorrowFormReq{Borrowform: formObj})
	if err != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

func selectFormByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	req := request.(*reqSelectFormByID)
	ctx := c.Request().Context()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select Form By ID", Data: req}
	// get shard ID
	shardID := core.GetShardID(uint32(req.DocID))
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	ser, ok := shardService[uint64(shardID)]
	if !ok {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return

	}

	resp, err := ser.Docmanager.SelectBorrowFormByID(ctx, &docmanagerModel.SelectBorrowFormByIDReq{FormID: req.FormID})
	if err != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, resp.Borrowform, lg, false, nil
}

func updateStatus(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	fmt.Println("in update status")
	req := request.(*reqUpdateStatus)
	ctx := c.Request().Context()
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Update status", Data: req}
	//get shard id
	shardID := core.GetShardID(uint32(req.DocID))
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}

	ser, ok := shardService[uint64(shardID)]
	if !ok {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return

	}
	db := core.GetDB()
	dDAO := dao.GetDocCacheDAO()
	//update status on cache
	//update doc status
	//update borrowform status
	er := dDAO.UpdateBorrowFormStatus(ctx, db, req.FormID, req.Status)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	ctx1 := context.Background()
	resp, err := ser.Docmanager.UpdateBorrowFormStatus(ctx1, &docmanagerModel.UpdateBorrowFormStatusReq{FormID: req.FormID, Status: int32(req.Status)})
	if err != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, nil, lg, false, nil
}
