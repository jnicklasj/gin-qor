package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Description string `json:"description" gorm:"text;not null;`
}
