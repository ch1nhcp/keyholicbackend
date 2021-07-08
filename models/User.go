package models

type User struct {
	Id         int    `json:"Id"`
	Email      string `json:"Email"`
	Name       string `json:"Name"`
	Permission bool   `json:"Permission"`
}
