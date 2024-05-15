package dto

type UserDTO struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Email  string     `json:"email"`
	Status string     `json:"status"`
	Orders []OrderDTO `json:"orders,omitempty"`
}
