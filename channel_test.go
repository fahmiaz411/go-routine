package goroutine

import (
	"fmt"
	"strconv"
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

func ChanParam (c chan <- string){
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

func OnlyIn(channel chan <- string) {
	// time.Sleep(time.Second)
	channel <- "hey"
}

func OnlyOut(channel chan string){
	defer close(channel)
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel (t *testing.T){
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)
}

func TestBufferChannel (t *testing.T){
	channel := make(chan string, 3)
	
	go func(){
		defer close(channel)
		for i:= 1; i<=10;i++{
			channel <- "data " + strconv.Itoa(i)
		}
	}()

	for data := range channel{
		fmt.Println(data)
	}
	
	fmt.Println("end")
}

func TestSelectChannel (t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer func () {
		close(channel1)
		close(channel2)
	}()

	go OnlyIn(channel1)
	go OnlyIn(channel2)

	counter := 0

	for {
		select {
			case data:= <- channel1:
				fmt.Println("data 1", data)
				counter++
			case data := <- channel2:
				fmt.Println("data 2", data)
				counter++
			default:
				fmt.Println("wait data")
		}

		if counter == 2 {
			break
		}
	}


	fmt.Println("end")
}