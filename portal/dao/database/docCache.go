package database

import (
	"context"
	"fmt"
	"time"

	"github.com/linxGnu/mssqlx"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
	"github.com/pinezapple/LibraryProject20201/skeleton/model/docmanagerModel"
)

const (
	sqlSelectAllDocInCache       = "SELECT * FROM doc"
	sqlSelectAllDoc0InCache      = "SELECT *FROM doc WHERE status = 0"
	sqlSelectAllFormInCache      = "select br.id_borrow, br.id_doc, d.doc_name, br.id_cus, br.id_lib, br.status, br.start_at, br.end_at from borrowform as br join doc as d where br.id_doc=d.id_doc"
	sqlSelectFormInCacheByStatus = "SELECT * FROM borrowform WHERE status = ?"
	sqlSaveDocToCache            = "INSERT INTO doc(id_doc, doc_name, doc_author, doc_type, doc_description, fee) VALUES (?,?,?,?,?,?)"
	sqlDeleteDocToCache          = "DELETE FROM doc WHERE id_doc = ?"
	sqlUpdateStatusDocToCache    = "UPDATE doc SET status = ?, id_borrow = ? WHERE id_doc= ?"
	sqlUpdateDocToCache          = "UPDATE doc SET doc_name = ?, doc_author = ?, doc_type =?, doc_description = ?, fee = ?, updated_at = ? WHERE id_doc = ?"
	sqlSaveBorrowForm            = "INSERT INTO borrowform(id_borrow, id_doc, id_cus, id_lib, status, start_at, end_at) VALUE (?,?,?,?,?,?,?)"
	sqlUpdateBorrowFormStatus    = "UPDATE borrowform SET status = ?, updated_at = ? WHERE id_borrow = ?"
	sqlUpdateDocStatus           = "UPDATE doc SET status = ? WHERE id_doc = ?"
	sqlSelectBorrowFormByID      = "SELECT * FROM borrowform WHERE id_borrow = ?"
	sqlSelecetIdDoc              = "SELECT id_doc FROM doc WHERE id_borrow = ?"
)

type CacheBorrowForm struct {
	ID        uint64      `protobuf:"varint,1,opt,name=ID,proto3" json:"id_borrow" db:"id_borrow"`
	DocID     uint64      `protobuf:"varint,2,opt,name=DocID,proto3" json:"id_doc" db:"id_doc"`
	Name      string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"doc_name" db:"doc_name"`
	CusID     uint64      `protobuf:"varint,3,opt,name=CusID,proto3" json:"id_cus" db:"id_cus"`
	LibID     uint64      `protobuf:"varint,4,opt,name=LibID,proto3" json:"id_lib" db:"id_lib"`
	Status    int32       `protobuf:"varint,5,opt,name=Status,proto3" json:"status" db:"status"`
	StartAt   *model.Time `protobuf:"bytes,6,opt,name=Start_at,json=StartAt,proto3" json:"start_at" db:"start_at"`
	EndAt     *model.Time `protobuf:"bytes,7,opt,name=End_at,json=EndAt,proto3" json:"end_at" db:"end_at"`
	CreatedAt *model.Time `protobuf:"bytes,8,opt,name=Created_at,json=CreatedAt,proto3" json:"created_at" db:"created_at"`
	UpdatedAt *model.Time `protobuf:"bytes,9,opt,name=Updated_at,json=UpdatedAt,proto3" json:"updated_at" db:"updated_at"`
}

type IDocCache interface {
	// Select all doc from cachedsa
	SelectAllDocFromCache(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Doc, err error)
	SelectAllDoc0FromCache(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Doc, err error)
	SelectAllFormFromCache(ctx context.Context, db *mssqlx.DBs) (result []*CacheBorrowForm, err error)
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

	err = db.SelectContext(ctx, &result, sqlSelectAllDocInCache)
	return
}

func (d *docCacheDAO) SelectAllDoc0FromCache(ctx context.Context, db *mssqlx.DBs) (result []*docmanagerModel.Doc, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllDoc0InCache)
	return
}

func (d *docCacheDAO) SelectAllFormFromCache(ctx context.Context, db *mssqlx.DBs) (result []*CacheBorrowForm, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllFormInCache)
	return
}

func (d *docCacheDAO) SaveDoc(ctx context.Context, db *mssqlx.DBs, doc *docmanagerModel.Doc) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	_, err = db.ExecContext(ctx, sqlSaveDocToCache, doc.ID, doc.Name, doc.Author, doc.Type, doc.Description, doc.Fee)
	return
}

func (d *docCacheDAO) UpdateDoc(ctx context.Context, db *mssqlx.DBs, doc *docmanagerModel.Doc) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	fmt.Println("in update Doc")
	fmt.Println(doc)
	fmt.Println(doc.ID)
	_, err = db.Exec(sqlUpdateDocToCache, doc.Name, doc.Author, doc.Type, doc.Description, doc.Fee, time.Now(), doc.ID)
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

	_, err = db.ExecContext(ctx, sqlUpdateStatusDocToCache, form.Status, form.ID, form.DocID)
	return
}

func (d *docCacheDAO) UpdateBorrowFormStatus(ctx context.Context, db *mssqlx.DBs, id uint64, status int) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}
	_, err = db.Exec(sqlUpdateBorrowFormStatus, status, time.Now(), id)
	if err != nil {
		return err
	}
	var id_doc uint64
	err = db.Get(&id_doc, sqlSelecetIdDoc, id)
	if err != nil {
		return err
	}
	_, err = db.Exec(sqlUpdateDocStatus, status, id_doc)
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlUpdateStatusDocToCache, status, id, id_doc)
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
