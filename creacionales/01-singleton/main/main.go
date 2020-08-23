package main

import (
	"fmt"
	"sync"

	client_one "github.com/fenriz07/patrones-go/creacionales/01-singleton/client-one"
	client_two "github.com/fenriz07/patrones-go/creacionales/01-singleton/client-two"
	"github.com/fenriz07/patrones-go/creacionales/01-singleton/singleton"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(200)

	for i := 0; i < 100; i++ {

		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()

		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()

	}

	wg.Wait()

	p := singleton.GetInstance()

	age := p.GetAge()

	fmt.Println(age)
}
