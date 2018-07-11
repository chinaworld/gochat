package tool

import (
	"net"
	"fmt"
)

func Heartbeat(con *net.TCPListener) {
	//go func() {
	//
	//	time.After()
	//}()
	fmt.Println("心跳")
}
