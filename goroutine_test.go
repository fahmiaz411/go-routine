package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHello() {
	for i:=0; i < 10000; i++ {

		fmt.Println(i, "Hello routine")
	}
}

func TestRunHello(t *testing.T){
	go RunHello()
	fmt.Println("Hello test")
	time.Sleep(time.Second / 2)
}

func DisplayNum (n int){
	fmt.Println(n)
}

func TestMemoryLeak(t *testing.T){
	for i:= 0; i < 100000; i++{
		go DisplayNum(i)
	}
}