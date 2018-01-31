package channelpool

import (
	"testing"
	//"log"
)

func TestNewChannel(t *testing.T) {
	size := []int{1,2,3,4,5,6}
	for _,i := range size{
		_, err := NewChannel(int32(i))
		if err!=nil{
			t.Fatalf("error in new channel:%v",err)
		}

	}
}

func TestMyChannel_Get(t *testing.T) {
	a := make(chan interface{},5)
	a<-3
	b:= make(chan interface{},3)
	b<-"hello"
	b<-"world"
	data := []myChannel{
		{ch:a},
		{ch:b},
	}
	result,err := data[0].Get()
	if err!=nil{
		t.Fatalf("get func error:%v",err)
	}
	if result != 3 {
		t.Fatalf("get result error:%v",result)
	}
	resultstr, err := data[1].Get()
	if err!=nil{
		t.Fatalf("get func error:%v",err)
	}
	if resultstr != "hello" {
		t.Fatalf("get result error:%v",resultstr)
	}
	resultstr1, err := data[1].Get()
	if err!=nil{
		t.Fatalf("get func error:%v",err)
	}
	if resultstr1 != "world" {
		t.Fatalf("get result error:%v",resultstr1)
	}

}