package main

import (
	"fmt"
	"math/rand"
	"sync"
)
type BankAccount struct{
	name string
	Balance int
}
var mutex=&sync.Mutex{}
func withdrawl(account *BankAccount,amount int,done chan bool){
	mutex.Lock()
	if account.Balance >=amount {
		account.Balance -= amount
		fmt.Println(amount, " amount is credited to your account", "Balance amount is ", account.Balance)
	} else {
		fmt.Println("Insufficient Balance")
	}
	mutex.Unlock()
	done<-true
}
func deposit(account *BankAccount,amount int,done chan bool){
	mutex.Lock()
	account.Balance+=amount
	fmt.Println(amount," amount is credited to your account","Balance amount is ",account.Balance)
	mutex.Unlock()
	done<-true
}
func GetBalance (account *BankAccount){
	mutex.Lock()
	fmt.Println("account balance is ", account.Balance)
	mutex.Unlock()

}
func main(){
	fmt.Println("Hello")
	account:=&BankAccount{"Suresh",500}
	const total=10
	done:=make(chan bool,total)
	done2:=make(chan bool,total)
	for trans:=0;trans<total;trans++{
		amount:=rand.Intn(1000)
		go deposit(account,amount,done)
		go withdrawl(account,amount,done2)
	}
	for i:=0;i<total;i++{
		<-done
		<-done2
	}
	GetBalance(account)
}
