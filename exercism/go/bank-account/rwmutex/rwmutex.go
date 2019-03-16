package rwmutex

import "sync"

type Account struct {
	balance int64
	closed  bool
	sync.RWMutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, closed: false}
}

/*

	without synchronization i.e. without protection for the critical section,
	if there are 10 +ve writers and 10 -ve writers, what's the minimum value the balance could reach?
	-9? since for -ve writers to enter the update statement a.balance + amount >=0 needs to be true.
	therefore at least one +ve writer should succeed.
	once that's done, all +ve/-ve writers could be inside the if check updating the balance and
	ending up over writing each other's values.
	so after an initial +1, all 10 -ve writers one by one overwrite each positive value of the 9 +ve writers.
	write a JMH style program to test this out?
	"accepted", "interesting", "unexpected" ... ?
*/

// conditional write -- can we do something better?
// we require mutual exclusion only in write case

// Deposit a positive amount into balance
// Withdraws the amount if negative
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return a.balance, false
	}

	if amount >= 0 || (a.balance+amount) >= 0 {
		a.balance += amount // modification
		return a.balance, true
	}

	return a.balance, false
}

// only reads here

func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()
	if a.closed {
		return a.balance, false
	}
	return a.balance, true
}

// conditional write -- can we do something better?
// we require mutual exclusion only in write case

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return a.balance, false
	}
	a.closed = true // modification
	payout := a.balance
	a.balance = 0 // modification
	return payout, true
}
