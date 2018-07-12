package model

import (
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
)

type User struct {
	Id         int    `db:"id"`
	UserName   string `db:"user_name"`
	Password   string `db:"password"`
	CreateTime int    `db:"create_time"`
}

type loginData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (*User) GetTableName() string {
	return "user"
}

func Login(logindata []byte) bool {
	data := loginData{}
	if err := json.Unmarshal(logindata, &data); err != nil {
		fmt.Println(err.Error())
		return false
	}



}
