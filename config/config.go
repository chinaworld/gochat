package config

import "net"

type Config struct {
	Addr *net.TCPAddr
	Ip   net.IP
	Port int
}

var Db_user = "root"
var Db_password = "0"
var Db_host = "0"
var Db_port = "3306"
var Db_DB = "gochat"

func (this *Config) SetConfig() {
	this.Addr = &net.TCPAddr{this.Ip, this.Port, ""}
}
