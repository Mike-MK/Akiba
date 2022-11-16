package models

import (
	"fmt"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
}

func (u *User) SaveUser()(*User,error){
	
	err := DB.Create(u).Error
	if err != nil{
		fmt.Println("------------error------",err)
		return &User{},err
	}
	return u,nil
}

func (u *User)BeforeSave() error{
	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	u.Password = string(hashed)

	// trim username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}