package decorator

type TomatoPizza struct {}

func (t *TomatoPizza) DoPizza() string {
	return "I am a Tomato Pizza"
}
