package client

import (
	"testing"
	"net"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"time"
	"gochat/tool"
	"strconv"
)

const (
	addr = "127.0.0.1:8888"
)

type loginData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func Test_cli_main(t *testing.T) {

	for i := 4; i < 1000; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("连接服务端失败:", err.Error())
			return
		}
		fmt.Println("已连接服务器")
		defer conn.Close()
		go Client(conn, strconv.Itoa(i), strconv.Itoa(i))
	}
	time.Sleep(time.Hour)
}

func Client(conn net.Conn, username string, password string) {
	sms := make([]byte, 128)

	login := loginData{UserName: username, Password: password}
	data, _ := json.Marshal(login)

	sms = []byte(data)
	fmt.Print("请输入要发送的消息:")
	//	_, err := fmt.Scan(&sms)
	//	if err != nil {
	//		fmt.Println("数据输入异常:", err.Error())
	//	}

	sms = append(sms, '\n')
	conn.Write(sms)

	for {
		msg := tool.Msg{UserId: 1006, FormUserId: 1, Content: password}
		data, _ := json.Marshal(msg)

		conn.Write(append([]byte(data), '\n'))

		//time.Sleep(3 * time.Second)
		//
		go func() {
			buf := make([]byte, 128)
			c, err := conn.Read(buf)
			if err != nil {
				fmt.Println("读取服务器数据异常:", err.Error())
			}
			fmt.Println(string(buf[0:c]))
		}()
	}

}

func Test_cli_main2(t *testing.T) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	go Client(conn, strconv.Itoa(2), strconv.Itoa(2))
	time.Sleep(time.Hour)
}

func Test_cli_main3(t *testing.T) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	go Client(conn, strconv.Itoa(3), strconv.Itoa(3))
	time.Sleep(time.Hour)
}

func Test_cli_main4(t *testing.T) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	go Client(conn, strconv.Itoa(4), strconv.Itoa(4))
	time.Sleep(time.Hour)
}
