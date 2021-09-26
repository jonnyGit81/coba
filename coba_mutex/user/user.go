package user

import "sync"

type User struct {
	Bal   int
	mutex sync.Mutex
}

func (u *User) SetBalance(val int) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.Bal += val
}

func (u *User) UpdateBalance(val int) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.SetBalance(val)
}
