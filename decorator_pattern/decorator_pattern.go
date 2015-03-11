package main

import (
	"fmt"
)

type PC interface {
	cost() int64
	getDescription() string
}

type GamingPc struct{}

func (g *GamingPc) cost() int64 { return 600000 }

func (p *GamingPc) getDescription() string {
	return "Gaming PC"
}

type OfficePc struct{}

func (g *OfficePc) cost() int64 { return 300000 }

func (p *OfficePc) getDescription() string {
	return "Office Pc"
}

type DeveloperPc struct{}

func (g *DeveloperPc) cost() int64 { return 1000000 }

func (p *DeveloperPc) getDescription() string {
	return "Developer Pc"
}

type SuperPc struct{}

func (g *SuperPc) cost() int64 { return 2000000 }

func (p *SuperPc) getDescription() string {
	return "SuperPc"
}

type PcDecorator struct {
	pc PC
}

func (c *PcDecorator) cost() int64 {
	return c.pc.cost()
}

func (c *PcDecorator) getDescription() string {
	return c.pc.getDescription()
}

type Ram2GB struct {
	*PcDecorator
}

func ram2GB(pc PC) *Ram2GB {
	ram := new(Ram2GB)
	ram.PcDecorator = &PcDecorator{pc}
	return ram
}

func (ram *Ram2GB) cost() int64 {
	return ram.pc.cost() + 20000
}

func (ram *Ram2GB) getDescription() string {
	return ram.pc.getDescription() + " + ram2GB"
}

type Ssd64GB struct {
	*PcDecorator
}

func ssd64GB(pc PC) *Ssd64GB {
	ssd := new(Ssd64GB)
	ssd.PcDecorator = &PcDecorator{pc}
	return ssd
}

func (ssd *Ssd64GB) cost() int64 {
	return ssd.pc.cost() + 70000
}

func (ssd *Ssd64GB) getDescription() string {
	return ssd.pc.getDescription() + " + ssd64GB"
}

type Cooler struct {
	*PcDecorator
}

func cooler(pc PC) *Cooler {
	cool := new(Cooler)
	cool.PcDecorator = &PcDecorator{pc}
	return cool
}

func (cool *Cooler) cost() int64 {
	return cool.pc.cost() + 10000
}

func (cool *Cooler) getDescription() string {
	return cool.pc.getDescription() + " + cooler"
}

type Usb30 struct {
	*PcDecorator
}

func usb30(pc PC) *Usb30 {
	usb := new(Usb30)
	usb.PcDecorator = &PcDecorator{pc}
	return usb
}

func (usb *Usb30) cost() int64 {
	return usb.pc.cost() + 10000
}

func (usb *Usb30) getDescription() string {
	return usb.pc.getDescription() + " + usb30"
}

func main() {
	fmt.Println("**********************")
	fmt.Println(" BUY PC here!         ")
	fmt.Println("**********************")
	fmt.Println("gaming Pc : 600000")
	fmt.Println("office Pc : 300000")
	fmt.Println("developer Pc : 1000000")
	fmt.Println("super Pc : 2000000")
	fmt.Println("**********************")
	fmt.Println("#optional item")
	fmt.Println("ram2GB : 20000")
	fmt.Println("ssd64GB : 20000")
	fmt.Println("cooler : 20000")
	fmt.Println("usb30 : 20000")
	fmt.Println("**********************")

	gamingpc := new(GamingPc)
	gamingpc1 := ram2GB(gamingpc)
	gamingpc1 = ram2GB(gamingpc1)
	gamingpc2 := ssd64GB(gamingpc1)
	gamingpc3 := cooler(gamingpc2)
	gamingpc4 := usb30(gamingpc3)

	fmt.Printf("1. %s = %d\n", gamingpc4.getDescription(), gamingpc4.cost())

	superpc := new(SuperPc)
	superpc1 := ssd64GB(superpc)
	fmt.Printf("2. %s = %d\n", superpc1.getDescription(), superpc1.cost())

}
