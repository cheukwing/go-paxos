package paxos

type acceptor struct {
	id        int
	apn       int
	apv       int
	maxpn     int
	receives  chan message
	proposers []chan message
}

func MakeAcceptor(id int, receives chan message, proposers []chan message) *acceptor {
	a := new(acceptor)
	a.id = id
	a.apn = 0
	a.apv = 0
	a.maxpn = 0
	a.receives = receives
	a.proposers = proposers
	return a
}

func (a *acceptor) run() {
	for {
		msg := <-a.receives
		switch msg.t {
		case Prepare:
			if msg.pn > a.maxpn {
				a.maxpn = msg.pn
			}
			a.proposers[msg.from] <- NewPromiseMessage(a.id, a.apn, a.apv)
		case Accept:
			if msg.pn >= a.maxpn {
				a.maxpn = msg.pn
				a.apn = msg.pn
				a.apv = msg.pv
			}
			a.proposers[msg.from] <- NewAcceptedMessage(a.id, a.maxpn)
		default:
		}
	}

}
