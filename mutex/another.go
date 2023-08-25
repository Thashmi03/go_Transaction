package main

import (
	"fmt"
	"sync"
)

type BankAccount struct{
	balance float64
	mutex sync.Mutex
}

func (acc * BankAccount)Deposit(amount float64){
	acc.mutex.Lock()
	defer acc.mutex.Unlock()

	fmt.Printf("depositing %.2f\n",amount)
	acc.balance+=amount
	fmt.Printf("New Balance :%.2f\n",acc.balance)
}
func(acc *BankAccount)Withdraw(amount float64){
	acc.mutex.Lock()
	defer acc.mutex.Unlock()

	if acc.balance>=amount{
		fmt.Printf("withdrawing %.2f\n",amount)
		acc.balance-=amount
		fmt.Printf("new balance:%.2f\n",acc.balance)
	}else{
		fmt.Println("insufficient balance")
	}
}


func main(){
	b:=&BankAccount{balance:1000.0}
	var wg sync.WaitGroup

	for i:=0; i<5; i++{
		wg.Add(1)
		go func (index int) {
		defer wg.Done()
		if index%2 != 0{
		b.Deposit(2000.0)
		}else{
		b.Withdraw(35000.0)
		}
		}(i)
		}
		wg.Wait()
		fmt.Println(b.balance)
}