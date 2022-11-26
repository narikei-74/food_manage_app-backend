package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);"`
}