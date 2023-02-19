package model

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       int32  `json:"uid" orm:"auto"`
	Email    string `json:"email" orm:"size(100)"`
	Password string `json:"pw" orm:"size(128)"`
	Name     string `json:"name" orm:"size(50)"`
}

func (u *User) EncryptPassword() {
	checksum := sha256.Sum256([]byte(u.Password))
	u.Password = hex.EncodeToString(checksum[:])
}

type UserDetail struct {
	User
}

func init() {
	orm.RegisterModel(new(User))
}
