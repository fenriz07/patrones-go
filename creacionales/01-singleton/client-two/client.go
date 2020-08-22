package client_two

import "github.com/fenriz07/patrones-go/creacionales/01-singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
