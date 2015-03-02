# -*- coding: utf-8 -*-


class Building(object):

    def __init__(self, action, work):
        self.do_action = action()
        self.do_work = work()

    def action(self):
        self.do_action.action()

    def work(self):
        self.do_work.work()

    def livehuman(self):
        print "people live here!!"


class Officetel(Building):

    def __init__(self):
        Building.__init__(self, Sleeping, Nowork)
        print "it's just normal house"


class Policeoffice(Building):

    def __init__(self):
        Building.__init__(self, Catchcriminal, Hardwork)
        print "it's policeoffice!!"


class Firestation(Building):

    def __init__(self):
        Building.__init__(self, Stopfire, Easywork)
        print "it's fire station!!"


class Action():

    def action(self):
        pass


class Stopfire(Action):

    def action(self):
        print "we get fire off !!"


class Catchcriminal(Action):

    def action(self):
        print "NYPD!!"


class Sleeping(Action):

    def action(self):
        print "Zzzz....."


class Work():

    def work(self):
        pass


class Nowork(Work):

    def work(self):
        print "Noworking"


class Hardwork(Work):

    def work(self):
        print "so hard!!"


class Easywork(Work):

    def work(self):
        print "easy work :)"

if __name__ == "__main__":

    officetel = Officetel()
    officetel.action()
    officetel.work()
    officetel.livehuman()

    print "\n**********************"
    policeoffice = Policeoffice()
    policeoffice.action()
    policeoffice.work()
    policeoffice.livehuman()
    policeoffice.do_action = Sleeping()
    policeoffice.action()

    print "\n**********************"

    firestation = Firestation()
    firestation.action()
    firestation.work()
    firestation.livehuman()
    firestation.do_action = Sleeping()
    firestation.action()
