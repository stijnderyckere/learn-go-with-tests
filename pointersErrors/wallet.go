package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

//Bitcoin is just an int
type Bitcoin int

//Wallet : has a balance
type Wallet struct {
	balance Bitcoin
}

//Deposit : adds an amount to wallet balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Withdraw : removes an amount to wallet balance
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

//Balance : returns balance amount of a wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
