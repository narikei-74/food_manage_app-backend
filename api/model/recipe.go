package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Recipe struct {
	gorm.Model
	Name string `gorm:"type:varchar(30);not null;comment:'レシピ名';"`
	Rakuten_id string `gorm:"type:varchar(100);comment:'楽天のレシピID';"`
	Cooking_time uint `gorm:"type:int(11);comment:'調理時間（分）';"`
	Image_key string `gorm:"type:varchar(100);not null;comment:'レシピ画像のキー';"`
	Dish_category uint `gorm:"type:tinyint(1);not null;comment:'栄養バランスカテゴリ。1が「主食」、2が「主菜」、3が「副菜」、4が「汁物」。';"`
	Rakuten_url string `gorm:"type:varchar(100);comment:'楽天のレシピURL';"`
	Rakuten_flag uint `gorm:"type:varchar(100);defalut:1;not null;comment:'楽天レシピフラグ。0が「楽天レシピ';"`
	Recipe_materials []Recipe_material
	Recipe_categories []Recipe_category
}