package config

import "net"

type Config struct {
	Addr *net.TCPAddr
	Ip   net.IP
	Port int
}

func (this *Config) SetConfig() {
	this.Addr = &net.TCPAddr{this.Ip, this.Port, ""}
}
