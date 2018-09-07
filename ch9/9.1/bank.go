// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdraw)

type withdraw struct {
	amount int
	result chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	wd := withdraw{amount, make(chan bool)}
	withdraws <- wd
	return <-wd.result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wd := <-withdraws:
			if balance >= wd.amount {
				balance -= wd.amount
				wd.result <- true
			} else {
				wd.result <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
