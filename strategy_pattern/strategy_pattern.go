package main

import (
	"fmt"
)

type Building struct {
	do_action Actioner
	do_work   Worker
	live      Liver
}

func (this *Building) SetAction(a Actioner) {
	this.do_action = a
}

func (this *Building) SetWork(w Worker) {
	this.do_work = w
}

type Liver interface {
	live()
}

func live() {
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

func (b *Building) Officetel() *Building {
	return &Building{new(Sleeping), new(Nowork), b.live}
}

func (b *Building) Policeoffice() *Building {
	return &Building{new(Catchcriminal), new(Hardwork), b.live}
}

func (b *Building) Firestation() *Building {
	return &Building{new(Stopfire), new(Easywork), b.live}
}

func main() {
	var b Building

	officetel := b.Officetel()
	officetel.do_action.action()
	officetel.do_work.work()
	officetel.SetAction(new(Sleeping))
	officetel.do_action.action()

	fmt.Println("***********************")
	policeoffice := b.Policeoffice()
	policeoffice.do_action.action()
	policeoffice.do_work.work()
	policeoffice.SetAction(new(Sleeping))
	policeoffice.do_action.action()

	fmt.Println("***********************")
	firestation := b.Firestation()
	firestation.do_action.action()
	firestation.do_work.work()
	firestation.SetWork(new(Nowork))
	firestation.do_work.work()
}
