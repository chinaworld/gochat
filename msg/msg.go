package msg

import (
	"encoding/json"
	"net"
	"fmt"
	"gochat/model"
	"time"
)

type Msg struct {
	UserId     int    `json:"user_id"`
	FormUserId int    `json:"form_user_id"`
	Content    string `json:"content"`
}

func (this *Msg) InitMsg(data []byte) error {
	if err := json.Unmarshal(data, this); err != nil {
		return err
	}
	return nil
}

type ReturnMsg struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

func (this *ReturnMsg) LoginErroMsg() ([]byte, error) {
	this.UserId = 0xff
	this.Content = "登录失败，请检查你的用户名和密码"

	data, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *Msg) RelayMsg(id_map *UserMap) bool {

	con_interface, ok := id_map.Map.Load(this.FormUserId)
	if !ok {
		//todo  持久化数据  等待下次上线再去拉
		histor := model.HistoricalMsg{UserId: this.FormUserId, SendUser: this.UserId, Context: this.Content, CreateTime: time.Now().Unix()}
		model.Insert(&histor)
		fmt.Println("用户不在线")
		return false
	}
	con := con_interface.(*net.TCPConn)
	data := append([]byte(this.Content), '\n')
	con.Write(data)
	return true
}
