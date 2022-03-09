package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for i := 0; i < 100; i++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	group.Wait()

	fmt.Println(x)
}
