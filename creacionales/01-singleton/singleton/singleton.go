package singleton

import "sync"

type person struct {
	name string
	age  int
	mux  sync.Mutex
}

var p *person
var once sync.Once

/*GetInstance Export Singleton*/
func GetInstance() *person {

	once.Do(func() {
		p = &person{}
	})

	return p
}

func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age

}
