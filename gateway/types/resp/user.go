package resp

type GetUserInfoResp struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	AvatarPath     string `json:"avatar_path"`
	Phone          string `json:"phone"`
	Role           int64  `json:"role"`
	Signature      string `json:"signature"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	DefaultAddress string `json:"default_address"` // id
}

type GetAddressListResp struct {
	Addresses []Address `json:"address_list"`
}

type Address struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
