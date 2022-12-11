package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type My_recipe struct {
	gorm.Model
	UserID uint `gorm:"type:int(11);not null;comment:'ユーザーID';"`
	RecipeID uint `gorm:"type:int(11);not null;comment:'レシピID';"`
	Date string `gorm:"type:date;comment:'日付';"`
	Index uint `gorm:"type:tinyint(1);comment:'レシピ表示順';"`
	Recipe Recipe
}