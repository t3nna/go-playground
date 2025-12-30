package main

import (
	"fmt"
	"time"
)

func main() {

	msgCh := make(chan int, 10)
	disconnectCh := make(chan struct{})
	go func() {
		for {

			select {
			case val := <-msgCh:
				fmt.Println(val)
			case <-disconnectCh:
				for {
					select {
					case val := <-msgCh:

						fmt.Println(val)
					default:
						return
					}
				}

			}
		}
	}()

	for i := 0; i < 10; i++ {
		msgCh <- i
	}
	disconnectCh <- struct{}{}

	time.Sleep(1 * time.Second)
}
