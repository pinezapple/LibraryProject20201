package doc

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	dao "github.com/pinezapple/LibraryProject20201/portal/dao/database"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
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

func saveDoc(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	// generate ID
	// save Doc to cache
	// save Doc to docmanager
	return
}

func delDoc(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	//del doc on cache
	//del doc to docmanager
	return
}

func updateDoc(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	//update doc on cache
	//update doc to docmanager
	return
}

func selectDocByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	//get shard id
	//select doc
	return
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
	//get shard id
	//save to cache
	//save to docmanager
	return
}

func selectFormByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	//get doc id from cache
	//get shard id by doc_id
	//get from from docmanager
	return
}

func updateStatus(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	//update status on cache
	//update doc status
	//update borrowform status
	//get shard id
	//update status on docmanager
	//update doc status
	//update borrowform status

	return
}
