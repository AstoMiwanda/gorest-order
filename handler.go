package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest-order/model"
	"github.com/pushm0v/gorest-order/service"
)

type OrderHandler struct {
	orderService service.OrderService
	custService  service.CustomerService
	notifService service.NotifService
}

func NewOrderHandler(orderService service.OrderService, custService service.CustomerService, notifService service.NotifService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		custService:  custService,
		notifService: notifService,
	}
}

func (s *OrderHandler) responseBuilder(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Response{
		Message: message,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error : %v", err)
	}
}

func (s *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	order, err := s.orderService.GetOrder(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Get Customer error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, order)
}

func (s *OrderHandler) Post(w http.ResponseWriter, r *http.Request) {

	var order = &model.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	cust, err := s.custService.GetCustomer(order.CustomerId)
	if err != nil {
		errMsg := fmt.Sprintf("Error get customer : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	order, err = s.orderService.CreateOrder(order)
	if err != nil {
		errMsg := fmt.Sprintf("Create order error : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.notifService.SendEmailOrderCreated(cust, order)
	if err != nil {
		log.Printf("Error sending email : %v\n", err)
	}

	w.WriteHeader(http.StatusCreated)
	s.responseBuilder(w, "order created")
}

func (s *OrderHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	var order = &model.Order{}
	err = json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	_, err = s.custService.GetCustomer(order.CustomerId)
	if err != nil {
		errMsg := fmt.Sprintf("Error get customer : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	_, err = s.orderService.UpdateOrder(orderID, order)
	if err != nil {
		errMsg := fmt.Sprintf("Update customer error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	s.responseBuilder(w, "order updated")
}

func (s *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.orderService.DeleteOrder(orderID)
	if err != nil {
		errMsg := fmt.Sprintf("Delete order error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "customer deleted")
}

func (s *OrderHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.responseBuilder(w, "not found")
}
