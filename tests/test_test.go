package tests

import "testing"
import (
	"gochat/config"
	"fmt"
	"gochat/model"
	"reflect"
)

func Test01(t *testing.T) {

	sql_link := config.Db_user + ":" + config.Db_password + "@tcp(" + config.Db_host +
		":" + config.Db_port + ")/" + config.Db_DB + "?charset=utf8"
	fmt.Println(sql_link)
}

type User struct {
	Id       int    `db:"id"`
	UserName string `db:"user_name"`
	PassWord string `db:"password"'`
	Time     string `db:"time"'`
}

func (*User) GetTableName() string {
	return "user"
}
func TestDb(t *testing.T) {

	user := []User{}
	//sql := "select * from user where user_name = ? "
	//model.Query(sql, &user, "dadfaf")
	sql := "select * from user "

	model.Querys(sql, &user)
	fmt.Println(user)
}


func TestRef(t *testing.T) {

	a := []User{}
	ty := reflect.TypeOf(&a).Elem()
	//ty1 := reflect.PtrTo(ty)
	//v := reflect.New(ty)

	//	ty := v.Type()

	//v = reflect.Indirect(v)

	fmt.Println(ty.Kind())
	//fmt.Println(v.Elem().NumField())
}