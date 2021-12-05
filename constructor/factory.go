package constructor

import "fmt"

type Restaurant interface {
	GetFood()
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "h":
		return &HotPot{}
	case "n":
		return &Noodles{}
	}
	return nil
}

type HotPot struct{}

func (h *HotPot) GetFood() {
	fmt.Println("get hotpot")
}

type Noodles struct{}

func (n *Noodles) GetFood() {
	fmt.Println("get noodles")
}
