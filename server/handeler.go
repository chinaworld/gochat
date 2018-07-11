package server

import (
	"net"

	"gochat/tool"
	"time"
	"context"
	"fmt"
)

func ConHandler(con *net.TCPConn) {

	data := make([]byte, 1000)

	ctx, close := context.WithCancel(context.Background())
	dataChan := make(chan []byte, 1)
	go func(ctx2 context.Context) {

		for {

			select {
			case <-ctx2.Done():
				return
			default:
				con.Read(data)
				if data != nil {
					dataChan <- data
					data = nil
				}
			}

		}
	}(ctx)

	go func(cancelFunc context.CancelFunc) {
		for {
			select {
			case d := <-dataChan:
				if len(d) < 1 {
					continue
				}
				msg := tool.Msg{MsgByte: d}
				err := msg.InitMsg()
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(string(d))

			case <-time.After(5 * time.Second):
				close()
				con.Close()
				return
			}
		}
	}(close)

}
