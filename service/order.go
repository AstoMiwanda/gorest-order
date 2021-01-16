package service

import (
	"github.com/pushm0v/gorest-order/model"
	"github.com/pushm0v/gorest-order/repository"
)

type OrderService interface {
	GetOrder(id int) (m *model.Order, err error)
	CreateOrder(m *model.Order) (*model.Order, error)
	UpdateOrder(id int, cust *model.Order) (m *model.Order, err error)
	DeleteOrder(id int) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (c *orderService) getOrderById(id int) (m *model.Order, err error) {
	m, err = c.orderRepository.FindOne(id)
	return
}

func (c *orderService) GetOrder(id int) (m *model.Order, err error) {
	return c.getOrderById(id)
}

func (c *orderService) CreateOrder(m *model.Order) (*model.Order, error) {
	return c.orderRepository.Create(m)
}

func (c *orderService) UpdateOrder(id int, m *model.Order) (*model.Order, error) {
	existingOrder, err := c.getOrderById(id)
	if err != nil {
		return nil, err
	}

	var updateValue = model.Order{
		SKU:   m.SKU,
		Qty:   m.Qty,
		Price: m.Price,
	}

	return c.orderRepository.Update(existingOrder, updateValue)
}

func (c *orderService) DeleteOrder(id int) error {
	existingOrder, err := c.getOrderById(id)
	if err != nil {
		return err
	}

	return c.orderRepository.Delete(existingOrder)
}
