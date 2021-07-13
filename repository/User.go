package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
	"finalbackend/util"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(User *models.User) error {
	var user models.User
	database.DB.Raw("SELECT * FROM `users` WHERE email = ? ", User.Email).Scan(&user)
	if user.Email == User.Email {
		return errors.New("email exist")
	} else {
		hashpassword, _ := util.HashPassword(User.Password)
		User.Password = hashpassword
		database.DB.Create(&User)
		return nil
	}

}

func Login(User *models.User) (*models.User, error) {
	var user models.User
	database.DB.Raw("SELECT * FROM `users` WHERE `email` = ?", User.Email).Scan(&user)
	if user.Id > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(User.Password))
		if err != nil {
			return nil, errors.New("error")
		} else {
			return &user, nil
		}
	}
	return nil, errors.New("error")
}
