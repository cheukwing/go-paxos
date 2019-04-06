package paxos

import "fmt"

type proposer struct {
	id        int
	pv        int
	pn        int
	receives  chan message
	acceptors []chan message
	learners  []chan message
}

func NewProposer(id, v int, receives chan message, acceptors, learners []chan message) *proposer {
	p := new(proposer)
	p.id = id
	p.pv = v
	p.pn = 0
	p.receives = receives
	p.acceptors = acceptors
	p.learners = learners
	return p
}

func (p *proposer) run() {
	fmt.Printf("Proposer %v: started\n", p.id)
	decided := false
	for !decided {
		p.prepare()
		responded := make(map[int]bool)
		max := p.pn
		for len(responded) < len(p.acceptors)/2+1 {
			msg := <-p.receives
			switch msg.t {
			case Promise:
				responded[msg.from] = true
				if msg.pn > max {
					p.pv = msg.pv
					max = msg.pn
				}
			default:
			}
		}

		p.accept()
		responded = make(map[int]bool)
		max = p.pn
		for len(responded) < len(p.acceptors)/2+1 {
			msg := <-p.receives
			switch msg.t {
			case Accepted:
				responded[msg.from] = true
				if msg.pn > max {
					max = msg.pn
				}
			default:
			}
		}

		if p.pn == max {
			break
		}
		p.pn = max
	}

	p.chosen()
}

func (p *proposer) prepare() {
	p.pn++
	msg := NewPrepareMessage(p.id, p.pn)
	fmt.Printf("Proposer %v: sending Prepare\n", p.id)
	broadcast(p.acceptors, msg)
}

func (p *proposer) accept() {
	msg := NewAcceptMessage(p.id, p.pn, p.pv)
	fmt.Printf("Proposer %v: sending Accept\n", p.id)
	broadcast(p.acceptors, msg)
}

func (p *proposer) chosen() {
	msg := NewChosenMessage(p.id, p.pv)
	fmt.Printf("Proposer %v: sending Chosen\n", p.id)
	broadcast(p.learners, msg)
}

func broadcast(peers []chan message, msg message) {
	for _, peer := range peers {
		peer <- msg
	}
}
