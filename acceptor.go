package paxos

import "fmt"

type acceptor struct {
	id        int
	apn       int
	apv       int
	maxpn     int
	receives  chan message
	proposers []chan message
}

// NewAcceptor makes a new acceptor component with identifier id, receival
// channel receives, and channels to proposers through proposers.
func NewAcceptor(id int, receives chan message, proposers []chan message) *acceptor {
	a := new(acceptor)
	a.id = id
	a.apn = 0
	a.apv = 0
	a.maxpn = 0
	a.receives = receives
	a.proposers = proposers
	return a
}

// Run starts the acceptor's Paxos algorithm.:w
func (a *acceptor) Run() {
	fmt.Printf("Acceptor %v: started\n", a.id)
	for {
		msg := <-a.receives
		switch msg.t {
		// PHASE 1: Prepare-Promise
		case Prepare:
			if msg.pn > a.maxpn {
				a.maxpn = msg.pn
			}
			a.proposers[msg.from] <- NewPromiseMessage(a.id, a.apn, a.apv)
			fmt.Printf("Acceptor %v: sending Promise\n", a.id)
		// PHASE 2: Accept-Accepted
		case Accept:
			if msg.pn >= a.maxpn {
				a.maxpn = msg.pn
				a.apn = msg.pn
				a.apv = msg.pv
			}
			a.proposers[msg.from] <- NewAcceptedMessage(a.id, a.maxpn)
			fmt.Printf("Acceptor %v: sending Accepted\n", a.id)
		default:
		}
	}

}
