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

	ctx, close := context.WithCancel(context.Background())
	dataChan := make(chan []byte, 100)
	go func(ctx2 context.Context) {
		buf := bufio.NewReader(con)
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
				fmt.Println(id, "用户", msg)
				//todo 消息转发
				key := msg.RelayMsg(id_map)
				if !key {
					//con.Write([]byte("用户不在线"))
				}

			case <-time.After(10 * time.Second):
				close()
				id_map.Map.Delete(id)
				fmt.Println(con.LocalAddr(), "心跳断开，断开连接")
				con.Close()
				return
			}
		}
	}(close)
}
