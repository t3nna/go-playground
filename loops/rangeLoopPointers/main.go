package main

import (
	"fmt"
	"log"
)

type Customer struct {
	ID      string
	Balance float64
}
type Store struct {
	m map[string]*Customer
}

func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		s.m[customer.ID] = &customer
	}
}

func main() {
	s := Store{map[string]*Customer{}}

	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})

	for _, val := range s.m {
		log.Println(val)
	}
}
