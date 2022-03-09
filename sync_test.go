package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup, i int){
	defer group.Done()
	group.Add(1)

	fmt.Println("Hello", i)
}

func TestWaitGroup(t *testing.T){
	group := &sync.WaitGroup{}

	for i:= 0; i < 100; i++{
		go RunAsync(group, i+1)
	}

	group.Wait()
}

var counter = 0

func OnlyOnce(){
	counter++
}

func TestOnce(t *testing.T){
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i:= 0; i < 100; i++{
		go func (){
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}

func TestPool (t *testing.T){
	pool := sync.Pool{
		New: func() interface{} {
			return "N"
		},
	}

	pool.Put("fahmi")
	pool.Put("ega")
	pool.Put("soni")

	for i:= 0; i < 10; i++{
		go func(){
			data:= pool.Get()
			fmt.Println(data)
			// pool.Put(data)
		}()
	}

	time.Sleep(time.Second)
}

func AddToMap (data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(value, value)
	
}

func TestMap(t *testing.T){
	data:= &sync.Map{}
	group := &sync.WaitGroup{}

	for i:=0; i<100; i++{
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, value		)
		return true
	})
}

var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCondition(value int){
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T){
	for i:= 0; i < 10; i++{
		go WaitCondition(i)
	}
	
	go func(){
		for i:= 0; i < 10; i++{
			time.Sleep(time.Second)
			cond.Signal()
		}
	}()

	// go func(){
	// 	time.Sleep(time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}