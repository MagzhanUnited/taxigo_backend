package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BookFood) SaveBook() (*BookFood, error) {
	var err error
	err = DB.Create(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func PostBook(c *gin.Context) {
	var bookFood BookFood

	if err := c.ShouldBindJSON(&bookFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	_, err = bookFood.SaveBook()
	if err != nil {
		fmt.Println("ERROR: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

type BookFood struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:255;not null" json:"name"`
	Price string `gorm:"size:255;not null" json:"price"`
	Place string `gorm:"size:255;not null" json:"place" binding:"required"`
	Image string `gorm:"size:1000; not null" json:"image"`
}

type inputFood struct {
	Place string `gorm:"size:255;not null" json:"place" binding:"required"`
}

func GetBook(c *gin.Context) {
	var inputFood BookFood
	if err := c.ShouldBindJSON(&inputFood); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var books []BookFood
	err := DB.Model(&BookFood{}).Where("place = ?", inputFood.Place).Find(&books).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
