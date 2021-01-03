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
	sqlInsertDocument     = "INSERT INTO documents(doc_id, doc_name, category_id) VALUES (?,?)"

	sqlSelectAllCategories  = "SELECT * FROM categories"
	sqlSelectCategoriesByID = "SELECT * FROM categories WHERE category_id = ?"
	sqlInsertCategories     = "INSERT INTO categories(category_id, category_name, doc_description) VALUE (?,?,?)"

	sqlSelectAllDocumentVersion  = "SELECT * FROM documents_version"
	sqlSelectDocumentVersionByID = "SELECT * FROM documents_version WHERE documents_version = ?"
	sqlInsertDocumentVersion     = "INSERT INTO documents_version(document_version, doc_id, version, doc_description, author_id, fee, price) VALUES (?,?,?,?,?,?,?)"

	sqlSelectAllAuthor  = "SELECT * FROM authors"
	sqlSelectAuthorByID = "SELECT * FROM authors WHERE author_id = ?"
	sqlInsertAuthor     = "INSERT INTO authors(author_id, author_name, description) VALUES (?,?,?)"
)

type IDocDAO interface {
	SelectAllDocument(ctx context.Context, db *mssqlx.DBs) (result []*model.DocumentsDAOobj, err error)
	SelectDocumentByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.DocumentsDAOobj, err error)
	SaveDocument(ctx context.Context, db *mssqlx.DBs, doc *model.DocumentsDAOobj) (err error)

	SelectAllCategories(ctx context.Context, db *mssqlx.DBs) (result []*model.CategoriesDAOobj, err error)
	SelectCategoriesByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.CategoriesDAOobj, err error)
	SaveCategories(ctx context.Context, db *mssqlx.DBs, cat *model.CategoriesDAOobj) (err error)

	SelectAllDocumentVersion(ctx context.Context, db *mssqlx.DBs) (result []*model.DocumentVersionDAOobj, err error)
	SelectDocumentVersionByID(ctx context.Context, db *mssqlx.DBs, id string) (result *model.DocumentVersionDAOobj, err error)
	SaveDocumentVersion(ctx context.Context, db *mssqlx.DBs, docver *model.DocumentVersionDAOobj) (err error)

	SelectAllAuthor(ctx context.Context, db *mssqlx.DBs) (result []*model.AuthorDAOobj, err error)
	SelectAuthorByID(ctx context.Context, db *mssqlx.DBs, id uint64) (result *model.AuthorDAOobj, err error)
	SaveAuthor(ctx context.Context, db *mssqlx.DBs, aut *model.AuthorDAOobj) (err error)
}

type DocCacheDAO struct{}

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

func SaveCategories(ctx context.Context, db *mssqlx.DBs, cat *model.CategoriesDAOobj) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	_, err = db.ExecContext(ctx, sqlInsertCategories, cat.CategoryID, cat.CategoryName, cat.DocDescription)
	return
}

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

func SaveAuthor(ctx context.Context, db *mssqlx.DBs, aut *model.AuthorDAOobj) (err error) {
	// Validate input
	if db == nil {
		err = core.ErrDBObjNull
		return
	}
	_, err = db.ExecContext(ctx, sqlInsertAuthor, aut.AuthorID, aut.AuthorName, aut.Description)
	return
}
