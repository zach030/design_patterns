package constructor

import "fmt"

type Lunch interface {
	Cook()
}

type Rice struct {}

func (r *Rice)Cook(){
	fmt.Println("cooking rice")
}

type Tomato struct {}

func (t *Tomato)Cook(){
	fmt.Println("cooking tomato")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {}

func (s *SimpleLunchFactory) CreateFood() Lunch{
	return &Rice{}
}

func (s *SimpleLunchFactory)CreateVegetable()Lunch{
	return &Tomato{}
}

func NewSimpleLunchFactory()SimpleLunchFactory{
	return SimpleLunchFactory{}
}