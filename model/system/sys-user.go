package system

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Age      uint   `json:"age"`
	Sex      uint8  `gorm:"comment:0:男, 1:女" json:"sex"`
}
