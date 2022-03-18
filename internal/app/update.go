package app

import "dishes-admin-mod/internal/app/update"

func InitUpdate() (*update.Update, func(), error) {
	u := update.Update{}
	u.FailedChan = make(map[uint64]chan struct{})
	u.SuccessChan = make(map[uint64]chan struct{})
	return &u, func() {}, nil
}
