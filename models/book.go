package models

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type BookReq struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Amount   string `gorm:"size:255;not null" json:"amount" binding:"required"`
	Special  string `gorm:"size:255;not null" json:"special" binding:"required"`
	Date     string `gorm:"size:255;not null" json:"date" binding:"required"`
	Time     string `gorm:"size:255;not null" json:"time" binding:"required"`
	Person   string `gorm:"size:255;not null" json:"person" binding:"required"`
	Number   string `gorm:"size:255;not null" json:"number" binding:"required"`
	Accepted string `gorm:"size:255;not null; default:Күтілуде" json:"accepted"`
	Place    string `gorm:"size:255;not null" json:"place" binding:"required"`
}

func UpdateBook(id uint, updatedBook *BookReq) (*BookReq, error) {
	var book BookReq
	err := DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}

	err = DB.Model(&book).Updates(updatedBook).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func UpdateBookReq(c *gin.Context) {
	id := c.Param("id")

	var updatedBook BookReq
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	book, err := UpdateBook(uint(uintID), &updatedBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (b *BookReq) SaveBook() (*BookReq, error) {
	var err error
	err = DB.Create(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Book(c *gin.Context) {
	var book BookReq

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	cleanedNumber := strings.ReplaceAll(book.Number, " ", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, "(", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, ")", "")
	book.Number = strings.ReplaceAll(cleanedNumber, "-", "")
	_, err = book.SaveBook()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetAllBook() ([]BookReq, error) {
	var bookReqs []BookReq
	err := DB.Find(&bookReqs).Error
	if err != nil {
		return nil, err
	}
	return bookReqs, nil
}
func GetAllBookReqs(c *gin.Context) {
	bookReqs, err := GetAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookReqs)
}

func GetBookByNumber(number string) ([]BookReq, error) {
	var bookReqs []BookReq
	err := DB.Where("number = ?", number).Find(&bookReqs).Error
	if err != nil {
		return nil, err
	}
	return bookReqs, nil
}

func GetBookReqsByNumber(c *gin.Context) {
	number := c.Param("number")
	cleanedNumber := strings.ReplaceAll(number, " ", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, "(", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, ")", "")
	number = strings.ReplaceAll(cleanedNumber, "-", "")

	bookReqs, err := GetBookByNumber(number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookReqs)
}
