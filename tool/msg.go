package tool

type Msg struct {
	//MsgByte    []byte
	UserId     int    `json:"user_id"`
	FormUserId int    `json:"form_user_id"`
	Content    string `json:"content"`
}

func (this *Msg) InitMsg() error {
	//str := strings.Split(string(this.MsgByte), "@")
	//var err error
	//if len(str) < 2 {
	//	fmt.Println("数据错误")
	//	return errors.New("数据错误")
	//}
	//
	//data2 := []byte(str[1])
	//this.Content = string(data2)
	//
	//userinfo := strings.Split(str[0], ":")
	//
	//this.UserId, err = strconv.Atoi(userinfo[0])
	//if err != nil {
	//	return err
	//}
	//
	//this.FormUserId, err = strconv.Atoi(userinfo[1])
	//if err != nil {
	//	return err
	//}
	//
	//return nil
}

//
//type SystemMsg struct {
//	Id       int    `json:"id"`
//	Priority int    `json:"priority"`
//	Context  string `json:"context"`
//}
