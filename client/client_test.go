package client

import (
	"testing"
	"net"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"time"
	"strconv"
	"bufio"
	"gochat/tool"
)

const (
	addr = "127.0.0.1:8888"
)

type loginData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
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
