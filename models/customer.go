package models

type Customer struct {
	Name  string `json:"name"`
	Phone string `json:"phone" binding:"required"`
}
