package tool

import "sync"

type UserMap struct {
	Map sync.Map
}

func (this *UserMap) InitMap(){
	//this.Map =
}