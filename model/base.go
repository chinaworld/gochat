package model

import (
	"fmt"
	"database/sql"
	"gochat/config"
	"reflect"
	"strconv"
	"strings"
	"gochat/tool"
)

type BaseDb interface {
	GetTableName() string
}

type ConnectionPoll struct {
	max_connection  int
	connection_poll []*sql.DB
	db              *sql.DB
}

func (this *ConnectionPoll) GetDb() (*sql.DB, error) {
	if len(this.connection_poll) > 0 {
		return this.connection_poll[0], nil
	}

	sql_link := config.Db_user + ":" + config.Db_password + "@tcp(" + config.Db_host +
		":" + config.Db_port + ")/" + config.Db_DB + "?charset=utf8"
	db, err := sql.Open("mysql", sql_link)
	return db, err
}

func (this *ConnectionPoll) Init() {

}

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
		tool.LogDebug.Println("[Err]", err)
		panic(err)
	}
}

//just select one! If have more than one show the frist one
func Query(sql_string string, baseDb BaseDb, parse ...interface{}) (error) {

	//var column_index []reflect.Value
	var rows *sql.Rows
	var err error

	if parse != nil {
		stemt, err := db.Prepare(sql_string)
		checkErr(err)
		rows, err = stemt.Query(parse...)

	} else {
		rows, err = db.Query(sql_string)
		checkErr(err)
	}

	defer rows.Close()
	checkErr(err)

	column, err := rows.Columns()

	t := reflect.TypeOf(baseDb).Elem()
	v := reflect.ValueOf(baseDb).Elem()

	dest := make([]interface{}, len(column))
	row := make([]string, len(column))
	for k, _ := range dest {
		dest[k] = &row[k]
	}
	//	for i := 0; rows.Next(); i++ {
	rows.Next()
	err = rows.Scan(dest...)
	if err != nil {
		fmt.Println(err.Error())
		//return err
	}

	for k, v1 := range column {
		for i := 0; i < t.NumField(); i++ {
			if v1 == t.Field(i).Tag.Get("db") {

				vu := reflect.ValueOf(*dest[k].(*string)).String() //.Convert(v.Field(i).Type())
				//fmt.Println(vu)
				//check type is int or int32 ...
				if v.Field(i).Type().Kind() >= 2 && v.Field(i).Type().Kind() <= 6 {
					if vu != "" {

						in, err := strconv.Atoi(vu)
						if err != nil {
							panic(err.Error())
						}
						v.Field(i).Set(reflect.ValueOf(in))
					} else {
						v.Field(i).Set(reflect.ValueOf(0))
					}
				}
				if v.Field(i).Type().Kind() == reflect.String {
					v.Field(i).Set(reflect.ValueOf(vu))
				}
				break
			}

		}

	}
	return nil
}

func Querys(sql_string string, baseDb interface{}, parse ...interface{}) (error) {

	//var column_index []reflect.Value
	var rows *sql.Rows
	var err error

	if parse != nil {
		stemt, err := db.Prepare(sql_string)
		checkErr(err)
		rows, err = stemt.Query(parse...)

	} else {
		rows, err = db.Query(sql_string)
		checkErr(err)
	}

	strt := reflect.TypeOf(baseDb).Elem().Elem()
	slicev := reflect.ValueOf(baseDb).Elem()

	defer rows.Close()
	checkErr(err)

	column, err := rows.Columns()

	dest := make([]interface{}, len(column))
	row := make([]string, len(column))
	for k, _ := range dest {
		dest[k] = &row[k]
	}

	for arry_index := 0; rows.Next(); arry_index++ {
		vs := make([]reflect.Value, 0)
		strv := reflect.New(strt).Elem()

		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println(err.Error())

			//return err
		}

		for k, v1 := range column {
			for i := 0; i < strv.NumField(); i++ {

				if v1 == strv.Type().Field(i).Tag.Get("db") {

					vu := reflect.ValueOf(*dest[k].(*string)).String() //.Convert(v.Field(i).Type())

					//check type is int or int32 ...
					if strv.Field(i).Type().Kind() >= 2 && strv.Field(i).Type().Kind() <= 6 {
						if vu != "" {
							in, err := strconv.Atoi(vu)
							if err != nil {
								panic(err.Error())
							}
							strv.Field(i).Set(reflect.ValueOf(in).Convert(strv.Field(i).Type()))
						} else {
							strv.Field(i).Set(reflect.ValueOf(0))
						}
						break
					}

					if strv.Field(i).Type().Kind() == reflect.String {
						strv.Field(i).Set(reflect.ValueOf(vu))
						break
					}
					break
				}

			}

		}
		vs = append(vs, strv)
		val_arr1 := reflect.Append(slicev, vs...)
		slicev.Set(val_arr1)

	}

	return err
}

func Insert(baseDb BaseDb) (int, error) {

	//处理sql语句
	v := reflect.ValueOf(baseDb)
	var table_name string
	var clouns []string
	var values []interface{}
	//fmt.Println(v.Type().Kind())
	for i := 0; i < v.NumMethod(); i++ {
		if v.Type().Method(i).Name == "GetTableName" {
			vs := v.Method(i).Call(nil)
			table_name = vs[0].Interface().(string)
		}
	}

	for i := 0; i < v.Type().Elem().NumField(); i++ {
		c := v.Type().Elem().Field(i).Tag.Get("db")
		cs := strings.Split(c, ":")
		if len(cs) > 1 && cs[1] == "pr" {
			continue
		}
		clouns = append(clouns, c)
		values = append(values, v.Elem().Field(i).Interface())

	}

	// 用+拼接效率很低
	sql_string := "INSERT INTO " + table_name + "("
	value := "("
	for k, v := range clouns {
		sql_string = sql_string + v
		value = value + "?"
		if k+1 != len(clouns) {
			sql_string = sql_string + ","
			value = value + ","
		}

	}
	sql_string = sql_string + ")" + "VALUES " + value + ")"

	var err error

	fmt.Println(sql_string)
	fmt.Println(values)
	stemt, err := db.Prepare(sql_string)
	checkErr(err)

	defer stemt.Close()

	rows, err := stemt.Exec(values...)
	checkErr(err)

	id, err := rows.LastInsertId()
	checkErr(err)

	return int(id), err
}
