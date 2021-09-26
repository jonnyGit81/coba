package user

import "sync"

type User struct {
	Bal   int
	setMu sync.Mutex
	upMu  sync.Mutex
}

func (u *User) SetBalance(bal int) {
	u.setMu.Lock()
	defer u.setMu.Unlock()

	u.Bal = bal
}

func (u *User) UpdateBal(bal int) {
	u.upMu.Lock()
	defer u.upMu.Unlock()
	prev := u.Bal
	total := prev + bal
	u.SetBalance(total)
}
