package repositories

import (
	"Chise1/imooc-product/datamodels"
	"Chise1/imooc-product/db"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type IOrder interface {
	Insert(*datamodels.Order) (int64, error)
	Delete(int64) error
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int64]map[string]string, error)
}
type OrderManager struct {
	mysqlConn *gorm.DB
}

func NewOrderManager() IOrder {
	res := &OrderManager{}
	var err error
	res.mysqlConn, err = db.GetDB()
	if err != nil {
		panic("con error!")
	} else {
		err := res.mysqlConn.AutoMigrate(&datamodels.Order{})
		if err != nil {
			fmt.Println(err)
		}
	}
	return res
}
func (o *OrderManager) Insert(order *datamodels.Order) (int64, error) {
	tx := o.mysqlConn.Create(&order)
	return order.Id, tx.Error
}
func (o *OrderManager) Delete(orderId int64) error {
	tx := o.mysqlConn.Delete(&datamodels.Order{}, orderId)
	return tx.Error
}
func (o *OrderManager) Update(order *datamodels.Order) error {
	tx := o.mysqlConn.Model(&order).Select("OrderStatus").Updates(&order)
	return tx.Error
}
func (o *OrderManager) SelectByKey(orderId int64) (*datamodels.Order, error) {
	var order datamodels.Order
	tx := o.mysqlConn.First(&order, orderId)
	return &order, tx.Error
}

func (o *OrderManager) SelectAll() ([]*datamodels.Order, error) {
	var orders []*datamodels.Order
	tx := o.mysqlConn.Find(&orders)
	return orders, tx.Error
}
func (o *OrderManager) SelectAllWithInfo() (map[int64]map[string]string, error) {
	var orders []*datamodels.Order
	tx := o.mysqlConn.Model(&datamodels.Order{}).Joins("Product").Find(&orders)
	result := make(map[int64]map[string]string)
	for _, order := range orders {
		r := make(map[string]string)
		r["OrderStatus"] = strconv.FormatInt(order.OrderStatus, 10)
		r["ProductName"] = order.Product.ProductName
		r["ID"] = strconv.FormatInt(order.Id, 10)
		result[order.Id] = r
	}
	return result, tx.Error
}
