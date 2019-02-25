package account

type Account struct {
	balance int64
	closed bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, false}
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.closed {
		return a.balance, false
	}

	if amount >= 0 || (a.balance + amount) >= 0{
		a.balance += amount
		return a.balance, true
	}

	return a.balance, false
}

func (a *Account) Balance() (int64, bool){
	if a.closed {
		return a.balance, false
	}
	return a.balance, true
}


func (a *Account) Close() (int64, bool){
	if a.closed {
		return a.balance, false
	}
	a.closed = true
	payout := a.balance
	a.balance = 0
	return payout, true
}