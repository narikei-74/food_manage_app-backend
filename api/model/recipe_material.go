package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Recipe_material struct {
	gorm.Model
	RecipeID uint `gorm:"type:int(11);not null;comment:'レシピID';"`
	FoodID uint `gorm:"type:int(11);not null;comment:'食材ID';"`
	grams uint `gorm:"type:int(11);comment:'食材の量（グラム）';"`
	quantity uint `gorm:"type:int(11);comment:'食材の量（数）';"`
	quantity_label string `gorm:"type:varchar(20);comment:'食材の量（ラベル）';"`
	Food Food
}