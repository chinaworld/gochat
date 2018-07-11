package tests

import "testing"
import (
	"gochat/config"
	"fmt"
)

func Test01(t *testing.T) {

	sql_link := config.Db_user + ":" + config.Db_password + "@tcp(" + config.Db_host +
		":" + config.Db_port + ")/" + config.Db_DB + "?charset=utf8"
	fmt.Println(sql_link)
}
