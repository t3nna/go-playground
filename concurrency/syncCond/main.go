package main

import (
	"fmt"
	"sync"
	"time"
)

type Donation struct {
	balance int
	cond    *sync.Cond
}

func main() {

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// listening on goroutines
	f := func(dGoal int) {
		donation.cond.L.Lock()
		for donation.balance < dGoal {
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached \n", donation.balance)
		donation.cond.L.Unlock()
	}
	go f(10)
	go f(15)

	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()

		if donation.balance == 16 {
			return
		}
	}

}
