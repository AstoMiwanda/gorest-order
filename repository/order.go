package repository

import (
	"github.com/pushm0v/gorest-order/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(m *model.Order) (*model.Order, error)
	Update(m *model.Order, updateValue interface{}) (*model.Order, error)
	Delete(m *model.Order) error
	FindOne(id int) (*model.Order, error)
}

type orderRepository struct {
	dbConnection *gorm.DB
}

func NewOrderRepository(dbConnection *gorm.DB) OrderRepository {
	return &orderRepository{dbConnection: dbConnection}
}

func (c *orderRepository) Create(m *model.Order) (*model.Order, error) {
	var err = c.dbConnection.Create(m).Error
	return m, err
}

func (c *orderRepository) FindOne(id int) (m *model.Order, err error) {
	m = &model.Order{}
	err = c.dbConnection.First(m, id).Error

	return
}

func (c *orderRepository) Update(m *model.Order, updateValue interface{}) (*model.Order, error) {
	return m, c.dbConnection.Model(m).Updates(updateValue).Error
}

func (c *orderRepository) Delete(m *model.Order) error {
	return c.dbConnection.Delete(m).Error
}
