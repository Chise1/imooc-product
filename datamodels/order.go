package datamodels

type Order struct {
	Id          int64 `gorm:"primaryKey"`
	UserId      int64
	ProductId   int64
	OrderStatus int64
	Product Product `gorm:"foreignKey:ProductId;references:Id;OnDelete:SET NULL;"`
}

const (
	OrderWait    = iota //0
	OrderSuccess        //1
	OrderFailed         //2
)
