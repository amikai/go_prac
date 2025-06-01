package syncex

import "sync"

type Something struct {
	Count int
}

var onceF = sync.OnceValue(func() *Something {
	sth := &Something{}
	sth.Count++
	return sth
})

func GetInstance() *Something {
	return onceF()
}
