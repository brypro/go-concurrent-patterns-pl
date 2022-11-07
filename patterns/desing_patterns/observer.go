package main

import "fmt"

/*
Observer is a behavioral design pattern
that lets you define a subscription mechanism
to notify multiple objects about any events
that happen to the object they’re observing.
*/

func main() {
	nvidiaItem := NewItem("Nvidia RTX 3080")
	firstObserver := &EmailClient{id: "first@code.cl"}
	secondObserver := &EmailClient{id: "second@code.cl"}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvaliable()

}

type Observer interface {
	getId() string
	updateValue(string)
}
type Topic interface {
	register(observer Observer) // register observer
	broadcast()                 // notifyAllObservers()
}

type Item struct {
	observers []Observer
	name      string
	avaliable bool
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}
func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func NewItem(name string) *Item {
	return &Item{
		name:      name,
		avaliable: false,
	}
}
func (i *Item) UpdateAvaliable() {
	fmt.Println("Item", i.name, "is avaliable")
	i.avaliable = true
	i.broadcast()
}

type EmailClient struct {
	id string
}

func (e *EmailClient) getId() string {
	return e.id
}
func (e *EmailClient) updateValue(name string) {
	fmt.Println("EmailClient", e.id, "is notified about", name)
}