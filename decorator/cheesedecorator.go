package decorator

type CheeseDecorator struct {
	pizza IPizza
}

func NewCheeseDecorator(pizza IPizza) *CheeseDecorator {
	return &CheeseDecorator{pizza: pizza}
}

func (c *CheeseDecorator) DoPizza() string {
	pizzaType := c.pizza.DoPizza()
	return pizzaType + c.addFlavour()
}

func (c *CheeseDecorator) SetPizza(pizza IPizza) {
	c.pizza = pizza
}

func (c *CheeseDecorator) addFlavour() string {
	return " Cheese"
}
