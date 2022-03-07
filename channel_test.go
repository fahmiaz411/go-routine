package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)
	
	go (func (){
		channel <- "fahmi"
		fmt.Println("selesai kirim data")	
	})()

	data := <- channel
	fmt.Println(data)

}

func ChanParam (c chan string){
	time.Sleep(5 * time.Second)
	c <- "test param"
}

func TestParam (t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go ChanParam(channel)

	data := <- channel

	fmt.Println(data)
}

func BenchmarkParam(b *testing.B) {
	for i:=0; i< b.N; i++{
		channel := make(chan string)
		defer close(channel)

		go ChanParam(channel)

		data := <- channel

		fmt.Println(data)
	}
}