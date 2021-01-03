package dao

import (
	"context"
	"encoding/json"

	"github.com/linxGnu/mssqlx"
	"github.com/pinezapple/LibraryProject20201/docmanager/core"
	"github.com/pinezapple/LibraryProject20201/docmanager/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

const (
	sqlSelectAllBarcode          = "SELECT * FROM barcodes"
	sqlSelectAllAvailableBarcode = "SELECT * FROM barcodes WHERE status = 0"
	sqlSelectAllDamagedBarcode   = "SELECT * FROM barcodes WHERE status = 3"
	sqlSelectAllSellingBarcode   = "SELECT * FROM barcodes WHERE status = 4"
	sqlSelectBarcodeByID         = "SELECT * FROM barcodes WHERE barcode_id = ?"
	sqlInsertNewBarcode          = "INSERT INTO barcodes(barcode_id, document_version, status) VALUES (?,?,?,?)"
	sqlUpdateBarcodeSaleBillID   = "UPDATE barcodes SET sale_bill_id = ? WHERE barcode_id = ?"
	sqlUpdateBarcodeStatus       = "UPDATE barcodes SET status = ? WHERE barcode_id = ?"

	sqlSelectAllSaleBill  = "SELECT * FROM sale_bill"
	sqlSelectSaleBillByID = "SELECT * FROM sale_bill WHERE sale_bill_id = ?"
	sqlInsertSaleBill     = "INSERT INTO sale_bill(sale_bill_id, barcode_id, sale_price) VALUES (?,?,?)"

	sqlSelectAllBorrowForm           = "SELECT * FROM borrow_form"
	sqlSelectAllUnreturnedBorrowForm = "SELECT * FROM borrow_form WHERE status <> 3"
	sqlSelectBorrowFormByID          = "SELECT * FROM borrow_form WHERE borrow_form_id = ?"
	sqlInsertBorrowForm              = "INSERT INTO borrow_form(borrow_form_id, librarian_id, barcode_id, status, borrow_start_time, borrow_end_time) VALUES (?,?,?,?,?,?)"
	sqlUpdateBorrowFormStatus        = "UPDATE borrow_form SET status = ? WHERE borrow_form_id = ?"

	sqlSelectAllPayment            = "SELECT * FROM payments"
	sqlSelectPaymentByID           = "SELECT * FROM payments WHERE payments_id = ?"
	sqlInsertPayment               = "INSERT INTO payments(payments_id, borrow_form_id, barcode_id, barcode_status, money) VALUES (?,?,?,?,?)"
	sqlSelectPaymentByBorrowFormID = "SELECT * FROM payments WHERE borrow_form_id = ?"
)

type IDocDAO interface {
	SelectAllBarcode(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Barcode, err error)
	SelectAllBarcodeAvail(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Barcode, err error)
	SelectBarcodeByID(ctx context.Context, db *mssqlx.DBs, barcodeID uint64) (result *docmanagerModel.Barcode, err error)
	InsertBarcode(ctx context.Context, db *mssqlx.DBs, barcode *docmanagerModel.Barcode) (err error)
	UpdateBarcodeSaleBill(ctx context.Context, db *mssqlx.DBs, barcodeID, saleBillID uint64) (err error)
	UpdateBarcodeStatus(ctx context.Context, db *mssqlx.DBs, barcodeID uint64, status int) (err error)

	SelectAllSaleBill(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.SaleBill, err error)
	SelectSaleBillByID(ctx context.Context, db *mssqlx.DBs, saleBillID uint64) (result *docmanagerModel.SaleBill, err error)
	InsertSaleBill(ctx context.Context, db *mssqlx.DBs, saleBill *docmanagerModel.SaleBill) (err error)

	SelectAllBorrowForm(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.BorrowForm, err error)
	SelectBorrowFormByID(ctx context.Context, db *mssqlx.DBs, borrowFormID uint64) (result *docmanagerModel.BorrowForm, err error)
	InsertBorrowForm(ctx context.Context, db *mssqlx.DBs, borrowForm *docmanagerModel.BorrowForm) (err error)
	UpdateBorrowFormStatus(ctx context.Context, db *mssqlx.DBs, borrowFormID uint64, status int) (err error)

	SelectAllPayment(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Payment, err error)
	SelectPaymentByID(ctx context.Context, db *mssqlx.DBs, paymentsID uint64) (result *docmanagerModel.Payment, err error)
	InsertPayment(ctx context.Context, db *mssqlx.DBs, payment *docmanagerModel.Payment) (err error)
	SelectPaymentByBorrowFormID(ctx context.Context, db *mssqlx.DBs, borrowFormID uint64) (result *docmanagerModel.Payment, err error)
}

type docDAO struct {
}

func SelectAllBarcode(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Barcode, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllBarcode)
	return
}

func SelectAllAvailableBarcode(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Barcode, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllAvailableBarcode)
	return
}

func SelectAllDamagedBarcode(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Barcode, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllDamagedBarcode)
	return
}

