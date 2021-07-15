package model

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"
)

// 不用默认的 gorm.Model, 而是自己定义，加深理解
type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female 表示 女， male表示男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户，2表示管理员'"`
}

func (user *User) GetMd5Str(code string) (string, error) {
	// 生成md5的值
	m := md5.New()
	_, err := io.WriteString(m, code)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}

