package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {

	now := time.Now()

	id := getUserByName("john")
	println(id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message, 2)
	quit := make(chan struct{})

	// as we have two go routines
	wg.Add(2)

	go getUserChats(id, ch, wg, quit)
	go getUserFriends(id, ch, wg)

loop:
	for {

		select {
		case val, _ := <-ch:
			log.Println(val)

		case <-quit:
			log.Println("quit")
			break loop

		}
	}

	wg.Wait()
	// is optional to close channel
	close(ch)

	// don't use range without closing chanel, deadlock will kill you
	//for msg := range ch {
	//	log.Println(msg)
	//}
	//for {
	//	val, ok := <-ch
	//	if !ok {
	//		break
	//	}
	//	log.Println(val)
	//}

	log.Println(time.Since(now))
}
func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)

	ch <- &Message{
		friends: []string{

			"bar",
			"kate",
			"toe",
			"den",
		},
	}
	defer wg.Done()
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup, quit chan<- struct{}) {
	time.Sleep(time.Second * 2)
	quit <- struct{}{}
	ch <- &Message{
		chats: []string{
			"john",
			"ivan",
			"toe",
			"den",
		},
	}
	defer wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s-2", name)
}
