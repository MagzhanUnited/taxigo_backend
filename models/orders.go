package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name" binding:"required"`
	Amount      string `gorm:"size:255;not null" json:"amount" binding:"required"`
	Price       string `gorm:"size:255;not null" json:"price" binding:"required"`
	OrderListID uint   `gorm:"not null" json:"orderListID"`
}

type OrderList struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Place  string  `gorm:"size:255;not null" json:"place" binding:"required"`
	Number string  `gorm:"size:255;not null" json:"number" binding:"required"`
	Orders []Order `gorm:"foreignKey:OrderListID" json:"orders"`
}

func GetAllOrderReqs(c *gin.Context) {
	orderReqs, err := GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderReqs)
}
func GetAllOrders() ([]OrderList, error) {
	var order []OrderList
	err := DB.Preload("Orders").Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, err
}

func PostOrderReq(c *gin.Context) {
	var order OrderList

	if err := c.ShouldBindJSON(&order); err != nil {
		fmt.Println("SHOULD BIND", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	_, err = order.SaveOrder()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (b *OrderList) SaveOrder() (*OrderList, error) {
	var err error
	err = DB.Create(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetOrderByNumber(number string) ([]OrderList, error) {
	var orderList []OrderList
	fmt.Println("number:", number)
	err := DB.Preload("Orders").Model(&OrderList{}).Where("number = ?", number).Find(&orderList).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(orderList)
	return orderList, nil
}
func GetOrderReqsByNumber(c *gin.Context) {
	number := c.Param("number")
	// cleanedNumber := strings.ReplaceAll(number, " ", "")
	// cleanedNumber = strings.ReplaceAll(cleanedNumber, "(", "")
	// cleanedNumber = strings.ReplaceAll(cleanedNumber, ")", "")
	// number = strings.ReplaceAll(cleanedNumber, "-", "")
	orderReq, err := GetOrderByNumber(number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderReq)
}

// func GetBookByNumber(number string) ([]BookReq, error) {
// 	var bookReqs []BookReq
// 	err := DB.Where("number = ?", number).Find(&bookReqs).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bookReqs, nil
// }

// func GetBookReqsByNumber(c *gin.Context) {
// 	number := c.Param("number")
// 	cleanedNumber := strings.ReplaceAll(number, " ", "")
// 	cleanedNumber = strings.ReplaceAll(cleanedNumber, "(", "")
// 	cleanedNumber = strings.ReplaceAll(cleanedNumber, ")", "")
// 	number = strings.ReplaceAll(cleanedNumber, "-", "")

// 	bookReqs, err := GetBookByNumber(number)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, bookReqs)
// }

// func PostBook(c *gin.Context) {
// 	var bookFood BookFood

// 	if err := c.ShouldBindJSON(&bookFood); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var err error
// 	_, err = bookFood.SaveBook()
// 	if err != nil {
// 		fmt.Println("ERROR: ", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "success"})
// }
// func (b *BookFood) SaveBook() (*BookFood, error) {
// 	var err error
// 	err = DB.Create(&b).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return b, nil
// }