func SelectAllSellingBarcode(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Barcode, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllSellingBarcode)
	return
}

func SelectBarcodeByID(ctx context.Context, db *mssqlx.DBs, barcodeID uint64) (result *docmanagerModel.Barcode, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	result = &docmanagerModel.Barcode{}
	err = db.GetContext(ctx, result, sqlSelectBarcodeByID, barcodeID)
	return
}

func InsertBarcode(ctx context.Context, db *mssqlx.DBs, barcode *docmanagerModel.Barcode) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.ExecContext(ctx, sqlInsertNewBarcode, barcode.ID, barcode.DocVer, barcode.Status)
	return

}

func UpdateBarcodeSaleBill(ctx context.Context, db *mssqlx.DBs, barcodeID, saleBillID uint64) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.Exec(sqlUpdateBarcodeSaleBillID, saleBillID, barcodeID)
	return
}

func UpdateBarcodeStatus(ctx context.Context, db *mssqlx.DBs, barcodeID uint64, status uint64) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.Exec(sqlUpdateBarcodeStatus, status, barcodeID)
	return

}

func SelectAllSaleBill(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.SaleBill, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	var tempResp []*model.SaleBillDAOobj
	err = db.SelectContext(ctx, &tempResp, sqlSelectAllSaleBill)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tempResp); i++ {
		var barcode, price []uint64
		err = json.Unmarshal(tempResp[i].BarcodeId, &barcode)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(tempResp[i].Price, &price)
		if err != nil {
			return nil, err
		}

		tmp := &docmanagerModel.SaleBill{
			ID:        tempResp[i].ID,
			BarcodeID: barcode,
			Price:     price,
			CreatedAt: tempResp[i].CreatedAt,
		}
		result = append(result, tmp)
	}
	return
}

func SelectSaleBillByID(ctx context.Context, db *mssqlx.DBs, saleBillID uint64) (result *docmanagerModel.SaleBill, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	var tempResp = &model.SaleBillDAOobj{}

	err = db.GetContext(ctx, tempResp, sqlSelectSaleBillByID, saleBillID)
	if err != nil {
		return nil, err
	}

	var barcode, price []uint64
	err = json.Unmarshal(tempResp.BarcodeId, &barcode)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tempResp.Price, &price)
	if err != nil {
		return nil, err
	}

	result = &docmanagerModel.SaleBill{
		ID:        tempResp.ID,
		BarcodeID: barcode,
		Price:     price,
		CreatedAt: tempResp.CreatedAt,
	}

	return
}

func InsertSaleBill(ctx context.Context, db *mssqlx.DBs, saleBill *docmanagerModel.SaleBill) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}
	barcode, err := json.Marshal(saleBill.BarcodeID)
	if err != nil {
		return
	}
	price, err := json.Marshal(saleBill.Price)
	if err != nil {
		return
	}

	_, err = db.Exec(sqlInsertPayment, saleBill.ID, barcode, price)
	return
}

func SelectAllBorrowForm(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.BorrowForm, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	var tempResp []*model.BorrowFormDAOobj
	err = db.SelectContext(ctx, &tempResp, sqlSelectAllBorrowForm)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tempResp); i++ {
		var barcode []uint64
		err = json.Unmarshal(tempResp[i].BarcodeID, &barcode)
		if err != nil {
			return nil, err
		}

		tmp := &docmanagerModel.BorrowForm{
			ID:          tempResp[i].ID,
			LibrarianID: tempResp[i].LibrarianID,
			Status:      tempResp[i].Status,
			BarcodeID:   barcode,
			StartTime:   tempResp[i].StartTime,
			EndTime:     tempResp[i].EndTime,
			CreatedAt:   tempResp[i].CreatedAt,
			UpdatedAt:   tempResp[i].UpdatedAt,
		}
		result = append(result, tmp)
	}
	return

}

func SelectAllUnreturnedBorrowForm(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.BorrowForm, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	var tempResp []*model.BorrowFormDAOobj
	err = db.SelectContext(ctx, &tempResp, sqlSelectAllUnreturnedBorrowForm)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tempResp); i++ {
		var barcode []uint64
		err = json.Unmarshal(tempResp[i].BarcodeID, &barcode)
		if err != nil {
			return nil, err
		}

		tmp := &docmanagerModel.BorrowForm{
			ID:          tempResp[i].ID,
			LibrarianID: tempResp[i].LibrarianID,
			Status:      tempResp[i].Status,
			BarcodeID:   barcode,
			StartTime:   tempResp[i].StartTime,
			EndTime:     tempResp[i].EndTime,
			CreatedAt:   tempResp[i].CreatedAt,
			UpdatedAt:   tempResp[i].UpdatedAt,
		}
		result = append(result, tmp)
	}
	return

}

