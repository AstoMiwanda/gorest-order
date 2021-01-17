package service

import (
	"fmt"

	"github.com/pushm0v/gorest-order/client"
	"github.com/pushm0v/gorest-order/model"
)

type NotifService interface {
	SendEmailOrderCreated(cust *model.Customer, order *model.Order) (err error)
}

type notifService struct {
	notifClient client.GorestNotif
}

func NewNotifService(notifClient client.GorestNotif) NotifService {
	return &notifService{
		notifClient: notifClient,
	}
}

func (n *notifService) SendEmailOrderCreated(cust *model.Customer, order *model.Order) (err error) {
	var m = new(model.EmailMessage)
	m.Destination = cust.Email
	m.DestinationName = cust.Name
	m.Subject = fmt.Sprintf("We received your order #%s!", order.OrderNumber)
	m.Body = fmt.Sprintf("Hi %s !, we would like to inform you that your order has been processing.", cust.Name)

	return n.notifClient.SendEmail(m)
}
