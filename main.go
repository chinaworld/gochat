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

	tool.LogDebug, logfile = tool.NewLog()
	runtime.GOMAXPROCS(runtime.NumCPU())
	config := config.Config{}

	defer func() {
		if err := recover(); err != nil {
			tool.LogDebug.Println(err) //这里的err其实就是panic传入的内容，55
		}
	}()
	config.Port = 8888
	config.Ip = net.ParseIP("localhost")
	config.SetConfig()
	tcp_listen := server.GetServer(&config)

	defer logfile.Close()

	server.ServerRun(tcp_listen)
}
