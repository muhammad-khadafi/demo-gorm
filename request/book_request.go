package request

import "encoding/json"

type BookRequest struct {
	ID          json.Number `json:"id" binding:"required,number"`
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description"`
}
