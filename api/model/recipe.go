package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Recipe struct {
	gorm.Model
	Name string `gorm:"type:varchar(30);not null;comment:'レシピ名';"`
	Cooking_time uint `gorm:"type:int(11);comment:'調理時間（分）';"`
	Image_key string `gorm:"type:varchar(100);not null;comment:'レシピ画像のキー';"`
	Dish_category uint `gorm:"type:tinyint(1);not null;comment:'栄養バランスカテゴリ。1が「主食」、2が「主菜」、3が「副菜」、4が「汁物」。';"`
	UserID uint `gorm:"type:int(11);comment:'レシピ製作者。nullの場合はオリジナルレシピ。';"`
	How_to_cook string `gorm:"type:text;comment:'作り方';"`
	Private_flag uint `gorm:"type:tinyint(1);not null;comment:'プライベートフラグ。0が公開、1が非公開。'";`
	Is_ok_public uint `gorm:"type:tinyint(1);not null;comment:'公開してもしてもいいか。0が公開不可、1が公開可。'";`
	Recipe_materials []Recipe_material
	Recipe_categories []Recipe_category
}