package cache

import (
	"context"

	"github.com/linxGnu/mssqlx"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/model"
)

const (
	sqlSelectAllDocument  = "SELECT * FROM documents"
	sqlSelectDocumentByID = "SELECT * FROM documents WHERE doc_id = ?"
	sqlFirstDocumentID    = "SELECT doc_id FROM documents WHERE doc_name = ? AND category_id = ? LIMIT 1"
	sqlInsertDocument     = "INSERT INTO documents(doc_id, doc_name, category_id) VALUES (?,?)"

	sqlSelectAllCategories  = "SELECT * FROM categories"
	sqlSelectCategoryID     = "SELECT category_id FROM categories WHERE category_name = ?"
	sqlSelectCategoriesByID = "SELECT * FROM categories WHERE category_id = ?"
	sqlInsertCategories     = "INSERT INTO categories(category_id, category_name, doc_description) VALUE (?,?,?)"

	sqlSelectAllDocumentVersion  = "SELECT * FROM documents_version"
	sqlSelectDocumentVersionByID = "SELECT * FROM documents_version WHERE documents_version = ?"
	sqlFirstDocumentVersion      = "SELECT document_version FROM documents_version WHERE doc_id = ? AND version = ? AND doc_description = ? AND author_id = ? AND fee = ? AND price = ? LIMIT 1"
	sqlInsertDocumentVersion     = "INSERT INTO documents_version(document_version, doc_id, version, doc_description, author_id, fee, price) VALUES (?,?,?,?,?,?,?)"

	sqlSelectAllAuthor  = "SELECT * FROM authors"
	sqlSelectAuthorByID = "SELECT * FROM authors WHERE author_id = ?"
	sqlSelectAuthorID   = "SELECT author_id FROM authors WHERE author_name = ?"
	sqlInsertAuthor     = "INSERT INTO authors(author_id, author_name, description) VALUES (?,?,?)"

	sqlSelectDocverFromCache = "SELECT document_version FROM barcode_cache WHERE barcode_id = ?"
	sqlInsertDocverToCache   = "INSERT INTO barcode_cache(barcore_id, document_version) VALUES (?,?)"
)

type IDocDAO interface {
	SelectAllDocument(ctx context.Context, db *mssqlx.DBs) (result []*model.DocumentsDAOobj, err error)
	SelectDocumentByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.DocumentsDAOobj, err error)
	SaveDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (err error)
	FirstOrCreateDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (docID uint64, err error)

	SelectAllCategories(ctx context.Context, db *mssqlx.DBs) (result []*model.CategoriesDAOobj, err error)
	SelectCategoriesByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.CategoriesDAOobj, err error)
	SelectCategoryID(ctx context.Context, db *mssqlx.DBs, category string) (catID uint64, err error)
	SaveCategories(ctx context.Context, db *mssqlx.DBs, cat *model.CategoriesDAOobj) (err error)

	SelectAllDocumentVersion(ctx context.Context, db *mssqlx.DBs) (result []*model.DocumentVersionDAOobj, err error)
	SelectDocumentVersionByID(ctx context.Context, db *mssqlx.DBs, id string) (result *model.DocumentVersionDAOobj, err error)
	SaveDocumentVersion(ctx context.Context, db *mssqlx.DBs, docver *model.DocumentVersionDAOobj) (err error)
	FirstOrCreateDocumentVersion(ctx context.Context, db *mssqlx.DBs, docver *model.DocumentVersionDAOobj) (docVerID string, err error)

	SelectAllAuthor(ctx context.Context, db *mssqlx.DBs) (result []*model.AuthorDAOobj, err error)
	SelectAuthorByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.AuthorDAOobj, err error)
	SelectAuthorID(ctx context.Context, db *mssqlx.DBs, authorName string) (authorID uint64, err error)
	SaveAuthor(ctx context.Context, db *mssqlx.DBs, aut *model.AuthorDAOobj) (err error)

	SelectDocVerFromCacheByBarcode(ctx context.Context, db *mssqlx.DBs, barcodeID string) (docver string, err error)
}

type DocCacheDAO struct{}

// -----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- DOCUMENTS ----------------------------------------------------------

func SelectAllDocument(ctx context.Context, db *mssqlx.DBs) (result []*model.DocumentsDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllDocument)
	return
}

func SelectDocumentByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.DocumentsDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	result = &model.DocumentsDAOobj{}
	err = db.GetContext(ctx, result, sqlSelectDocumentByID, id)
	return
}

func SaveDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	_, err = db.ExecContext(ctx, sqlInsertDocument, doc.DocID, doc.DocName, doc.CategoryID)
	return
}

func FirstOrCreateDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (docID uint64, err error) {
	if db == nil {
		return 0, core.ErrDBObjNull
	}

	if err = db.GetContext(ctx, &docID, sqlFirstDocumentID, doc.DocName, doc.CategoryID); err != nil {
		return 0, err
	}

	if docID == 0 {
		if _, err = db.ExecContext(ctx, sqlInsertDocument, doc.DocID, doc.DocName, doc.CategoryID); err != nil {
			return 0, err
		}
	}

	return docID, nil
}

// ------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- CATEGORIES ----------------------------------------------------------

func SelectAllCategories(ctx context.Context, db *mssqlx.DBs) (result []*model.CategoriesDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllCategories)
	return
}

func SelectCategoriesByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.CategoriesDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	result = &model.CategoriesDAOobj{}
	err = db.GetContext(ctx, result, sqlSelectCategoriesByID, id)
	return
}

func SelectCategoryID(ctx context.Context, db *mssqlx.DBs, category string) (catID uint64, err error) {
	if db == nil {
		return 0, core.ErrDBObjNull
	}

	if err = db.GetContext(ctx, &catID, sqlSelectCategoryID, category); err != nil {
		return 0, err
	}

	return catID, nil
}

func SaveCategories(ctx context.Context, db *mssqlx.DBs, cat *model.CategoriesDAOobj) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	_, err = db.ExecContext(ctx, sqlInsertCategories, cat.CategoryID, cat.CategoryName, cat.DocDescription)
	return
}

// ------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- DOCUMENT VERSION ----------------------------------------------------------

func SelectAllDocumentVersion(ctx context.Context, db *mssqlx.DBs) (result []*model.DocumentVersionDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllDocumentVersion)
	return
}

func SelectDocumentVersionByID(ctx context.Context, db *mssqlx.DBs, id string) (result *model.DocumentVersionDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	result = &model.DocumentVersionDAOobj{}
	err = db.GetContext(ctx, result, sqlSelectDocumentVersionByID, id)
	return
}

func SaveDocumentVersion(ctx context.Context, db *mssqlx.DBs, docver *model.DocumentVersionDAOobj) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	_, err = db.ExecContext(ctx, sqlInsertDocumentVersion, docver.DocumentVersion, docver.DocID, docver.Version, docver.DocDescription, docver.AuthorID, docver.Fee, docver.Price)
	return
}

func FirstOrCreateDocumentVersion(ctx context.Context, db *mssqlx.DBs, docver *model.DocumentVersionDAOobj) (docVerID string, err error) {
	if db == nil {
		return "", core.ErrDBObjNull
	}
	// doc_id, version, doc_description, author_id, fee, price
	if err = db.GetContext(ctx, &docVerID, sqlFirstDocumentVersion, docver.DocID, docver.Version, docver.DocDescription, docver.AuthorID, docver.Fee, docver.Price); err != nil {
		return "", err
	}

	if docVerID == "" {
		if _, err = db.Exec(sqlInsertDocumentVersion, docver.DocumentVersion, docver.DocID, docver.Version, docver.DocDescription, docver.AuthorID, docver.Fee, docver.Price); err != nil {
			return "", err
		}
	}

	return docVerID, nil
}

// --------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- AUTHOR ----------------------------------------------------------

func SelectAllAuthor(ctx context.Context, db *mssqlx.DBs) (result []*model.AuthorDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, &result, sqlSelectAllAuthor)
	return
}

func SelectAuthorByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.AuthorDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	result = &model.AuthorDAOobj{}
	err = db.GetContext(ctx, result, sqlSelectAuthorByID, id)
	return
}

func SelectAuthorID(ctx context.Context, db *mssqlx.DBs, authorName string) (authorID uint64, err error) {
	if db == nil {
		return 0, core.ErrDBObjNull
	}

	if err = db.GetContext(ctx, &authorID, sqlSelectAuthorID, authorName); err != nil {
		return 0, err
	}

	return authorID, nil
}

func SaveAuthor(ctx context.Context, db *mssqlx.DBs, aut *model.AuthorDAOobj) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	_, err = db.ExecContext(ctx, sqlInsertAuthor, aut.AuthorID, aut.AuthorName, aut.Description)
	return
}

func SelectDocVerFromCacheByBarcode(ctx context.Context, db *mssqlx.DBs, barcodeID uint64) (docver string, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.GetContext(ctx, docver, sqlSelectDocverFromCache, barcodeID)
	return
}

func SaveDocverToCache(ctx context.Context, db *mssqlx.DBs, barcode, docver uint64) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	_, err = db.ExecContext(ctx, sqlInsertDocverToCache, barcode, docver)
	return
}
