package singleton

import "sync"

var once sync.Once

type Something struct{}

var instance *Something

func GetInstance() *Something {
	once.Do(func() {
		instance = &Something{}
	})
	return instance
}

func UnsafeGetInstance() *Something {
	if instance == nil {
		instance = &Something{}
	}
	return instance
}
