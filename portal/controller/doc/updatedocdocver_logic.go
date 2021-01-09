package doc

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/dao/cache"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
)

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
