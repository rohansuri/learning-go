package account

// to submit the final solution to exercism

type Account interface {
	Deposit(amount int64) (int64, bool)
	Balance() (int64, bool)
	Close() (int64, bool)
}

func Open(initialDeposit int64, open func(int64) *Account) *Account {
	return open(initialDeposit)
}
