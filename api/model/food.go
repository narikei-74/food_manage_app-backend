package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Food struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;comment:'食材名';"`
	Amazon_id string `gorm:"type:varchar(100);not null;comment:'amazonの商品ID';"`
	Gram uint `gorm:"type:int(11);comment:'食材の量（グラム）';"`
	Quantity uint `gorm:"type:int(11);comment:'食材の量（数）';"`
	Price uint `gorm:"type:int(11);not null;comment:'食材の値段';"`
}