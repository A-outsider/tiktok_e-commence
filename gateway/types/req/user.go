package req

type ModifyUserInfoReq struct {
	Name      string `form:"name"`
	Signature string `form:"signature"`
	Gender    int    `form:"gender"`
	Birthday  string `form:"birthday"`
}

type DeleteUserReq struct {
	Phone    string `form:"phone" binding:"required,phone"`
	AuthCode string `form:"auth_code" binding:"required"`
}

// 地址

type AddAddressReq struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
	Phone   string `form:"phone" binding:"required"`
}

type ModifyAddressReq struct {
	Aid     string `form:"aid" binding:"required"`
	Address string `form:"address"`
	Name    string `form:"name" `
	Phone   string `form:"phone" `
}

type DeleteAddressReq struct {
	Aid string `form:"aid" binding:"required"`
}

type SetDefaultAddressReq struct {
	Aid string `form:"aid" binding:"required"`
}
