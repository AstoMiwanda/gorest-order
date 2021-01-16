package service

import (
	"github.com/pushm0v/gorest-order/client"
	"github.com/pushm0v/gorest-order/model"
)

type CustomerService interface {
	GetCustomer(id int) (cust *model.Customer, err error)
}

type customerService struct {
	gorestClient client.GorestClient
}

func NewCustomerService(gorestClient client.GorestClient) CustomerService {
	return &customerService{
		gorestClient: gorestClient,
	}
}

func (c *customerService) GetCustomer(id int) (cust *model.Customer, err error) {
	cust, err = c.gorestClient.GetCustomerByID(id)

	return
}
