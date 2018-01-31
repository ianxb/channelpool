package channelpool

import (
	"sync"
	"errors"
)

type Pool interface {
	Get() (chan interface{}, error)
	Put(chan interface{}) error
	Len() int32
	GetCurrentChannelCount() int32
	Close()error
}

type myPool struct {
	maxCount int32
	lock sync.RWMutex
	count int32
	ch chan Channel
}

func Init(){

}

func NewPool(size int32) (Pool,error){
	if size == 0{
		return nil,errors.New("this size of NewPool is 0")
	}
	ch := make(chan Channel,size )
	return &myPool{
		ch:ch,
		maxCount:size,
		count:0,
		},nil
}

func (this *myPool)Get()(chan interface{},error){
	this.lock.RLock()
	defer this.lock.RUnlock()
	if this.count == 0{
		return nil, errors.New("the count of channel in pool is 0")
	}

}