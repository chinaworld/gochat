package main

import (
	"gochat/config"
	"net"
	"gochat/server"
)

func main() {
	config := config.Config{}
	config.Port = 8888
	config.Ip = net.ParseIP("127.0.0.1")
	config.SetConfig()
	tcp_listen := server.GetServer(&config)
	server.ServerRun(tcp_listen)
}
