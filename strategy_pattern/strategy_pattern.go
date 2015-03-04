package main

import (
	"fmt"
)

type Building struct {
	do_action Actioner
	do_work   Worker
}

func (b *Building) SetAction(a Actioner) {
	b.do_action = a
}

func (b *Building) SetWork(w Worker) {
	b.do_work = w
}

func (b *Building) live() {
	fmt.Println("im alive here!!")
}

type Actioner interface {
	action()
}

type Stopfire struct{}

func (this *Stopfire) action() {
	fmt.Println("we get fire off !!")
}

type Catchcriminal struct{}

func (this *Catchcriminal) action() {
	fmt.Println("NYPD!!")
}

type Sleeping struct{}

func (this *Sleeping) action() {
	fmt.Println("Zzzz.......")
}

type Worker interface {
	work()
}

type Nowork struct{}

func (this *Nowork) work() {
	fmt.Println("Noworking")
}

type Hardwork struct{}

func (this *Hardwork) work() {
	fmt.Println("so hard!!")
}

type Easywork struct{}

func (this *Easywork) work() {
	fmt.Println("easy work :) ")
}

func Officetel() *Building {
	return &Building{new(Sleeping), new(Nowork)}
}

func Policeoffice() *Building {
	return &Building{new(Catchcriminal), new(Hardwork)}
}

func Firestation() *Building {
	return &Building{new(Stopfire), new(Easywork)}
}

func main() {

	officetel := Officetel()
	officetel.do_action.action()
	officetel.do_work.work()
	officetel.SetAction(new(Sleeping))
	officetel.do_action.action()
	officetel.live()

	fmt.Println("***********************")
	policeoffice := Policeoffice()
	policeoffice.do_action.action()
	policeoffice.do_work.work()
	policeoffice.SetAction(new(Sleeping))
	policeoffice.do_action.action()
	policeoffice.live()

	fmt.Println("***********************")
	firestation := Firestation()
	firestation.do_action.action()
	firestation.do_work.work()
	firestation.SetWork(new(Nowork))
	firestation.do_work.work()
	firestation.live()
}
