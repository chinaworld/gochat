package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gochat/config"
)

var db *sql.DB

func init() {

	var err error
	sql_link := config.Db_user + ":" + config.Db_password + "@tcp(" + config.Db_host +
		":" + config.Db_port + ")/" + config.Db_DB + "?charset=utf8"
	db, err = sql.Open("mysql", sql_link)
	checkErr(err)

}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func query() {
	db, err := sql.Open("mysql", "root:@/shopvisit")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM shopvisit.announcement")
	checkErr(err)

	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		//
		//for i := range values {
		//	scanArgs[i] = &amp;
		//	values[i]
		//}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	rows.Close()
}
