package model

type Order struct {
	Id          int    `json:"id" gorm:"autoIncrement;primaryKey"`
	OrderNumber string `json:"order_number"`
	CustomerId  int    `json:"customer_id"`
	SKU         string `json:"sku"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
}
