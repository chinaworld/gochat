package client

import (
	"testing"
	"net"
	"fmt"
	"time"
	"strconv"
	"bufio"
	"gochat/tool"
	"strings"
	"encoding/json"
)

const (
	addr = "127.0.0.1:8888"
)

type loginData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Msg struct {
	UserId     int    `json:"user_id"`
	FormUserId int    `json:"form_user_id"`
	Content    string `json:"content"`
}

func main() {
	con, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	buf := bufio.NewReader(con)
	var user_name string
	var password string
	fmt.Println("please login")
	fmt.Println("please enter your user name")
	fmt.Scanln(&user_name)
	fmt.Println("please enter your user password")
	fmt.Scanln(&password)

	login := loginData{UserName: user_name, Password: password}
	data, _ := json.Marshal(login)
	sms := []byte(data)
	sms = append(sms, '\n')
	con.Write(sms)

	id_string, err := buf.ReadString('\n')
	id_strings := strings.Split(id_string, "\n")
	id, _ := strconv.Atoi(id_strings[0])

	fmt.Println("链接成功  你的ID为：", id)
	msg := []string{}

	go func() {
		for {

			m, err := buf.ReadString('\n')
			if err != nil {
				fmt.Println("a err:", err.Error())
			}
			msg = append(msg, m)
		}
	}()

	go func() {
		for {
			h := make([]byte, 1)
			h[0] = '\n'
			con.Write(h)
			time.Sleep(3 * time.Second)
		}
	}()

	var cmd string
	for {

		fmt.Scanln(&cmd)

		switch cmd {

		case "msg":
			if len(msg) > 0 {
				fmt.Println(msg[len(msg)-1])
			} else {
				fmt.Println("have noting1")
			}
			cmd = ""
		case "msg -all":
			for v := range msg {
				fmt.Println(v)
			}
			cmd = ""

		case "push":
		Again:
			user_id_string := ""
			context := ""
			fmt.Println("please enter other's id ")
			fmt.Scanln(&user_id_string)
			user_id_int, err := strconv.Atoi(user_id_string)
			if err != nil {
				fmt.Println("enter have a err! please enter again")
				goto Again
			}
			fmt.Println("please enter context")
			fmt.Scanln(&context)

			m := Msg{UserId: id, FormUserId: user_id_int, Content: context}
			data, err := json.Marshal(m)
			if err != nil {
				fmt.Println("enter have a err! please enter again")
			}
			data = append(data, '\n')
			_, err = con.Write(data)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OJBK")
			}
			cmd = ""

		case "help":
			fmt.Println("msg : 获取最后一条消息")
			fmt.Println("msg -all: 获取所有消息")
			fmt.Println("push : 准备发送一条消息")
			fmt.Println("提供了若干测试账号 \n id：2108	username：0	password:0 ")
			fmt.Println("id：2109	username：1	password:1 ")
			fmt.Println("以此类推")

		default:
			fmt.Println("you can enter 'help' get help")
		}

	}

}

func Test_cli_main(t *testing.T) {

	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("连接服务端失败:", err.Error())
			return
		}
		fmt.Println("已连接服务器")

		go Client(conn, strconv.Itoa(i), strconv.Itoa(i), i)
	}
	time.Sleep(time.Hour)
}

func Client(conn net.Conn, username string, password string, i int) {
	//登录

	login := loginData{UserName: username, Password: password}
	data, _ := json.Marshal(login)
	sms := []byte(data)

	sms = append(sms, '\n')
	conn.Write(sms)

	go func() {
		for {
			//发送消息
			msg := tool.Msg{UserId: 2108 + i, FormUserId: 2108, Content: password}
			data, _ := json.Marshal(msg)

			conn.Write(append([]byte(data), '\n'))
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			defer conn.Close()
			//接受消息
			buf := bufio.NewReader(conn)
			c, err := buf.ReadString('\n')
			if err != nil {
				fmt.Println("读取服务器数据异常:", err.Error())
				break
			}
			fmt.Println(c)
		}

	}()

}
