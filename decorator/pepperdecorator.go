package decorator

type PepperDecorator struct {
	pizza IPizza
}

func NewPepperDecorator(pizza IPizza) *PepperDecorator {
	return &PepperDecorator{pizza: pizza}
}

func (p *PepperDecorator) DoPizza() string {
	pizzaType := p.pizza.DoPizza()
	return pizzaType + p.addFlavour()
}

func (p *PepperDecorator) SetPizza(pizza IPizza) {
	p.pizza = pizza
}

func (p *PepperDecorator) addFlavour() string {
	return " Pepper"
}