package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"avicena/models"
	"avicena/utils/token"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Password string `json:"password" binding:"required"`
	Number   string `json:"number" binding:"required"`
}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})

}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Password = input.Password
	u.Number = input.Number
	var username string
	var status string
	token, err, username, status := models.LoginCheck(u.Number, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password or number incorrec"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "token": token, "username": username, "status": status})
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Number   string `json:"number" binding:"required"`
	Status   string `json:"status"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password
	cleanedNumber := strings.ReplaceAll(input.Number, " ", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, "(", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, ")", "")
	input.Number = strings.ReplaceAll(cleanedNumber, "-", "")
	u.Number = input.Number
	u.Status = input.Status
	print(u.Password)
	normPassword := u.Password
	var err error
	fmt.Println("u.Number", u.Number)
	_, err = u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password or number incorrec"})
		return
	}
	var username string
	fmt.Println("u.Number", u.Number)
	token, err, username, status := models.LoginCheck(u.Number, normPassword)
	fmt.Println(token)
	fmt.Println(username)
	fmt.Println(status)
	c.JSON(http.StatusOK, gin.H{"message": "success", "token": token, "username": username, "status": status})
}
