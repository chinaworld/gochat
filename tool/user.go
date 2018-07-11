package tool

import "sync"

type UserMap struct {
	UserMap sync.Map
}

//func (this *UserMap)