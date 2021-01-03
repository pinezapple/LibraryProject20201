package model

import "github.com/pinezapple/LibraryProject20201/skeleton/model"

type SelectBarcodeByIDReq struct {
	BarcodeID uint64 `json:"barcode_id"`
}

type BarcodeFrontEndResp struct {
	BarcodeID   uint64      `json:"barcode_id"`
	Status      uint64      `json:"status"`
	DocName     string      `json:"doc_name"`
	Version     string      `json:"version"`
	Author      string      `json:"author"`
	Fee         uint64      `json:"fee"`
	Price       uint64      `json:"price"`
	Description string      `json:"description"`
	Category    string      `json:"category"`
	SaleBillID  uint64      `json:"sale_bill_id"`
	CreatedAt   *model.Time `json:"created_at"`
}
type SelectPaymentByIDReq struct {
	PaymentID uint64 `json:"payment_id"`
}

type SelectSaleBillByIDReq struct {
	SaleBillID uint64 `json:"sale_bill_id"`
}
