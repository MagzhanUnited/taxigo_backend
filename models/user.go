package models

import (
	"avicena/utils/token"
	"errors"
	"fmt"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `gorm:"size:255;not null ;" json:"username"`
	Password string `gorm:"size:1000; not null;" json:"password"`
	Number   string `gorm:"size:255; not null;unique" json:"number"`
	ID       uint   `gorm:"primaryKey" json:"id"`
	Status   string `gorm:"size:255;default:Қолданушы" json:"status"`
}

// func registerCheck(){}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUserByID(uid uint) (User, error) {
	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found")
	}
	u.PrepareGive()
	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func LoginCheck(number string, password string) (string, error, string, string) {
	var err error

	u := User{}
	cleanedNumber := strings.ReplaceAll(number, " ", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, "(", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, ")", "")
	number = strings.ReplaceAll(cleanedNumber, "-", "")

	err = DB.Model(User{}).Where("number = ?", number).Take(&u).Error
	if err != nil {
		return "", err, "", ""
	}
	fmt.Println("norm password: ", password)
	err = VerifyPassword(password, u.Password)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return "", err, "", ""
	}
	var token_ string
	token_, err = token.GenerateToken(u.ID)
	if err != nil {
		return "", err, "", ""
	}
	return token_, nil, u.Username, u.Status
}
func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	fmt.Println("BeforeSave: ", u.Number)
	if err != nil {
		return err
	}
	fmt.Println("##############")
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Number = html.EscapeString(strings.TrimSpace(u.Number))

	return nil

}
func (u *User) SaveUser() (*User, error) {
	fmt.Println("Number that's gonna be saved:", u.Number)
	u.BeforeSave()

	var err error
	fmt.Println("AAAA", u.Number)
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
