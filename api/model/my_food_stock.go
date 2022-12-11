package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type My_food_stock struct {
	gorm.Model
	UserID uint `gorm:"type:int(11);not null;comment:'ユーザーID';"`
	FoodID uint `gorm:"type:int(11);not null;comment:'食材ID';"`
	Gram uint `gorm:"type:int(11);comment:'食材の量（グラム）';"`
	Quantity float32 `gorm:"type:float;comment:'食材の量（数）';"`
	Food Food
}