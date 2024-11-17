package req

type AddProductReq struct {
	Bid         string   `json:"bid" form:"bid"`
	Uid         string   `json:"uid" form:"uid"`
	Description string   `json:"description" form:"description"`
	Categories  []string `json:"categories" form:"categories"`
	Price       float64  `json:"price" form:"price"`
	Picture     string   `json:"picture" form:"picture"	`
}

type SearchProductByCategoryReq struct {
	Category string `json:"category"`
}
