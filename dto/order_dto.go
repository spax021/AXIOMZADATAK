package dto

type OrderDTO struct {
	ID          int     `json:"id"`
	UserID      int     `json:"userId"`
	TotalAmount float64 `json:"totalAmount"`
	Status      string  `json:"status"`
}
