package single_threaded

// Serves as a starter to try out various thread safety primitives around the basic non thread safe version.

type Account struct {
	balance int64
	closed  bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, closed: false}
}

// Deposit a positive amount into balance
// Withdraws the amount if negative
func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.closed {
		return a.balance, false
	}

	if amount >= 0 || (a.balance+amount) >= 0 {
		a.balance += amount // modification
		return a.balance, true
	}

	return a.balance, false
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return a.balance, false
	}
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	if a.closed {
		return a.balance, false
	}
	a.closed = true // modification
	payout := a.balance
	a.balance = 0 // modification
	return payout, true
}
