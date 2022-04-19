package models

import "github.com/jinzhu/gorm"

type Orders struct {
	gorm.Model
	CustomerName string  `json:"customer_name"`
	Items        []Items `gorm:"foreignKey:OrderID"`
}

type Items struct {
	gorm.Model
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}

type OrdersCreate struct {
	CustomerName string        `json:"customer_name"`
	Items        []ItemsCreate `json:"items"`
}

type ItemsCreate struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
