package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Category    string  `json:"category" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Imagepath   string  `json:"imagepath" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Size        string  `json:"size" validate:"required"`
	Quantity    uint32  `json:"quantity" validate:"required"`
}
