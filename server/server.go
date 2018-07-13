package server

import (
	"gochat/config"
	"net"
	"fmt"
	"gochat/tool"
	"bufio"
	"gochat/model"
	"github.com/gin-gonic/gin/json"
)

var id_map = &tool.UserMap{}

func GetServer(config *config.Config) *net.TCPListener {

	tcp_listen, err := net.ListenTCP("tcp", config.Addr)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		return nil
	}
	return tcp_listen

}

func ServerRun(tcp_listen *net.TCPListener) error {

	for {
		con, err := tcp_listen.AcceptTCP()
		if err != nil {
			fmt.Println(err.Error())
		}
		login_buf := bufio.NewReader(con)
		message, err := login_buf.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			con.Close()
		}

		id, key := model.Login([]byte(message))
		if !key {
			retuenmsg := tool.ReturnMsg{}
			data, err := retuenmsg.LoginErroMsg()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println("登录失败")
			con.Write(data)
			continue
		}

		fmt.Println("登录成功 建立连接", con.RemoteAddr())
		id_map.Map.Store(id, con)

		go func(gocon *net.TCPConn) {
			h, err := model.GetHistoricalMsg(id)
			//拉历史消息
			if len(h) > 0 {
				if err != nil {
					fmt.Println(err.Error())
				}
				lr := model.LoginReturn{Id: id, Histor: h}

				lrb, err := json.Marshal(lr)
				if err != nil {
					fmt.Println(err.Error())
				}
				gocon.Write(append(lrb, '\n'))
			}
		}(con)

		go ConHandler(con, id)

	}
}
