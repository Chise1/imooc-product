package services

import (
	"Chise1/imooc-product/datamodels"
	"Chise1/imooc-product/repositories"
)

type IOrderService interface {
	GetOrderByID(int64) (*datamodels.Order, error)
	DeleteOrderByID(int64) error
	UpdateOrder(*datamodels.Order) error
	InsertOrder(*datamodels.Order) (int64, error)
	GetAllOrder() ([]*datamodels.Order, error)
	GetAllOrderInfo() (map[int64]map[string]string, error)
}
type OrderService struct {
	OrderRepository repositories.IOrder
}

func NewOrderService(repository repositories.IOrder) IOrderService {
	return &OrderService{repository}
}

func (o *OrderService) GetOrderByID(orderID int64) (*datamodels.Order, error) {
	return o.GetOrderByID(orderID)
}
func (o *OrderService) DeleteOrderByID(orderID int64) error {
	return o.OrderRepository.Delete(orderID)
}
func (o *OrderService) UpdateOrder(order *datamodels.Order) error {
	return o.UpdateOrder(order)
}
func (o *OrderService) InsertOrder(order *datamodels.Order) (int64, error) {
	return o.InsertOrder(order)
}
func (o *OrderService) GetAllOrder() ([]*datamodels.Order, error) {
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) GetAllOrderInfo() (map[int64]map[string]string, error) {
	return o.OrderRepository.SelectAllWithInfo()
}
