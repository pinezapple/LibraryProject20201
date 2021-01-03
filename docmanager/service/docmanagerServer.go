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

func (d *docManagerSrv) SelectAllAvailableBarcode(ctx context.Context, req *docmanagerModel.SelectAllAvailableBarcodeReq) (resp *docmanagerModel.SelectAllAvailableBarcodeResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all available barcode")

	//FIXME: add functions
	barcodes, err := dao.SelectAllBarcode(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllAvailableBarcodeResp{Code: 1, Message: err.Error()}, err
	}

	d.lg.Message = "Still using select all barcode"
	logger.LogWarning(d.lg)

	logger.LogInfo(d.lg, "RPC Resp: Select all available barcode OK")
	return &docmanagerModel.SelectAllAvailableBarcodeResp{Code: 0, Barcodes: barcodes}, nil
}

func (d *docManagerSrv) SelectAllSellingBarcode(ctx context.Context, req *docmanagerModel.SelectAllSellingBarcodeReq) (resp *docmanagerModel.SelectAllSellingBarcodeResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all selling barcode")

	sellingBarcodes, err := dao.SelectAllSellingBarcode(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllSellingBarcodeResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select all selling barcode OK")
	return &docmanagerModel.SelectAllSellingBarcodeResp{Code: 0, Barcodes: sellingBarcodes}, nil
}

func (d *docManagerSrv) SelectAllDamageBarcode(ctx context.Context, req *docmanagerModel.SelectAllDamageBarcodeReq) (resp *docmanagerModel.SelectAllDamageBarcodeResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all damage barcode")

	damageBarcodes, err := dao.SelectAllDamagedBarcode(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllDamageBarcodeResp{Code: 1, Message: err.Error()}, err
	}

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
		return &docmanagerModel.SelectAllBorrowFormResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select all borrow form OK")
	return &docmanagerModel.SelectAllBorrowFormResp{Code: 0, BorrowForms: borrowForms}, nil
}

func (d *docManagerSrv) SelectAllUnReturnBorrowForm(ctx context.Context, req *docmanagerModel.SelectAllUnReturnBorrowFormReq) (resp *docmanagerModel.SelectAllUnReturnBorrowFormResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all unreturn borrow form")

	//FIXME: add functions
	unReturnBorrowForms, err := dao.SelectAllUnreturnedBorrowForm(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllUnReturnBorrowFormResp{Code: 1, Message: err.Error()}, err
	}

	d.lg.Message = "Using select all borrow form func"
	logger.LogWarning(d.lg)

	logger.LogInfo(d.lg, "RPC Resp: Select all unreturn borrow form OK")
	return &docmanagerModel.SelectAllUnReturnBorrowFormResp{Code: 0, BorrowForms: unReturnBorrowForms}, nil
}

func (d *docManagerSrv) SelectBorrowFormByID(ctx context.Context, req *docmanagerModel.SelectBorrowFormByIDReq) (resp *docmanagerModel.SelectBorrowFormByIDResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select borrow form by ID")

	borrowForm, err := dao.SelectBorrowFormByID(ctx, core.GetDB(), req.BorrowFormID)
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectBorrowFormByIDResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select borrow form by ID OK")
	return &docmanagerModel.SelectBorrowFormByIDResp{Code: 0, Borrowform: borrowForm}, nil
}

func (d *docManagerSrv) SaveBorrowForm(ctx context.Context, req *docmanagerModel.SaveBorrowFormReq) (resp *docmanagerModel.SaveBorrowFormResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Save borrow form")

	if err = dao.InsertBorrowForm(ctx, core.GetDB(), req.Borrowform); err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SaveBorrowFormResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Save borrow form OK")
	return &docmanagerModel.SaveBorrowFormResp{Code: 0}, nil
}

// ---------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- PAYMENT ----------------------------------------------------------

func (d *docManagerSrv) SelectAllPayment(ctx context.Context, req *docmanagerModel.SelectAllPaymentReq) (resp *docmanagerModel.SelectAllPaymentResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all payment")

	payments, err := dao.SelectAllPayment(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllPaymentResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select all payment OK")
	return &docmanagerModel.SelectAllPaymentResp{Code: 0, Payments: payments}, nil
}

func (d *docManagerSrv) SelectPaymentByID(ctx context.Context, req *docmanagerModel.SelectPaymentByIDReq) (resp *docmanagerModel.SelectPaymentByIDResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select payment by id")

	payment, err := dao.SelectPaymentByID(ctx, core.GetDB(), req.PaymentID)
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectPaymentByIDResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select payment by id OK")
	return &docmanagerModel.SelectPaymentByIDResp{Code: 0, Payment: payment}, nil
}

func (d *docManagerSrv) SelectPaymentByBorrowFormID(ctx context.Context, req *docmanagerModel.SelectPaymentByBorrowFormIDReq) (resp *docmanagerModel.SelectPaymentByBorrowFormIDResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select payment by borrow form ID")

	payment, err := dao.SelectPaymentByBorrowFormID(ctx, core.GetDB(), req.BorrowFormID)
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectPaymentByBorrowFormIDResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select payment by borrow form ID OK")
	return &docmanagerModel.SelectPaymentByBorrowFormIDResp{Code: 0, Payment: payment}, nil
}

func (d *docManagerSrv) SavePayment(ctx context.Context, req *docmanagerModel.SavePaymentReq) (resp *docmanagerModel.SavePaymentResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Save payment")

	if err = dao.InsertPayment(ctx, core.GetDB(), req.Payment); err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SavePaymentResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Save payment OK")
	return &docmanagerModel.SavePaymentResp{Code: 0}, nil
}

// -----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- Sale bill ----------------------------------------------------------

func (d *docManagerSrv) SelectAllSaleBill(ctx context.Context, req *docmanagerModel.SelectAllSaleBillReq) (resp *docmanagerModel.SelectAllSaleBillResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select all sale bill")

	saleBills, err := dao.SelectAllSaleBill(ctx, core.GetDB())
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectAllSaleBillResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select all sale bill OK")
	return &docmanagerModel.SelectAllSaleBillResp{Code: 0, SaleBills: saleBills}, nil
}

func (d *docManagerSrv) SelectSaleBillByID(ctx context.Context, req *docmanagerModel.SelectSaleBillByIDReq) (resp *docmanagerModel.SelectSaleBillByIDResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Select sale bill by ID")

	saleBill, err := dao.SelectSaleBillByID(ctx, core.GetDB(), req.SaleBillID)
	if err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SelectSaleBillByIDResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Select sale bill by ID OK")
	return &docmanagerModel.SelectSaleBillByIDResp{Code: 0, SaleBill: saleBill}, nil
}

func (d *docManagerSrv) SaveSaleBill(ctx context.Context, req *docmanagerModel.SaveSaleBillReq) (resp *docmanagerModel.SaveSaleBillResp, err error) {
	logger.LogInfo(d.lg, "RPC Req: Save sale bill")

	if err = dao.InsertSaleBill(ctx, core.GetDB(), req.SaleBill); err != nil {
		logger.LogErr(d.lg, err)
		return &docmanagerModel.SaveSaleBillResp{Code: 1, Message: err.Error()}, err
	}

	logger.LogInfo(d.lg, "RPC Resp: Save sale bill OK")
	return &docmanagerModel.SaveSaleBillResp{Code: 0}, nil
}
