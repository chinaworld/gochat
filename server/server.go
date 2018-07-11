package server

import (
	"gochat/config"
	"net"
	"fmt"
	"gochat/tool"
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

	for {
		//tool.Heartbeat(tcp_listen)
		con, err := tcp_listen.Accept()
		if err != nil {
			fmt.Println("链接失败")
			continue
		}
		//defer func() {
		//	con.Close()
		//	fmt.Println("链接已经关闭")
		//}()
		data := make([]byte, 1000)
		fmt.Println("this conn from the :", con.LocalAddr().String())

		go func() {
			for {

				//select {
				//case data := <-resChan:
				//	doData(data)
				//case <-time.After(time.Second * 3):
				//	fmt.Println("request time out")
				//}
				con.Read(data)
				fmt.Println(string(data))
				msg := tool.Msg{MsgByte: data}
				data = nil
				err = msg.InitMsg()
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				_, err = con.Write([]byte("im fine"))
				if err != nil {
					con.Close()
					break
				}
			}
		}()
	}
}
