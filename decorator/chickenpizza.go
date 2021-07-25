package decorator

type ChickenPizza struct {}

func (c *ChickenPizza) DoPizza() string {
	return "I am a Chicken Pizza"
}