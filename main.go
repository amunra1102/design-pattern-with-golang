package main

import (
	"design-pattern/abstractfactory"
	"design-pattern/adapter"
	"design-pattern/bridge"
	"design-pattern/builder"
	"design-pattern/chainofresponsibility"
	"design-pattern/command"
	"design-pattern/composite"
	"design-pattern/decorator"
	"design-pattern/facade"
	"design-pattern/flyweight"
	"design-pattern/iterator"
	"design-pattern/mediator"
	"design-pattern/memento"
	"design-pattern/observer"
	"design-pattern/prototype"
	"design-pattern/proxy"
	"design-pattern/singleton"
	"design-pattern/state"
	"design-pattern/strategy"
	"design-pattern/templatemethod"
	"design-pattern/visitor"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	// Decorator pattern
	tomatoPizza := &decorator.TomatoPizza{}
	chickenPizza := &decorator.ChickenPizza{}

	fmt.Println(tomatoPizza.DoPizza())
	fmt.Println(chickenPizza.DoPizza())

	pepperDecorator := decorator.NewPepperDecorator(chickenPizza)
	fmt.Println(pepperDecorator.DoPizza())
	pepperDecorator.SetPizza(tomatoPizza)
	fmt.Println(pepperDecorator.DoPizza())

	cheeseDecorator := decorator.NewCheeseDecorator(chickenPizza)
	fmt.Println(cheeseDecorator.DoPizza())
	cheeseDecorator.SetPizza(tomatoPizza)
	fmt.Println(cheeseDecorator.DoPizza())

	fmt.Println("-------------------------------------")

	// Proxy pattern
	nginxServer := proxy.NewNginxServer()

	appStatusURL := "app/status"
	createUserUrl := "create/user"

	httpCode, message := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s\nHttpCode: %d\nMessage: %s\n\n", appStatusURL, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s\nHttpCode: %d\nMessage: %s\n\n", appStatusURL, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s\nHttpCode: %d\nMessage: %s\n\n", appStatusURL, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(createUserUrl, "POST")
	fmt.Printf("URL: %s\nHttpCode: %d\nMessage: %s\n\n", createUserUrl, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(createUserUrl, "POST")
	fmt.Printf("URL: %s\nHttpCode: %d\nMessage: %s\n\n", createUserUrl, httpCode, message)

	httpCode, message = nginxServer.HandleRequest("createUserUrl", "GET")
	fmt.Printf("URL: %s\nHttpCode: %d\nMessage: %s\n", createUserUrl, httpCode, message)

	fmt.Println("-------------------------------------")

	// Fly weight pattern
	factory := flyweight.NewFactory()

	for i := 0; i < 15; i++ {
		starContext := flyweight.NewContext(strconv.Itoa(i), 2)
		soldier := factory.CreateSoldier("Foot Man")
		soldier.Promote(starContext)
	}

	for i := 0; i < 30; i++ {
		starContext := flyweight.NewContext(strconv.Itoa(i), 2)
		soldier := factory.CreateSoldier("Sea Man")
		soldier.Promote(starContext)
	}

	fmt.Printf("Number of storage soldier map: %d\n", factory.GetSize())

	fmt.Println("-------------------------------------")

	// Facade pattern
	walletFacade := facade.NewWalletFacade("abc", 1234)
	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)

	fmt.Println()
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	fmt.Println("-------------------------------------")

	// Composite pattern
	file1Com := &composite.File{Name: "File1"}
	file2Com := &composite.File{Name: "File2"}
	file3Com := &composite.File{Name: "File3"}

	folder1Com := &composite.Folder{Name: "Folder 1"}
	folder1Com.Add(file1Com)

	folder2Com := &composite.Folder{Name: "Folder 2"}
	folder2Com.Add(file2Com)
	folder2Com.Add(file3Com)
	folder2Com.Add(folder1Com)

	folder2Com.Search("rose")

	fmt.Println("-------------------------------------")

	// Bridge pattern
	hpPrinter := &bridge.HP{}
	epsonPrinter := &bridge.Epson{}

	macComputer := &bridge.Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	windowComputer := &bridge.Window{}
	windowComputer.SetPrint(hpPrinter)
	windowComputer.Print()

	windowComputer.SetPrint(epsonPrinter)
	windowComputer.Print()

	fmt.Println("-------------------------------------")

	// Adapter pattern
	client := &adapter.Client{}

	window := &adapter.Window{}
	client.InsertSquareUSBInComputer(window)

	macAdapter := &adapter.MacAdapter{M: adapter.Mac{}}
	client.InsertSquareUSBInComputer(macAdapter)

	fmt.Println("-------------------------------------")

	// Visitor pattern
	square := &visitor.Square{Side: 2}
	circle := &visitor.Circle{Radius: 3}
	rectangle := &visitor.Rectangle{A: 4, B: 5}

	areaCal := &visitor.AreaCalculator{}
	square.Accept(areaCal)
	circle.Accept(areaCal)
	rectangle.Accept(areaCal)

	perimeterCal := &visitor.PerimeterCalculator{}
	square.Accept(perimeterCal)
	circle.Accept(perimeterCal)
	rectangle.Accept(perimeterCal)

	fmt.Println("-------------------------------------")

	// Template method pattern
	smsOTP := &templatemethod.SMS{}
	o := &templatemethod.OTP{ObjectOTP: smsOTP}
	o.GenAndSendOTP(4)

	emailOTP := &templatemethod.Email{}
	o = &templatemethod.OTP{ObjectOTP: emailOTP}
	o.GenAndSendOTP(4)

	fmt.Println("-------------------------------------")

	// Strategy pattern
	lfu := &strategy.LFU{}
	cache := strategy.InitCache(lfu)
	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")
	cache.Add("c2", "3")

	lru := &strategy.LRU{}
	cache.SetEvictionAlgo(lru)
	cache.Add("d", "4")

	fifo := &strategy.FIFO{}
	cache.SetEvictionAlgo(fifo)
	cache.Add("e", "5")

	fmt.Println("-------------------------------------")

	// State pattern
	machine := state.NewMachine()
	machine.On()
	machine.Off()
	machine.Off()

	fmt.Println("-------------------------------------")

	// Observer pattern
	shirtItem := observer.NewItem("Nike shirt")
	observerFirst := &observer.Customer{ID: "first@gmail.com"}
	observerSecond := &observer.Customer{ID: "second@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()

	shirtItem.DeRegister(observerSecond)

	shirtItem.UpdateAvailability()

	fmt.Println("-------------------------------------")

	// Memento Pattern
	careTaker := &memento.CareTaker{
		MementoArray: make([]*memento.Memento, 0),
	}

	originator := &memento.Originator{}
	originator.SetState("A")

	fmt.Printf("Originnator current state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originnator current state: %s\n", originator.GetState())

	careTaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originnator current state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(careTaker.GetMemento(1))
	fmt.Printf("Originnator current state: %s\n", originator.GetState())

	originator.RestoreMemento(careTaker.GetMemento(0))
	fmt.Printf("Originnator current state: %s\n", originator.GetState())

	fmt.Println("-------------------------------------")

	// Mediator pattern
	stationManager := mediator.NewStationManager()
	passengerTrain := &mediator.PassengerTrain{
		Mediator: stationManager,
	}

	goodsTrain := &mediator.GoodsTrain{
		Mediator: stationManager,
	}

	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()

	passengerTrain.Departure()

	fmt.Println("-------------------------------------")

	// Iterator pattern
	user1 := &iterator.User{
		Name: "Test 1",
		Age: 18,
	}

	user2 := &iterator.User{
		Name: "Test 2",
		Age: 28,
	}

	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}

	iter := userCollection.CreateIterator()

	for iter.HasNext() {
		user := iter.GetNext()
		fmt.Printf("User is %+v\n", user)
	}

	fmt.Println("-------------------------------------")

	// Command pattern
	tv := &command.Tivi{}
	onCommand := &command.OnCommand{
		Device: tv,
	}

	offCommand := &command.OffCommand{
		Device: tv,
	}

	onButton := &command.Button{
		Command: onCommand,
	}

	onButton.Press()

	offButton := &command.Button{Command: offCommand}
	offButton.Press()

	fmt.Println("-------------------------------------")

	// Chain of responsibility
	cashier := &chainofresponsibility.Cashier{}

	medical := &chainofresponsibility.Medicine{}
	medical.SetNext(cashier)

	doctor := &chainofresponsibility.Doctor{}
	doctor.SetNext(medical)

	reception := &chainofresponsibility.Reception{}
	reception.SetNext(doctor)

	patient := &chainofresponsibility.Patient{Name: "Patient test"}
	reception.Execute(patient)

	fmt.Println("-------------------------------------")

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