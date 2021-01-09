package model

import "github.com/pinezapple/LibraryProject20201/skeleton/model"

type DocumentsDAOobj struct {
	DocID      uint64      `json:"doc_id" db:"doc_id"`
	DocName    string      `json:"doc_name" db:"doc_name"`
	CategoryID uint64      `json:"category_id" db:"category_id"`
	CreatedAt  *model.Time `json:"created_at" db:"created_at"`
	UpdatedAt  *model.Time `json:"updated_at" db:"updated_at"`
}

type CategoriesDAOobj struct {
	CategoryID     uint64      `json:"category_id" db:"category_id"`
	CategoryName   string      `json:"category_name" db:"category_name"`
	DocDescription string      `json:"doc_description" db:"doc_description"`
	CreatedAt      *model.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *model.Time `json:"updated_at" db:"updated_at"`
}

type DocumentVersionDAOobj struct {
	DocumentVersion string      `json:"document_version" db:"document_version"`
	DocID           uint64      `json:"doc_id" db:"doc_id"`
	Version         string      `json:"version" db:"version"`
	DocDescription  string      `json:"doc_description" db:"doc_description"`
	Publisher       string      `json:"publisher" db:"publisher"`
	AuthorID        uint64      `json:"author_id" db:"author_id"`
	Fee             uint64      `json:"fee" db:"fee"`
	Price           uint64      `json:"price" db:"price"`
	CreatedAt       *model.Time `json:"created_at" db:"created_at"`
	UpdatedAt       *model.Time `json:"updated_at" db:"updated_at"`
}

type AuthorDAOobj struct {
	AuthorID    uint64      `json:"author_id" db:"author_id"`
	AuthorName  string      `json:"author_name" db:"author_name"`
	Description string      `json:"description" db:"description"`
	CreatedAt   *model.Time `json:"created_at" db:"created_at"`
	UpdatedAt   *model.Time `json:"updated_at" db:"updated_at"`
}

type BlackListDAOobj struct {
	UserID       uint64      `json:"user_id" db:"user_id"`
	BorrowFormID uint64      `json:"borrow_form_id" db:"borrow_form_id"`
	Money        uint64      `json:"money" db:"money"`
	CreatedAt    *model.Time `json:"created_at" db:"created_at"`
}
