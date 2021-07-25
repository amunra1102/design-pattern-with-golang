package adapter

import "fmt"

type Mac struct{}

func (m *Mac) insertInCirclePort() {
	fmt.Println("Insert circle port into Mac machine")
}