func SelectBorrowFormByID(ctx context.Context, db *mssqlx.DBs, borrowFormID uint64) (result *docmanagerModel.BorrowForm, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	var tempResp = &model.BorrowFormDAOobj{}

	err = db.GetContext(ctx, tempResp, sqlSelectBorrowFormByID, borrowFormID)
	if err != nil {
		return nil, err
	}

	var barcode []uint64
	err = json.Unmarshal(tempResp.BarcodeID, &barcode)
	if err != nil {
		return nil, err
	}

	result = &docmanagerModel.BorrowForm{
		ID:          tempResp.ID,
		LibrarianID: tempResp.LibrarianID,
		Status:      tempResp.Status,
		BarcodeID:   barcode,
		StartTime:   tempResp.StartTime,
		EndTime:     tempResp.EndTime,
		CreatedAt:   tempResp.CreatedAt,
		UpdatedAt:   tempResp.UpdatedAt,
	}

	return
}

func InsertBorrowForm(ctx context.Context, db *mssqlx.DBs, borrowForm *docmanagerModel.BorrowForm) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	barcode, err := json.Marshal(borrowForm.BarcodeID)
	if err != nil {
		return
	}

	_, err = db.Exec(sqlInsertPayment, borrowForm.ID, borrowForm.LibrarianID, borrowForm.Status, barcode, borrowForm.StartTime, borrowForm.EndTime)

	return
}

func UpdateBorrowFormStatus(ctx context.Context, db *mssqlx.DBs, borrowFormID uint64, status int) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.Exec(sqlUpdateBorrowFormStatus, status, borrowFormID)
	return
}

func SelectAllPayment(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Payment, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	var tempResp []*model.PaymentDAOobj
	err = db.SelectContext(ctx, &tempResp, sqlSelectAllPayment)
	for i := 0; i < len(tempResp); i++ {
		var barcode, barcodestatus, price []uint64

		err = json.Unmarshal(tempResp[i].BarcodeID, &barcode)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(tempResp[i].BarcodeStatus, &barcodestatus)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(tempResp[i].Money, &price)
		if err != nil {
			return nil, err
		}
		tmp := &docmanagerModel.Payment{
			ID:            tempResp[i].ID,
			BorrowFormID:  tempResp[i].BorrowFormID,
			BarcodeID:     barcode,
			BarcodeStatus: barcodestatus,
			Money:         price,
			CreatedAt:     tempResp[i].CreatedAt,
			UpdatedAt:     tempResp[i].UpdatedAt,
		}
		result = append(result, tmp)
	}
	return
}

func SelectPaymentByID(ctx context.Context, db *mssqlx.DBs, paymentsID uint64) (result *docmanagerModel.Payment, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	var tempResp = &model.PaymentDAOobj{}

	err = db.GetContext(ctx, tempResp, sqlSelectPaymentByID, paymentsID)
	if err != nil {
		return nil, err
	}

	var barcode, barcodestatus, price []uint64
	err = json.Unmarshal(tempResp.BarcodeID, &barcode)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tempResp.BarcodeStatus, &barcodestatus)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tempResp.Money, &price)
	if err != nil {
		return nil, err
	}

	result = &docmanagerModel.Payment{
		ID:            tempResp.ID,
		BorrowFormID:  tempResp.BorrowFormID,
		BarcodeID:     barcode,
		BarcodeStatus: barcodestatus,
		Money:         price,
		CreatedAt:     tempResp.CreatedAt,
		UpdatedAt:     tempResp.UpdatedAt,
	}

	return
}

func SelectPaymentByBorrowFormID(ctx context.Context, db *mssqlx.DBs, borrowFormID uint64) (result *docmanagerModel.Payment, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	var tempResp = &model.PaymentDAOobj{}

	err = db.GetContext(ctx, tempResp, sqlSelectPaymentByBorrowFormID, borrowFormID)
	if err != nil {
		return nil, err
	}

	var barcode, barcodestatus, price []uint64
	err = json.Unmarshal(tempResp.BarcodeID, &barcode)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tempResp.BarcodeStatus, &barcodestatus)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tempResp.Money, &price)
	if err != nil {
		return nil, err
	}

	result = &docmanagerModel.Payment{
		ID:            tempResp.ID,
		BorrowFormID:  tempResp.BorrowFormID,
		BarcodeID:     barcode,
		BarcodeStatus: barcodestatus,
		Money:         price,
		CreatedAt:     tempResp.CreatedAt,
		UpdatedAt:     tempResp.UpdatedAt,
	}

	return
}

func InsertPayment(ctx context.Context, db *mssqlx.DBs, payment *docmanagerModel.Payment) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	barcode, err := json.Marshal(payment.BarcodeID)
	if err != nil {
		return
	}
	barcodestatus, err := json.Marshal(payment.BarcodeStatus)
	if err != nil {
		return
	}
	money, err := json.Marshal(payment.Money)
	if err != nil {
		return
	}

	_, err = db.Exec(sqlInsertPayment, payment.ID, payment.BorrowFormID, barcode, barcodestatus, money)

	return
}
