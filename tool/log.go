package tool

import (
	"time"
	"os"
	"log"
	"runtime/debug"
)

var LogDebug *Logdebug

type Logdebug struct {
	Log   *log.Logger
	Debug bool
}

func NewLog() (*Logdebug, *os.File) {
	log.SetFlags(3)
	now_time := time.Now().Format("2006-01-02")
	file_name := now_time + "-log.log"

	file, err := os.Create(file_name)
	if err != nil {
		panic(err.Error())
		//return nil
	}

	//defer file.Close()
	debuglog := log.New(file, "[DEBUG]", log.LstdFlags)

	logdebug := Logdebug{Log: debuglog, Debug: true}

	return &logdebug, file
}

func (this *Logdebug) Println(v ...interface{}) {

	stack := string(debug.Stack())
	if this.Debug {
		log.Println(v)
		log.Println(stack)

	}
	this.Log.Println(v)
	this.Log.Println(stack)
}
