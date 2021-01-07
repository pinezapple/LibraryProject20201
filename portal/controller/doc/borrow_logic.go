package doc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
)

var (
	hours24 = time.Hour * 24
)

func createBorrowForm(c echo.Context, request interface{}) (statusCode int, data interface{}, lg *model.LogFormat, logResponse bool, err error) {
	ctx := c.Request().Context()
	req := request.(*portalModel.CreateBorrowFormReq)

	lg = &model.LogFormat{
		Source: c.Request().RemoteAddr,
		Action: "create borrow form request",
	}

	// gen borrowform ID
	borrowformUUID, er := uuid.NewUUID()
	if er != nil {
		statusCode, err = http.StatusInternalServerError, er
		return
	}

	// create borrowform to save
	rpcBorrowFormReq := docmanagerModel.SaveBorrowFormReq{
		Borrowform: &docmanagerModel.BorrowForm{
			ID:          uint64(core.GetHash(borrowformUUID.String())),
			LibrarianID: req.LibrarianID,
			BarcodeID:   req.Barcodes,
			StartTime:   model.NewTime(time.Now()),
			EndTime:     model.NewTime(time.Now().Add(hours24 * time.Duration(req.BorrowDays))),
		},
	}

	// get borrow form shard
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		fmt.Println("nil shardService")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("nil shardService")
		return
	}
	shardID := core.GetShardID(uint32(rpcBorrowFormReq.Borrowform.ID))
	ser, ok := shardService[uint64(shardID)]
	if !ok {
		fmt.Println("nil shardID")
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("no shard id")
		return
	}

	resp, er := ser.Docmanager.SaveBorrowForm(ctx, &rpcBorrowFormReq)
	if er != nil || resp.Code != 0 {
		statusCode, err = http.StatusInternalServerError, fmt.Errorf("grpc Error")
		return
	}

	return http.StatusOK, nil, lg, false, nil
}
