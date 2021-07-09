package models

type User struct {
	Id         int    `json:"Id"`
	Email      string `json:"Email"`
	Name       string `json:"Name"`
	Password   string `json:"Password"`
	Permission bool   `json:"Permission"`
}
