package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Auto_create_recipe_settings struct {
	gorm.Model
	UserID uint `gorm:"type:int(11);not null;comment:'ユーザーID';"`
	Hate_foods string `gorm:"type:text;comment:'アレルギー・嫌いな食材。jsonでfoodIDを持つ';"`
	Is_only_rice uint `gorm:"type:tinyint(1);not null;default:0;comment:'主食を白米のみにするかどうか';"`
	Is_only_meat uint `gorm:"type:tinyint(1);not null;default:0;comment:'主菜を肉料理のみにするかどうか';"`
	Is_only_fish uint `gorm:"type:tinyint(1);not null;default:0;comment:'主菜を魚料理のみにするかどうか';"`
}