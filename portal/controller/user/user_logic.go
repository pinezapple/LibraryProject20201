package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	dao "github.com/pinezapple/LibraryProject20201/portal/dao/database"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
)

func selectAllUser(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select All User", Data: ""}
	db := core.GetDB()
	uDAO := dao.GetUserDAO()

	res, er := uDAO.SelectAllUserFromCache(ctx, db)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, res, lg, false, nil
}

func selectUserByID(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*reqUserID)
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select All User", Data: ""}
	db := core.GetDB()
	uDAO := dao.GetUserDAO()

	res, er := uDAO.Select(ctx, db, req.ID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	fmt.Println(res)

	return http.StatusOK, res, lg, false, nil
}

func saveUser(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*reqUser)
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select All User", Data: ""}
	db := core.GetDB()
	uDAO := dao.GetUserDAO()

	res, er := uDAO.Create(ctx, db, core.GetConfig().WebServer.Secure.SipHashSum0, core.GetConfig().WebServer.Secure.SipHashSum1, req.Username, req.Password, req.Name, req.Role, req.Dob, req.Sex, req.PhoneNumber)
	if er != nil || res == nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}
	return http.StatusOK, nil, lg, false, nil
}

func updateUser(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*reqUser)
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select All User", Data: ""}
	db := core.GetDB()
	uDAO := dao.GetUserDAO()

	er := uDAO.Update(ctx, db, core.GetConfig().WebServer.Secure.SipHashSum0, core.GetConfig().WebServer.Secure.SipHashSum1, req.ID, req.Password, req.Name, req.Role, req.Dob, req.Sex, req.PhoneNumber)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, nil, lg, false, nil
}

func deleteUser(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*reqUserID)
	// Log login info
	lg = &model.LogFormat{Source: c.Request().RemoteAddr, Action: "Select All User", Data: ""}
	db := core.GetDB()
	uDAO := dao.GetUserDAO()

	er := uDAO.Delete(ctx, db, req.Username, req.ID)
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	return http.StatusOK, nil, lg, false, nil
}
