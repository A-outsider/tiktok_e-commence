package req

type CreatePaymentReq struct {
	OId    string  `form:"oid"`
	Amount float64 `form:"amount"`
}
