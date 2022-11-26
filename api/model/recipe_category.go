package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Recipe_category struct {
	gorm.Model
	RecipeID uint `gorm:"type:int(11);not null;comment:'レシピID';"`
	Category_name string `gorm:"type:varchar(50);not null;comment:'カテゴリ名';"`
}