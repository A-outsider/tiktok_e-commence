package req

type PlaceOrderReq struct {
	UserCurrency string       `json:"user_currency" form:"user_currency"`
	Address      *Address     `json:"address"  form:"address"`
	OrderItems   []*OrderItem `json:"order_items" form:"order_items"`
}

type Address struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address"  form:"address"`
}

type OrderItem struct {
	Item *CartItem `json:"item" form:"item"`
	Cost float64   `json:"cost" form:"cost"`
}

type CartItem struct {
	ProductId string `json:"product_id" form:"product_id"`
	Quantity  int64  `json:"quantity" form:"quantity"`
}

type ChangeStatusReq struct {
	OrderId string `json:"order_id" form:"order_id"`
}
