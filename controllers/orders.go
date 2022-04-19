package controllers

import (
	"assigment-2/models"
	"assigment-2/views"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type controllersOrders struct {
	db *gorm.DB
}

func NewControllerOrders(db *gorm.DB) *controllersOrders {
	return &controllersOrders{
		db: db,
	}
}

func (in *controllersOrders) GetOrders(c *gin.Context) {
	var (
		order  []models.Orders
		result gin.H
	)

	in.db.Preload("Items").Find(&order)

	if len(order) == 0 {
		result = gin.H{
			"msg": "There is no data in database",
		}
	} else {
		result = gin.H{
			"data":  order,
			"count": len(order),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (in *controllersOrders) CreateOrders(c *gin.Context) {
	var orderCreate models.OrdersCreate
	var order models.Orders

	err := json.NewDecoder(c.Request.Body).Decode(&orderCreate)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order.CustomerName = orderCreate.CustomerName
	err = in.db.Create(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, o := range orderCreate.Items {
		newItem := models.Items{
			ItemCode:    o.ItemCode,
			Description: o.Description,
			Quantity:    o.Quantity,
			OrderID:     int(order.ID),
		}
		order.Items = append(order.Items, newItem)
	}

	for _, item := range order.Items {
		err = in.db.Create(&item).Error
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	var orderViews views.OrdersCreateView

	orderViews.CustomerName = order.CustomerName
	orderViews.OrderedAt = order.CreatedAt.String()
	for _, o := range order.Items {
		newItem := views.ItemsCreateViews{
			ItemCode:    o.ItemCode,
			Description: o.Description,
			Quantity:    o.Quantity,
		}
		orderViews.Items = append(orderViews.Items, newItem)
	}

	c.JSON(http.StatusOK, orderViews)
}

func (in *controllersOrders) UpdateOrdersByID(c *gin.Context) {
	var (
		order    models.Orders
		newOrder models.Orders
	)

	id := c.Param("id")

	err := in.db.First(&order, id).Error
	if err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newOrder)
	if err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Model(&order).Updates(newOrder).Error
	if err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update data successfull !",
	})
}

func (in *controllersOrders) DeleteOrderByID(c *gin.Context) {
	var (
		order models.Orders
	)

	id := c.Param("id")

	err := in.db.First(&order, id).Error
	if err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Delete(&order).Error
	if err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete data successfull !",
	})
}
