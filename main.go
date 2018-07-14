package main

import (
	"gochat/config"
	"net"
	"gochat/server"
	"runtime"
	"gochat/tool"
	"os"
)

func main() {
	var logfile *os.File

	runtime.GOMAXPROCS(runtime.NumCPU())
	config := config.Config{}
	config.Port = 8888
	config.Ip = net.ParseIP("127.0.0.1")
	config.SetConfig()
	tcp_listen := server.GetServer(&config)

	tool.LogDebug, logfile = tool.NewLog()
	defer logfile.Close()

	server.ServerRun(tcp_listen)
}
