package tests

import "testing"
import (
	"gochat/config"
	"fmt"
	"reflect"
	"strconv"
	"gochat/model"
	"gochat/tool"
	"runtime"
)

func Test01(t *testing.T) {

	sql_link := config.Db_user + ":" + config.Db_password + "@tcp(" + config.Db_host +
		":" + config.Db_port + ")/" + config.Db_DB + "?charset=utf8"
	fmt.Println(sql_link)
}

type User struct {
	Id         int    `db:"id"`
	UserName   string `db:"user_name"`
	PassWord   string `db:"password"'`
	CreateTime int64  `db:"create_time"`
}

func (*User) GetTableName() string {
	return "user"
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

func TestLog(t *testing.T) {


	tool.LogDebug, _ = tool.NewLog()

	tool.LogDebug.Println("11")

	//panic("123123")

	//tool.LogDebug.Println(runtime.Caller(0))

		fmt.Println("11111")
}

func TestDb(t *testing.T) {

	//	user := User{UserName: "lishuyang", PassWord: "123123123", CreateTime: time.Now().Unix()}

	//sql := "select * from user where user_name = ? "
	//model.Query(sql, &user, "dadfaf")
	//	sql := "select * from user where id = ? and user_name = ?"
	//sql := "select * from where user_name = ? and password = ? "
	//sql := "select * from user where user_name = ? and password = ? "
	//sql := "select * from user where user_name = ?  "

	//model.Query(sql, &user, "1")
	//fmt.Println(user)
	//model.Insert(&user)
	//h, err := model.GetHistoricalMsg(2)
	//fmt.Println(len(h), err)

	for i := 1001; i < 50000; i++ {
		user := User{UserName: strconv.Itoa(i), PassWord: strconv.Itoa(i)}
		model.Insert(&user)
	}
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
