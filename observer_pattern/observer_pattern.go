package main

import (
	"fmt"
	"time"
)

var datetime = time.Now()

type Orderfooder interface {
	registerObserver(o *Observer)
	removeObserver(o *Observer)
	notifyObservers()
}

type Subject struct {
	observeruid int32
	location    string
	amountofcar int64
	accident    bool
	observers   []*Observer
}

func (s *Subject) registerObserver(o *Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) removeObserver(o *Observer) {
	for index, registobserver := range s.observers {
		if registobserver.observeruid == o.observeruid {
			s.observers = append(s.observers[:index], s.observers[index+1:]...)
		}
	}
}

func (s *Subject) notifyObservers() {
	if len(s.observers) == 0 {
		fmt.Println("please regist observer first")
	}
	for _, ob := range s.observers {
		ob.do_update.update(s.location, s.amountofcar, s.accident)
	}
}

func (s *Subject) setMeasurements(location string, amountofcar int64, accident bool) {
	fmt.Println("*********************************")
	s.location = location
	s.amountofcar = amountofcar
	s.accident = accident
	s.measurementsChanged()
}

func (s *Subject) measurementsChanged() {
	s.notifyObservers()
}

type Observer struct {
	observeruid int32
	location    string
	amountofcar int64
	accident    bool
	do_update   updater
	do_display  displayer
}

type updater interface {
	update(location string, amountofcar int64, accident bool)
}

type displayer interface {
	display()
}

type LiveTrafficNews struct {
	Observer
}

func (o *LiveTrafficNews) update(location string, amountofcar int64, accident bool) {
	o.location = location
	o.amountofcar = amountofcar
	o.accident = accident
	o.display()
}

func (o *LiveTrafficNews) display() {
	fmt.Printf("**CNN live traffic news**\nlocation : %s\namountofcar : %d\naccident : %v\ndatetime: %v\n\n", o.location, o.amountofcar, o.accident, datetime.Format(time.RFC3339))
}

type EmergencyService struct {
	Observer
}

func (o *EmergencyService) update(location string, amountofcar int64, accident bool) {
	if accident == true {
		o.location = location
		o.amountofcar = amountofcar
		o.display()
	}
}

func (o *EmergencyService) display() {
	var trafficjam string
	if o.amountofcar > 50000 {
		trafficjam = "terrible trafficjam"
	} else {
		trafficjam = "no trafficjam"
	}
	fmt.Printf("**EmergencyService**\naccident accept time : %v\nwe will there for you!\nnow %s on the road\nrun to %s right now!\n\n", datetime.Format(time.RFC3339), trafficjam, o.location)
}

type NavigationAlarm struct {
	Observer
}

func (o *NavigationAlarm) update(location string, amountofcar int64, accident bool) {
	o.location = location
	o.amountofcar = amountofcar
	o.accident = accident
	o.display()
}

func (o *NavigationAlarm) display() {
	if o.accident == true {
		fmt.Printf("**NavigationAlarm**\ndatetime : %v\nthere is car accident on %s,\n if our root drop by %s, we will change root.\n%d car on the road.\n\n", datetime.Format(time.RFC3339), o.location, o.location, o.amountofcar)
	} else {
		fmt.Printf("**NavigationAlarm**\ndatetime : %v\nsafe your driving!\n%s has %d car on the road. -by live traffic infomation\n\n", datetime.Format(time.RFC3339), o.location, o.amountofcar)
	}
}

func main() {
	var subject Subject
	livetrafficnews := &Observer{observeruid: 1, do_update: new(LiveTrafficNews), do_display: new(LiveTrafficNews)}
	emergencyservice := &Observer{observeruid: 2, do_update: new(EmergencyService), do_display: new(EmergencyService)}
	navigationalarm := &Observer{observeruid: 3, do_update: new(NavigationAlarm), do_display: new(NavigationAlarm)}
	subject.registerObserver(livetrafficnews)
	subject.registerObserver(emergencyservice)
	subject.setMeasurements("Chungdam", 72438, false)
	subject.registerObserver(navigationalarm)
	subject.setMeasurements("Bucheon", 32438, true)
	subject.removeObserver(livetrafficnews)
	subject.setMeasurements("Ilsan", 52438, false)
	subject.setMeasurements("Gangnam", 92438, true)
	subject.removeObserver(emergencyservice)
	subject.removeObserver(navigationalarm)
	subject.setMeasurements("incheon", 32438, true)
}
