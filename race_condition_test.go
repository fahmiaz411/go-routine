package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRace(t *testing.T){
	x := 0

	for i:= 0; i < 1000; i++ {
		go func (){
			for i:= 0; i < 100; i++{
				x++
			}
			fmt.Println(x)
		}()
	}

	time.Sleep(time.Second * 5)

	fmt.Println(x)
}
