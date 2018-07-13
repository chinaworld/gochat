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

type LoginReturn struct {
	Id     int `json:"id"`
	Histor []HistoricalMsg
}

func (*User) GetTableName() string {
	return "user"
}

func Login(logindata []byte) (int, bool) {
	data := loginData{}

	fmt.Println(string(logindata))
	if err := json.Unmarshal(logindata, &data); err != nil {
		fmt.Println(err.Error())
		return 0, false
	}
	sql := "select * from user where user_name = ? and password = ? "
	user := User{}

	err := Query(sql, &user, data.UserName, data.Password)
	if err != nil {
		fmt.Println(err.Error())
		return 0, false
	}
	if user.Id == 0 {
		return 0, false
	}
	return user.Id, true

}
