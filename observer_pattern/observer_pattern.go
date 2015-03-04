package main

import (
	"fmt"
)

type Orderfooder interface {
	registerObserver()
	removeObserver()
	notifyObservers()
}

type Subject struct {
	name       string
	price      int64
	address    string
	cardnumber string
	orderfood  Orderfooder
}

func (o *Orderfooder) registerObserver() {

}

func (o *Orderfooder) removeObserver() {

}

func (o *Orderfooder) notifyObservers() {

}

type Observer struct {
	Subject
}

type OrderAccept struct{}

func (o *OrderAccept) update(name string, price int64, address string, cardnumber string) {
	o.name = name
	o.price = price
	o.address = address
	o.cardnumber = cardnumber
	o.OrderAcceptDisplay.display()
}

func (o *OrderAccept) display() {
	fmt.Printf("name : %s\nprice : %d\naddress : %s\ncardnumber :%s\n", o.name, o.price, o.address, o.cardnumber)
}

func (o *Observer) OrderAccept() *Observer {
	return &Observer{}
}

func main() {
	fmt.Println("observer")
}
