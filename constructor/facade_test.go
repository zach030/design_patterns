package constructor

import "testing"

func TestCarFacade_CreateCompleteCar(t *testing.T) {
	cf := NewCarFacade()
	cf.CreateCompleteCar()
}