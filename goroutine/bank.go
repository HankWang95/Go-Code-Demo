package goroutine

var deposits = make(chan int)
var withdraw = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}
func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraw:
			balance -= amount
		case balances <- balance:

		}
	}
}

//func init() {
//	go teller()
//}

func Withdraw(amount int) bool {
	b := Balance()
	if amount <= b {
		withdraw <- amount
		return true
	} else {
		return false
	}

}
