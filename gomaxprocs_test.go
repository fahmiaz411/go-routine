package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T){
 
	for i:= 0; i<10; i++{
		go func(){
			time.Sleep(time.Second)
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println(totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println(totalGoroutine)

}