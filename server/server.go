package server

import (
	"gochat/config"
	"net"
	"gochat/tool"
	"bufio"
	"gochat/model"
	"strconv"
	"encoding/json"
	"gochat/msg"
	"time"
)

var id_map = &msg.UserMap{}

func GetServer(config *config.Config) *net.TCPListener {

	tcp_listen, err := net.ListenTCP("tcp", config.Addr)
	if err != nil {
		tool.LogDebug.Println(err)
		panic(err)
		return nil
	}
	return tcp_listen

}

func ServerRun(tcp_listen *net.TCPListener) error {

	for {
		con, err := tcp_listen.AcceptTCP()
		if err != nil {
			tool.LogDebug.Println("[Warning]", err)
		}
		login_buf := bufio.NewReader(con)
		message_chan := make(chan string, 1)

		//防止有傻逼在这个地方不登录  占用老子的资源
		go func() {
			message, err := login_buf.ReadString('\n')
			if err != nil {
				tool.LogDebug.Println("[Err]", err)
				con.Close()
			}
			message_chan <- message
		}()

		select {
		case <-time.After(60 * time.Second):
			con.Close()

		case message := <-message_chan:


			id, key := model.Login([]byte(message))
			if !key {
				retuenmsg := msg.ReturnMsg{}
				data, err := retuenmsg.LoginErroMsg()
				if err != nil {
					tool.LogDebug.Println(err)
					continue
				}
				tool.LogDebug.Println("[Warning]", con.RemoteAddr(), "登录失败")
				con.Write(data)
				continue
			}
			tool.LogDebug.Println("[Warning]", "登录成功 建立连接", con.RemoteAddr())
			id_string := strconv.Itoa(id)

			con.Write(append([]byte(id_string), '\n'))

			id_map.Map.Store(id, con)

			go func(gocon *net.TCPConn) {
				h, err := model.GetHistoricalMsg(id)
				//拉历史消息
				if len(h) > 0 {
					if err != nil {
						tool.LogDebug.Println("[Err]", err)
					}
					lr := model.LoginReturn{Id: id, Histor: h}

					lrb, err := json.Marshal(lr)
					if err != nil {
						tool.LogDebug.Println("[Err]", err)
					}
					gocon.Write(append(lrb, '\n'))
				}
			}(con)

			go ConHandler(con, id)
		}

	}
}
