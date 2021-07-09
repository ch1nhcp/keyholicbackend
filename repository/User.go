package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
	"finalbackend/util"
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
		return &user, nil
		err := util.CheckPasswordHash(user.Password, User.Password)
		if err != true {
			return &user, nil
		} else {
			return nil, errors.New("error")
		}
	}
	return nil, errors.New("error")
}
