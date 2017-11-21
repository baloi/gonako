package funding

type Fund struct {
	// balance is unexported (private because it is lowercase
	balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct without worrying
	// whether it's on the stack or heap: Go figures that out
	return &Fund {
		balance: initialBalance,
	}
}

// Methods start with *receiver, in this case a Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}