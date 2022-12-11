package model

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User_family_info struct {
	gorm.Model
	UserID uint `gorm:"type:int(11);not null;comment:'ユーザーID。ユーザーテーブルのプライマリーキー。';"`
	Sex uint `gorm:"type:tinyint(1);not null;default:0;comment:'性別。0が「無回答」、1が「男性」、2が「女性」。';"`
	Age string `gorm:"type:varchar(10);comment:'年代';"`
	Master_flag uint `gorm:"type:tinyint(1);not null;comment:'アプリ利用者フラグ。1が「アプリ利用者」、0が「その他」。';"`
}