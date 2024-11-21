package req

type AddItemReq struct {
	ProductId string `json:"product_id" form:"productId"`
	Quantity  int64  `json:"quantity" form:"quantity"`
}
