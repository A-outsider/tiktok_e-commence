package model

type Payment struct {
	PayId string `json:"pay_id" gorm:"column:pay_id; index" `
}
