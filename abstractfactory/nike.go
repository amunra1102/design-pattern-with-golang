package abstractfactory

type Nike struct {}

func (n *Nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "Nike",
			size: 15,
		},
	}
}

func (n *Nike) MakeShort() IShort {
	return &nikeShort{
		short: short{
			logo: "Nike",
			size: 16,
		},
	}
}