package server

import (
	"gochat/config"
	"net"
	"fmt"
	"gochat/tool"
	"github.com/astaxie/beego/orm"
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

		con, err := tcp_listen.AcceptTCP()
		if err != nil {
			fmt.Println("链接失败")
			continue
		}
		go ConHandler(con, id_map)

	}

	orm.NewOrm().Raw().QueryRows()
}
