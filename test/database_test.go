package test

import (
	"fmt"
	"testing"
)

type User struct {
	Id              int64
	Name            string
	Avatar          string
	BackgroundImage string
	Signature       string
	Username        string
	Password        string
}

func (User) TableName() string {
	return "sys_user"
}

func Test_userInsert(t *testing.T) {
	db := GetDbConnection2()

	user := User{int64(10), "aa", "aa", "aa", "aa", "aa", "aa"}

	res := db.Create(&user)

	fmt.Println(res)
}
