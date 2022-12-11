package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);comment:'ユーザー名';"`
	Password string `gorm:"type:varchar(100);comment:'パスワード';"`
	Email string `gorm:"type:varchar(100);comment:'メールアドレス';"`
	Guest_flag uint `gorm:"type:tinyint(1);default:1;not null;comment:'ゲストフラグ。1はゲスト。';"`
	User_family_infos []User_family_info
}