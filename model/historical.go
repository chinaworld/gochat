package model

type HistoricalMsg struct {
	Id         int    `json:"id" db:"id:pr"`
	UserId     int    `json:"user_id" db:"user_id"`
	SendUser   int    `json:"send_user" db:"send_user"`
	Context    string `json:"context" db:"context"`
	CreateTime int64  `json:"create_time" db:"create_time"`
}

func (*HistoricalMsg) GetTableName() string {
	return "historical"
}

func GetHistoricalMsg(user_id int) ([]HistoricalMsg, error) {

	h := []HistoricalMsg{}
	sql := "select * from historical where user_id = ?"
	err := Querys(sql, &h, user_id)
	return h, err
}
