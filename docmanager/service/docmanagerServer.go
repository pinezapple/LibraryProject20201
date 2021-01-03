package service

import (
	"context"

	"github.com/pinezapple/LibraryProject20201/docmanager/core"
	"github.com/pinezapple/LibraryProject20201/docmanager/dao"
	"github.com/pinezapple/LibraryProject20201/skeleton/logger"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

type docManagerSrv struct {
	lg *model.LogFormat
}

// ---------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- BARCODE ----------------------------------------------------------

func (d *docManagerSrv) SelectAllBarcode(ctx context.Context, req *docmanagerModel.SelectAllBarcodeReq) (resp *docmanagerModel.SelectAllBarcodeResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all barcode")

	barcodes, err := dao.SelectAllBarcode(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllBarcodeResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select all barcode OK")
	return &docmanagerModel.SelectAllBarcodeResp{Code: 0, Barcodes: barcodes}, nil
}

func (d *docManagerSrv) SelectAllSellingBarcode(ctx context.Context, req *docmanagerModel.SelectAllSellingBarcodeReq) (resp *docmanagerModel.SelectAllBarcodeResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all selling barcode")

	// FIXME: add function
	sellingBarcodes, err := dao.SelectAllBarcode(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllSellingBarcodeResp{Code: 1, Message: err.Error()}, err
	}

	d.lg.Message = "DAO: Still use select all"
	logger.LogWarning(d.lg)

	logger.LogInfo(d.lg, "RPC Resp: Select all selling barcode OK")
	return &docmanagerModel.SelectAllSellingBarcodeResp{Code: 0, Barcodes: sellingBarcodes}, nil
}

func (d *docManagerSrv) SelectAllDamageBarcode(ctx context.Context, req *docmanagerModel.SelectAllDamageBarcodeReq) (resp *docmanagerModel.SelectAllDamageBarcodeResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all damage barcode")

	// FIXME: add function
	damageBarcodes, err := dao.SelectAllBarcode(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllDamageBarcodeResp{Code: 1, Message: err.Error()}, err
	}

	d.lg.Message = "DAO: Still use select all"
	logger.LogWarning(d.lg)

	logger.LogInfo(d.lg, "RPC Resp: Select all damage barcode OK")
	return &docmanagerModel.SelectAllDamageBarcodeResp{Code: 0, Barcodes: damageBarcodes}, nil
}

func (d *docManagerSrv) SelectBarcodeByID(ctx context.Context, req *docmanagerModel.SelectBarcodeByIDReq) (resp *docmanagerModel.SelectBarcodeByIDResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select barcode by id")

	barcode, err := dao.SelectBarcodeByID(ctx, core.GetDB(), req.BarcodeID)
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectBarcodeByIDResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select barcode by ID OK")
	return &docmanagerModel.SelectBarcodeByIDResp{Code: 0, Barcode: barcode}, nil
}

func (d *docManagerSrv) SaveBarcode(ctx context.Context, req *docmanagerModel.SaveBarcodeReq) (resp *docmanagerModel.SaveBarcodeResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Save barcode")

	if err = dao.InsertBarcode(ctx, core.GetDB(), req.Barcode); err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SaveBarcodeResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Save barcode OK")
	return &docmanagerModel.SaveBarcodeResp{Code: 0}, nil
}

func (d *docManagerSrv) UpdateBarcode(ctx context.Context, req *docmanagerModel.UpdateBarcodeReq) (resp *docmanagerModel.UpdateBarcodeResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Update barcode")

	if req.Barcode != nil && req.Barcode.SaleBillID != 0 {
		if err = dao.UpdateBarcodeSaleBill(ctx, core.GetDB(), req.Barcode.ID, req.Barcode.SaleBillID); err != nil {
			logger.LogErr(d.lg, err)
			return &docmanagerModel.UpdateBarcodeResp{Code: 1, Message: err.Error()}, err
		}
		logger.LogInfo(d.lg, "Update barcode sale bill ID OK")
	}

	if req.Barcode != nil && req.Barcode.Status != 0 {
		if err = dao.UpdateBarcodeStatus(ctx, core.GetDB(), req.Barcode.ID, req.Barcode.Status); err != nil {
			logger.LogErr(d.lg, err)
			return &docmanagerModel.UpdateBarcodeResp{Code: 1, Message: err.Error()}, err
		}
		logger.LogInfo(d.lg, "Update barcode status OK")
	}

	logger.LogInfo(d.lg, "RPC Resp: Update barcode OK")
	return &docmanagerModel.UpdateBarcodeResp{Code: 0}, nil
}

// -------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- Borrow form ----------------------------------------------------------

func (d *docManagerSrv) SelectAllBorrowForm(ctx context.Context, req *docmanagerModel.SelectAllBorrowFormReq) (resp *docmanagerModel.SelectAllBorrowFormResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all borrow form")

	borrowForms, err := dao.SelectAllBorrowForm(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllBorrowFormReq{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select all borrow form OK")
	return &docmanagerModel.SelectAllBorrowFormResp{Code: 0, BorrowForms: borrowForms}, nil
}
