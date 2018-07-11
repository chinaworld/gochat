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
	tool.UserMap{}

	for {

		con, err := tcp_listen.AcceptTCP()
		if err != nil {
			fmt.Println("链接失败")
			continue
		}
		//conChan <- &con

		go ConHandler(con)
		//defer func() {
		//	con.Close()
		//	fmt.Println("链接已经关闭")
		//}()

		//data := make([]byte, 1000)
		//fmt.Println("this conn from the :", con.LocalAddr().String())

		//go func() {
		//	for {
		//
		//		reschan := make(chan []byte, 1)
		//		go func() {
		//
		//			con.Read(data)
		//			fmt.Println(string(data))
		//
		//			if data != nil {
		//				reschan <- data
		//				data = nil
		//			}
		//
		//		}()
		//
		//		select {
		//		case data := <-reschan:
		//			//doData(data)
		//			msg := tool.Msg{MsgByte: data}
		//			msg.InitMsg()
		//		case <-time.After(time.Second * 3):
		//			fmt.Println("request time out")
		//		}
		//		//err = msg.InitMsg()
		//		if err != nil {
		//			fmt.Println(err.Error())
		//			continue
		//		}
		//		_, err = con.Write([]byte("im fine"))
		//		if err != nil {
		//			con.Close()
		//			break
		//		}
		//	}
		//}()
	}
}
