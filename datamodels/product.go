package datamodels

type Product struct {
	Id           int64  `json:"id" gorm:"primaryKey"`
	ProductName  string `json:"ProductName"`
	ProductNum   int64  `json:"ProductNum"`
	ProductImage string `json:"ProductImage"`
	ProductUrl   string `json:"ProductUrl"`
}
