package entity

import (
	"time"
)

type Book struct {
	ID          uint
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
