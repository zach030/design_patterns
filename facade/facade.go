package facade

import "fmt"

type CarModel struct{}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel() {
	fmt.Println("set Car Model")
}

type CarEngine struct{}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (c *CarEngine) SetEngine() {
	fmt.Println("set Car Engine")
}

type CarBody struct{}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetCarBody() {
	fmt.Println("set Car Body")
}

type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (c *CarFacade) CreateCompleteCar() {
	c.model.SetModel()
	c.body.SetCarBody()
	c.engine.SetEngine()
}
