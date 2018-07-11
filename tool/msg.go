package tool

import (
	"strings"
	"fmt"
	"errors"
	"strconv"
)

type Msg struct {
	MsgByte    []byte
	UserId     int
	FormUserId int
	Content    string
}

func (this *Msg) InitMsg() error {
	str := strings.Split(string(this.MsgByte), "@")
	var err error
	if len(str) < 2 {
		fmt.Println("数据错误")
		return errors.New("数据错误")
	}

	data2 := []byte(str[1])
	this.Content = string(data2)

	userinfo := strings.Split(str[0], ":")

	this.UserId, err = strconv.Atoi(userinfo[0])
	if err != nil {
		return err
	}

	this.FormUserId, err = strconv.Atoi(userinfo[1])
	if err != nil {
		return err
	}

	return nil
}
