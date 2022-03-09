package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T){
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}
func TestAfterTimer(t *testing.T){
	timer := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer
	fmt.Println(time)
}

func TestAfterFuncTimer (t *testing.T){

	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(time.Second, func(){
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}

func TestTicker(t *testing.T){
	ticker := time.NewTicker(time.Second)

	count := 0

	for {
		if count < 5 {
			fmt.Println(<-ticker.C)
			count++
		} else {
			ticker.Stop()
			break
		}
	}
}

func TestTick(t *testing.T){
	tick := time.Tick(time.Second)

	count := 0

	for {
		if count < 5 {
			fmt.Println(<-tick)
			count++
		} else {
		}
	}
}