package server

import (
	"net"
	"gochat/tool"
	"bufio"
	"fmt"
	"time"
	"context"
)

func ConHandler(con *net.TCPConn, id int) {

	buf := bufio.NewReader(con)

	ctx, close := context.WithCancel(context.Background())
	dataChan := make(chan []byte, 1)
	go func(ctx2 context.Context) {

		for {
			select {
			case <-ctx2.Done():
				return
			default:
				//con.Read(data)
				data, err := buf.ReadString('\n')
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				dataChan <- []byte(data)
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
				msg := tool.Msg{}
				err := msg.InitMsg(d)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(msg)
				//todo 消息转发
				go msg.RelayMsg(id_map)

			case <-time.After(5 * time.Second):
				close()
				id_map.Map.Delete(id)
				fmt.Println(con.LocalAddr(), "心跳断开，断开连接")
				con.Close()
				return
			}
		}
	}(close)
}
