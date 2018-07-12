package model

import (
	"fmt"
	"database/sql"
	"gochat/config"
	"reflect"
	"strconv"
)

type BaseDb interface {
	GetTableName() string
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
		fmt.Println(err.Error())
		panic(err)
	}
}

//just select one! If have more than one show the frist one
func Query(sql string, baseDb BaseDb, parse ...interface{}) (error) {

	//var column_index []reflect.Value
	stemt, err := db.Prepare(sql)
	checkErr(err)

	rows, err := stemt.Query(parse...)
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

	err = rows.Scan(dest...)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for k, v1 := range column {
		for i := 0; i < t.NumField(); i++ {
			if v1 == t.Field(i).Tag.Get("db") {

				vu := reflect.ValueOf(*dest[k].(*string)).String() //.Convert(v.Field(i).Type())
				//fmt.Println(vu)
				//check type is int or int32 ...
				if v.Field(i).Type().Kind() >= 2 && v.Field(i).Type().Kind() <= 6 {
					in, err := strconv.Atoi(vu)
					if err != nil {
						panic(err.Error())
					}
					v.Field(i).Set(reflect.ValueOf(in))
				}

				if v.Field(i).Type().Kind() == reflect.String {
					v.Field(i).Set(reflect.ValueOf(vu))
				}

				break
			}

		}

	}

	//	}

	return err
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
			return err
		}

		for k, v1 := range column {
			for i := 0; i < strv.NumField(); i++ {

				if v1 == strv.Type().Field(i).Tag.Get("db") {

					vu := reflect.ValueOf(*dest[k].(*string)).String() //.Convert(v.Field(i).Type())

					//check type is int or int32 ...
					if strv.Field(i).Type().Kind() >= 2 && strv.Field(i).Type().Kind() <= 6 {
						in, err := strconv.Atoi(vu)
						if err != nil {
							panic(err.Error())
						}
						strv.Field(i).Set(reflect.ValueOf(in))
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
