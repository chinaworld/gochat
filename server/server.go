package server

import (
	"gochat/config"
	"net"
	"fmt"
	"gochat/tool"
	"gochat/model"
)

func GetServer(config *config.Config) *net.TCPListener {

	tcp_listen, err := net.ListenTCP("tcp", config.Addr)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return tcp_listen

}

func ServerRun(tcp_listen *net.TCPListener) error {
	id_map := &tool.UserMap{}

	for {
		logindata := make([]byte, 100)
		con, err := tcp_listen.AcceptTCP()
		if err != nil {
			fmt.Println(err.Error())
			con.Close()
		}
		_, err = con.Read(logindata)
		if err != nil {
			fmt.Println(err.Error())
			con.Close()
		}

		key := model.Login(logindata)
		if !key {
			tool.Msg{U}
			con.Write()
			continue
		}

		if err != nil {
			fmt.Println("链接失败")
			continue
		}
		go ConHandler(con, id_map)

	}
}
