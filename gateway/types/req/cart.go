package req

type AddItemReq struct {
	ProductId string `json:"product_id" form:"product_id"`
	Quantity  int64  `json:"quantity" form:"quantity"`
}
