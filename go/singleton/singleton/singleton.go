package singleton

import "sync"

var (
	p *person
	once sync.Once
)

func GetInstance() *person {
	/* El metodo once.Do se va a ejecuta una UNICA VEZ EN TODO EL PAQUETE!!!
	aunque tengamos mas funciones que lo usen, este solo se usara una sola vez */
	once.Do(func() {
		p = &person {}
	})

	return p
}

type person struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	defer p.mux.Unlock()
	
	p.mux.Lock()
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	defer p.mux.Unlock()

	p.mux.Lock()
	return p.age
}
