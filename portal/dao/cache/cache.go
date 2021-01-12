package cache

import (
	"context"
	"database/sql"
	"errors"

	"github.com/linxGnu/mssqlx"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/model"
)

const (
	sqlSelectAllDocument  = "SELECT * FROM documents"
	sqlSelectDocumentByID = "SELECT * FROM documents WHERE doc_id = ?"
	sqlFirstDocumentID    = "SELECT doc_id FROM documents WHERE doc_name = ? AND category_id = ? LIMIT 1"
	sqlInsertDocument     = "INSERT INTO documents(doc_id, doc_name, category_id) VALUES (?,?,?)"
	sqlUpdateDocumentByID = "UPDATE documents SET doc_name = ?, category_id = ? WHERE doc_id = ?"

	sqlSelectAllCategories  = "SELECT * FROM categories"
	sqlSelectCategoryID     = "SELECT category_id FROM categories WHERE category_name = ?"
	sqlSelectCategoriesByID = "SELECT * FROM categories WHERE category_id = ?"
	sqlInsertCategories     = "INSERT INTO categories(category_id, category_name, doc_description) VALUE (?,?,?)"

	sqlSelectAllDocumentVersion     = "SELECT * FROM document_version"
	sqlSelectDocumentVersionByID    = "SELECT * FROM document_version WHERE document_version_id = ?"
	sqlSelectDocumentVersionByDocID = "SELECT * FROM document_version WHERE doc_id = ?"

	sqlFirstDocumentVersion  = "SELECT document_version_id FROM document_version WHERE doc_id = ? AND document_version = ? AND doc_description = ? AND author_id = ? AND fee = ? AND price = ? LIMIT 1"
	sqlInsertDocumentVersion = "INSERT INTO document_version(document_version_id, document_version, doc_id, doc_description, publisher, author_id, fee, price) VALUES (?,?,?,?,?,?,?,?)"
	sqlUpdateDocumentVersion = "UPDATE document_version SET document_version = ?, publisher = ?, author_id = ?, price = ? WHERE document_version_id = ?"

	sqlSelectAllAuthor  = "SELECT * FROM authors"
	sqlSelectAuthorByID = "SELECT * FROM authors WHERE author_id = ?"
	sqlSelectAuthorID   = "SELECT author_id FROM authors WHERE author_name = ?"
	sqlInsertAuthor     = "INSERT INTO authors(author_id, author_name, description) VALUES (?,?,?)"

	sqlSelectDocverIDFromCache               = "SELECT document_version_id FROM barcode_cache WHERE barcode_id = ?"
	sqlSelectCountBarcodeFromCacheByDocVerID = "SELECT COUNT(barcode_id) FROM barcode_cache WHERE document_version_id = ?"
	sqlInsertDocverIDToCache                 = "INSERT INTO barcode_cache(barcode_id, document_version_id) VALUES (?,?)"

	sqlInsertBlackList           = "INSERT INTO black_list(user_id, borrow_form_id, money) VALUES (?,?,?)"
	sqlSelectFromBlackListWithID = "SELECT * FROM black_list WHERE user_id = ?"
)

type IDocDAO interface {
	SelectAllDocument(ctx context.Context, db *mssqlx.DBs) (result []*model.DocumentsDAOobj, err error)
	SelectDocumentByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.DocumentsDAOobj, err error)
	SaveDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (err error)
	FirstOrCreateDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (docID uint64, err error)
	UpdateDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (err error)

	SelectAllCategories(ctx context.Context, db *mssqlx.DBs) (result []*model.CategoriesDAOobj, err error)
	SelectCategoriesByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.CategoriesDAOobj, err error)
	SelectCategoryID(ctx context.Context, db *mssqlx.DBs, category string) (catID uint64, err error)
	SaveCategories(ctx context.Context, db *mssqlx.DBs, cat *model.CategoriesDAOobj) (err error)
	FirstOrCreateCategory(ctx context.Context, db *mssqlx.DBs, category string, catIDIfNotExist uint64) (catID uint64, err error)

	SelectAllDocumentVersion(ctx context.Context, db *mssqlx.DBs) (result []*model.DocumentVersionDAOobj, err error)
	SelectDocumentVersionByID(ctx context.Context, db *mssqlx.DBs, id string) (result *model.DocumentVersionDAOobj, err error)
	SaveDocumentVersion(ctx context.Context, db *mssqlx.DBs, docver *model.DocumentVersionDAOobj) (err error)
	FirstOrCreateDocumentVersion(ctx context.Context, db *mssqlx.DBs, docver *model.DocumentVersionDAOobj) (docVerID string, err error)

	SelectAllAuthor(ctx context.Context, db *mssqlx.DBs) (result []*model.AuthorDAOobj, err error)
	SelectAuthorByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.AuthorDAOobj, err error)
	SelectAuthorID(ctx context.Context, db *mssqlx.DBs, authorName string) (authorID uint64, err error)
	SaveAuthor(ctx context.Context, db *mssqlx.DBs, aut *model.AuthorDAOobj) (err error)
	FirstOrCreateAuthor(ctx context.Context, db *mssqlx.DBs, author string, authIDIfNotExist uint64) (authorID uint64, err error)

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
	docID = 0
	if db == nil {
		return 0, core.ErrDBObjNull
	}

	if err = db.GetContext(ctx, &docID, sqlFirstDocumentID, doc.DocName, doc.CategoryID); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	if docID == 0 {
		if _, err = db.ExecContext(ctx, sqlInsertDocument, doc.DocID, doc.DocName, doc.CategoryID); err != nil {
			return 0, err
		}
		return doc.DocID, nil
	}

	return docID, nil
}

func UpdateDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (err error) {
	if db == nil {
		return core.ErrDBObjNull
	}

	if _, err := db.ExecContext(ctx, sqlUpdateDocumentByID, doc.DocName, doc.CategoryID, doc.DocID); err != nil {
		return err
	}

	return nil
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

func FirstOrCreateCategory(ctx context.Context, db *mssqlx.DBs, category string, catIDIfNotExist uint64) (catID uint64, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	catID, err = SelectCategoryID(ctx, db, category)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	if catID == 0 {
		newCat := &model.CategoriesDAOobj{
			CategoryID:   catIDIfNotExist,
			CategoryName: category,
		}

		if err = SaveCategories(ctx, db, newCat); err != nil {
			return 0, err
		}

		return catIDIfNotExist, nil
	}

	return catID, nil
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

func SelectDocumentVersionByDocumentID(ctx context.Context, db *mssqlx.DBs, docid uint64) (result []*model.DocumentVersionDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	err = db.SelectContext(ctx, &result, sqlSelectDocumentVersionByDocID, docid)
	return
}

func SelectDocumentVersionByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.DocumentVersionDAOobj, err error) {
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
	_, err = db.ExecContext(ctx, sqlInsertDocumentVersion, docver.DocVerID, docver.DocumentVersion, docver.DocID, docver.DocDescription, docver.Publisher, docver.AuthorID, docver.Fee, docver.Price)
	return
}

func UpdateDocumentVersion(ctx context.Context, db *mssqlx.DBs, docVerID uint64, documentVersion string, publisher string, authorID uint64, price uint64) (err error) {
	// UPDATE documents_version SET document_version = ?, publisher = ?, author_id = ?, price = ? WHERE document_version_id = ?
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	if _, err = db.ExecContext(ctx, sqlUpdateDocumentVersion, documentVersion, publisher, authorID, price, docVerID); err != nil {
		return err
	}

	return nil
}

func FirstOrCreateDocumentVersion(ctx context.Context, db *mssqlx.DBs, docver *model.DocumentVersionDAOobj) (docVerID uint64, err error) {
	docVerID = 0
	if db == nil {
		return 0, core.ErrDBObjNull
	}

	// doc_id, version, doc_description, author_id, fee, price
	if err = db.GetContext(ctx, &docVerID, sqlFirstDocumentVersion, docver.DocID, docver.DocumentVersion, docver.DocDescription, docver.AuthorID, docver.Fee, docver.Price); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	if docVerID == 0 {
		// INSERT INTO documents_version(document_version_id, document_version, doc_id, doc_description, publisher, author_id, fee, price) VALUES (?,?,?,?,?,?,?,?)
		if _, err := db.Exec(sqlInsertDocumentVersion, docver.DocVerID, docver.DocumentVersion, docver.DocID, docver.DocDescription, docver.Publisher, docver.AuthorID, docver.Fee, docver.Price); err != nil {
			return 0, err
		}
		return docver.DocVerID, nil
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

func FirstOrCreateAuthor(ctx context.Context, db *mssqlx.DBs, author string, authIDIfNotExist uint64) (authorID uint64, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	authorID, err = SelectAuthorID(ctx, db, author)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}

	if authorID == 0 {
		newAuthor := &model.AuthorDAOobj{
			AuthorID:   authIDIfNotExist,
			AuthorName: author,
		}
		if err = SaveAuthor(ctx, db, newAuthor); err != nil {
			return 0, err
		}

		return authIDIfNotExist, nil
	}

	return authorID, nil
}

// --------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- CACHE -----------------------------------------------------------

func SelectDocVerIDFromCacheByBarcode(ctx context.Context, db *mssqlx.DBs, barcodeID uint64) (docverID uint64, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	//TODO: fix model
	err = db.GetContext(ctx, &docverID, sqlSelectDocverIDFromCache, barcodeID)
	return
}

func SelectCountBarcodeByDocverID(ctx context.Context, db *mssqlx.DBs, docverid uint64) (count uint64, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	//TODO: fix model
	err = db.GetContext(ctx, &count, sqlSelectCountBarcodeFromCacheByDocVerID, docverid)
	return
}

func SaveDocverIDToCache(ctx context.Context, db *mssqlx.DBs, barcode uint64, docverID uint64) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}

	//TODO: fix model
	_, err = db.ExecContext(ctx, sqlInsertDocverIDToCache, barcode, docverID)
	return
}

// --------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- BLACK LIST ------------------------------------------------------

func InsertIntoBlackList(ctx context.Context, db *mssqlx.DBs, userID, borrowFormID, money uint64) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	_, err = db.ExecContext(ctx, sqlInsertBlackList, userID, borrowFormID, money)
	return
}

func SelectFromBlackListByUserID(ctx context.Context, db *mssqlx.DBs, userID uint64) (result []*model.BlackListDAOobj, err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	err = db.SelectContext(ctx, &result, sqlSelectFromBlackListWithID, userID)

	return
}
