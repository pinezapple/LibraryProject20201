package model

type SelectBarcodeByIDReq struct {
	BarcodeID uint64 `json:"barcode_id"`
}
type SelectPaymentByIDReq struct {
	PaymentID uint64 `json:"payment_id"`
}

type SelectSaleBillByIDReq struct {
	SaleBillID uint64 `json:"sale_bill_id"`
}
