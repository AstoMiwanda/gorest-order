package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest-order/client"
	"github.com/pushm0v/gorest-order/model"
	"github.com/pushm0v/gorest-order/repository"
	"github.com/pushm0v/gorest-order/service"
)

func RestRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	orderRouter(api)
	r.Use(LoggingMiddleware)
	return r
}

func orderRouter(r *mux.Router) {
	var dbConn, err = NewDBConnection("order.db")
	if err != nil {
		log.Fatalf("DB Connection error : %v", err)
	}
	dbConn.AutoMigrate(&model.Customer{})
	var orderRepository = repository.NewOrderRepository(dbConn)
	var orderService = service.NewOrderService(orderRepository)
	var gorestNotifClient = client.NewGorestNotif(os.Getenv("GOREST_NOTIF_ADDR"))
	var gorestClient = client.NewGorestClient(os.Getenv("GOREST_ADDR"))
	var notifService = service.NewNotifService(gorestNotifClient)
	var custService = service.NewCustomerService(gorestClient)
	var orderHandler = NewOrderHandler(orderService, custService, notifService)

	r.HandleFunc("/orders/{id}", orderHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/orders", orderHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/orders/{id}", orderHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/orders/{id}", orderHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", orderHandler.NotFound)
}
