package main

import (
	"design-pattern/abstractfactory"
	"design-pattern/builder"
	"design-pattern/prototype"
	"design-pattern/singleton"
	"fmt"
	"time"
)

func main() {

	// Prototype Pattern
	file1 := &prototype.File{Name: "File 1"}
	file2 := &prototype.File{Name: "File 2"}
	file3 := &prototype.File{Name: "File 3"}

	folder1 := &prototype.Folder{
		Children: []prototype.INode{file1},
		Name: "folder 1",
	}
	folder2 := &prototype.Folder{
		Children: []prototype.INode{folder1, file2, file3},
		Name: "folder 2",
	}

	fmt.Println("Printing for Folder 2")
	folder2.Print("   ")
	cloneFolder := folder2.Clone()
	fmt.Println("Printing for clone Folder 2")
	cloneFolder.Print("   ")

	fmt.Println("-------------------------------------")

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