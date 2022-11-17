package models

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Number string `gorm:"unique"`
	UserID uint `gorm:"unique"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;unique"`
}

func (a *Account) SaveAccount() (*Account, error) {

	err := DB.Create(a).Error
	if err != nil {
		return &Account{}, err
	}
	return a, nil
}

func (a *Account) BeforeSave() error{
	var nonNumericRegex = regexp.MustCompile(`[^0-9 ]+`)
	formattedNumber := nonNumericRegex.ReplaceAllString(a.Number, "")
	if string(formattedNumber[0]) == "0"{
		formattedNumber = "254" + string(formattedNumber[1:])
	}

	a.Number = formattedNumber
	fmt.Println("formatted",a.Number)
	if len(a.Number) != 12 {
		return errors.New("number is invalid")
	}
	return nil
}