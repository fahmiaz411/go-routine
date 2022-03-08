package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
	user1.Lock()
	fmt.Println("lock user1")
	user1.Change(-amount)
	
	time.Sleep(time.Second)
	
	user2.Lock()
	fmt.Println("lock user2")
	user2.Change(amount)
	
	time.Sleep(time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock (t *testing.T){
	user1:= UserBalance{
		Name: "fahmi",
		Balance: 100,
	}

	user2 := UserBalance{
		Name: "ega",
		Balance: 100,
	}

	go Transfer(&user1, &user2, 10)
	go Transfer(&user2, &user1, 10)

	time.Sleep(5 * time.Second)

	fmt.Println("user", user1.Name, "balance", user1.Balance)
	fmt.Println("user", user2.Name, "balance", user2.Balance)
}


type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (acc *BankAccount) AddBalance(amount int) {
	acc.RWMutex.Lock()
	acc.Balance = acc.Balance + amount
	acc.RWMutex.Unlock()
}

func (acc *BankAccount) GetBalance() int {
	acc.RWMutex.RLock()
	bal := acc.Balance
	acc.RWMutex.RUnlock()

	return bal
}

func TestRWMutex (t *testing.T) {
	account := BankAccount{}

	for i:= 0; i < 100; i++{
		go func (){
			for i:= 0; i < 100; i++{
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println(account.GetBalance())
}

func TestMutex(t *testing.T){
	x := 0
	var mutex sync.Mutex

	for i:= 0; i < 1000; i++ {
		go func (){
			for i:= 0; i < 100; i++{
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(time.Second * 5)

	fmt.Println(x)
}

func BenchmarkMutex(b *testing.B) {
	for i:= 0; i < b.N; i++{
		x := 0
		var mutex sync.Mutex

		for i:= 0; i < 1000; i++ {
			go func (){
				for i:= 0; i < 100; i++{
					mutex.Lock()
					x++
					mutex.Unlock()
				}
			}()
		}
	}
}
func BenchmarkNoMutex(b *testing.B) {
	for i:= 0; i < b.N; i++{
		x := 0

		for i:= 0; i < 1000; i++ {
			go func (){
				for i:= 0; i < 100; i++{
					x++
				}
			}()
		}
	}
}