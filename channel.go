package channelpool

import (
	"errors"
	"fmt"
	"sync"
)

type Channel interface {
	Get() (interface{}, error)
	Put(data interface{}) (ok bool)
	Len() int32
	Close()
}
type myChannel struct {
	lock sync.Mutex
	ch   chan interface{}
}

func NewChannel(size int32) (Channel, error) {
	if size == 0 {
		errMsg := fmt.Sprintf("the size of channel is:%d", size)
		return nil, errors.New(errMsg)
	}
	ch := make(chan interface{}, size)
	return &myChannel{ch: ch}, nil
}

func (this *myChannel) Get() (interface{}, error) {
	data, err := <-this.ch
	if !err {
		return nil, errors.New("get_error")
	}
	return data, nil
}

func (this *myChannel) Put(data interface{}) (ok bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	select {
	case this.ch <- data:
		ok = true
	default:
		ok = false
	}
	return
}
func (this *myChannel) Len() int32 {
	return int32(len(this.ch))
}
func (this *myChannel) Close() {
	this.lock.Lock()
	defer this.lock.Unlock()
	close(this.ch)
}
