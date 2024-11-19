package req

type AddProductReq struct {
	Bid         string   `json:"bid" form:"bid"`
	Uid         string   `json:"uid" form:"uid"`
	Description string   `json:"description" form:"description"`
	Categories  []string `json:"categories" form:"categories"`
	Price       float64  `json:"price" form:"price"`
	Picture     string   `json:"picture" form:"picture"	`
	Name        string   `json:"name" form:"name"`
}

type SearchProductByCategoryReq struct {
	CategoryName string `query:"category" json:"categoryName"`
}

type SearchProductByQueryReq struct {
	Query string `query:"query" json:"query"`
}

type DeleteProductReq struct {
	Pid string `json:"pid" form:"pid"`
}

type GetProductReq struct {
	Id string `json:"id" query:"id"`
}
