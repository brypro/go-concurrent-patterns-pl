package main

import "fmt"

/*
factory is a creational design pattern
that provides an interface for creating objects in
a superclass, but allows subclasses to alter the type
of objects that will be created.

*/

type IProduct interface {
	getStock() int
	setStock(stock int)
	setName(name string)
	getName() string
}

type Computer struct {
	stock int
	name  string
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

// Laptop is a struct and implements IProduct interface methods from Computer struct
type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			stock: 10,
			name:  "Laptop Computer",
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			stock: 22,
			name:  "Desktop Computer",
		},
	}
}

func GetComputerFactory(computerType IProduct) (IProduct, error) {
	switch computerType.(type) {
	case *Laptop:
		return NewLaptop(), nil
	case *Desktop:
		return NewDesktop(), nil
	default:
		return nil, fmt.Errorf("Computer type not recognized")
	}
}

func PrintNameAndStock(product IProduct) {
	fmt.Printf("Product Name: %s\n", product.getName())
	fmt.Printf("Product Stock: %d\n", product.getStock())
}

func main() {
	lap, _ := GetComputerFactory(&Laptop{})
	desk, _ := GetComputerFactory(&Desktop{})
	PrintNameAndStock(lap)
	PrintNameAndStock(desk)
}
