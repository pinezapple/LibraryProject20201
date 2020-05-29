package dao

import (
	"context"
	"time"

	"github.com/linxGnu/mssqlx"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

const (
	sqlSelectAllDocInCache       = "SELECT * FROM doc"
	sqlSelectAllFormInCache      = "SELECT * FROM borrowform"
	sqlSelectFormInCacheByStatus = "SELECT * FROM borrowform WHERE status = ?"
	sqlSaveDocToCache            = "INSERT INTO doc(id_doc,doc_name,doc_author,doc_type,doc_description,status,fee) VALUES (?,?,?,?,?,?,?)"
	sqlDeleteDocToCache          = "DELETE FROM doc WHERE id_doc = ?"
	sqlUpdateStatusDocToCache    = "UPDATE doc SET status = ?, id_borrow = ?,updated_at = ? WHERE id_doc= ?"
	sqlUpdateDocToCache          = "UPDATE doc SET doc_name = ?, doc_author = ?, doc_type =?, doc_description = ?, status = ?, fee = ?, updated_at = ? WHERE id_doc = ?"
	sqlSaveBorrowForm            = "INSERT INTO borrowform(id_borrow, id_doc, id_cus, id_lib, status, start_at, end_at) VALUE (?,?,?,?,?,?,?)"
	sqlUpdateBorrowFormStatus    = "UPDATE borrowform SET status = ?, updated_at = ? WHERE id_borrow = ?"
	sqlSelectBorrowFormByID      = "SELECT * FROM borrowform WHERE id_borrow = ?"
	sqlSelecetIdDoc              = "SELECT id_doc FROM doc WHERE id_borrow = ?"
)

type IDocCache interface {
	// Select all doc from cachedsa
	SelectAllDocFromCache(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Doc, err error)
	SelectAllFormFromCache(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.BorrowForm, err error)
	SaveDoc(ctx context.Context, db *mssqlx.DBs, doc *docmanagerModel.Doc) (err error)
	UpdateDoc(ctx context.Context, db *mssqlx.DBs, doc *docmanagerModel.Doc) (err error)
	DelDoc(ctx context.Context, db *mssqlx.DBs, id uint64) (err error)
	//--------------- BorrowForm --------------
	SaveBorrowForm(ctx context.Context, db *mssqlx.DBs, form *docmanagerModel.BorrowForm) (err error)
	UpdateBorrowFormStatus(ctx context.Context, db *mssqlx.DBs, id uint64, status int) (err error)
	SelectBorrowFormByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *docmanagerModel.BorrowForm, err error)
}

type docCacheDAO struct{}

func (d *docCacheDAO) SelectAllDocFromCache(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Doc, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, result, sqlSelectAllDocInCache)
	return
}

func (d *docCacheDAO) SelectAllFormFromCache(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.BorrowForm, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, result, sqlSelectAllFormInCache)
	return
}

func (d *docCacheDAO) SaveDoc(ctx context.Context, db *mssqlx.DBs, doc *docmanagerModel.Doc) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.ExecContext(ctx, sqlSaveDocToCache, doc.ID, doc.Name, doc.Author, doc.Type, doc.Descriptor, doc.Status, doc.Fee)
	return
}

func (d *docCacheDAO) UpdateDoc(ctx context.Context, db *mssqlx.DBs, doc *docmanagerModel.Doc) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.ExecContext(ctx, sqlUpdateDocToCache, doc.Name, doc.Author, doc.Type, doc.Descriptor, doc.Status, doc.Fee, time.Now(), doc.ID)
	return
}

func (d *docCacheDAO) DelDoc(ctx context.Context, db *mssqlx.DBs, id uint64) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.ExecContext(ctx, sqlDeleteDocToCache, id)
	return
}

func (d *docCacheDAO) UpdateStatus(ctx context.Context, db *mssqlx.DBs, id uint64, status int) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.ExecContext(ctx, sqlUpdateStatusDocToCache, status, id)
	return
}

//-------------------------- Borrow Form ---------------------------
func (d *docCacheDAO) SaveBorrowForm(ctx context.Context, db *mssqlx.DBs, form *docmanagerModel.BorrowForm) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.ExecContext(ctx, sqlSaveBorrowForm, form.ID, form.DocID, form.CusID, form.LibID, form.Status, form.StartAt, form.EndAt)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, sqlUpdateStatusDocToCache, form.Status, form.ID, time.Now(), form.DocID)
	return
}

func (d *docCacheDAO) UpdateBorrowFormStatus(ctx context.Context, db *mssqlx.DBs, id uint64, status int) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}
	_, err = db.ExecContext(ctx, sqlUpdateBorrowFormStatus, status, time.Now(), id)
	if err != nil {
		return err
	}
	var id_doc uint64
	err = db.SelectContext(ctx, &id_doc, sqlSelecetIdDoc, id)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, sqlUpdateStatusDocToCache, status, id, time.Now(), id_doc)
	return
}

func (d *docCacheDAO) SelectBorrowFormByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *docmanagerModel.BorrowForm, err error) {
	if db == nil {
		return nil, core.ErrDBObjNull
	}

	err = db.SelectContext(ctx, result, sqlSelectBorrowFormByID, id)
	return
}

var dcDAO IDocCache = &docCacheDAO{}

func GetDocCacheDAO() IDocCache {
	return dcDAO
}

func SetDocCacheDAO(v IDocCache) {
	dcDAO = v
}
