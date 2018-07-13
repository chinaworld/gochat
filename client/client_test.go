package client

import (
	"testing"
	"net"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"time"
	"gochat/tool"
)

const (
	addr = "127.0.0.1:8888"
)

func Test_cli_main(t *testing.T) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	Client(conn)
}

type loginData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func Client(conn net.Conn) {
	sms := make([]byte, 128)

	login := loginData{UserName: "1", Password: "1"}
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
		msg := tool.Msg{UserId: 2, FormUserId: 2, Content: "this is a test"}
		data, _ := json.Marshal(msg)

		conn.Write(append([]byte(data), '\n'))

		buf := make([]byte, 128)
		c, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
		}
		fmt.Println(string(buf[0:c]))

		time.Sleep(3 * time.Second)

	}

}
