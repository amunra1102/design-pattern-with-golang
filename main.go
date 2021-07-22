package main

import (
	"design-pattern/abstractfactory"
	"design-pattern/builder"
	"design-pattern/singleton"
	"fmt"
	"time"
)

func main() {

	// Abstract Factory
	adidasFactory := abstractfactory.GetSportFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasShort := adidasFactory.MakeShort()
	printShoeDetail(adidasShoe)
	printShortDetail(adidasShort)

	nikeFactory := abstractfactory.GetSportFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()
	printShoeDetail(nikeShoe)
	printShortDetail(nikeShort)

	fmt.Println("-------------------------------------")

	// Builder Pattern
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.GetFloor())

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Printf("Igloor House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Igloor House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Igloor House Num Floor: %d\n", iglooHouse.GetFloor())

	fmt.Println("-------------------------------------")

	// Singleton Pattern
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("%p\n", singleton.GetInstance())
		}()
	}
	time.Sleep(time.Second * 20)
}

func printShoeDetail(s abstractfactory.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShortDetail(s abstractfactory.IShort) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}