package paxos

type proposer struct {
	id        int
	pv        int
	pn        int
	receives  chan message
	acceptors []chan message
	learners  []chan message
}

func NewProposer(id int, acceptors, learners []chan message) *proposer {
	p := new(proposer)
	p.id = id
	p.pv = 0
	p.pn = 0
	p.acceptors = acceptors
	p.learners = learners
	return p
}

func (p *proposer) run() {
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

		responded = make(map[int]bool)
		max = p.pn
		for len(responded) < len(p.acceptors)/2+1 {
			msg := <-p.receives
			switch msg.t {
			case Accepted:
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

	p.decide()
}

func (p *proposer) prepare() {
	p.pn++
	msg := NewPrepareMessage(p.id, p.pn)
	broadcast(p.acceptors, msg)
}

func (p *proposer) accept() {
	msg := NewAcceptMessage(p.id, p.pn, p.pv)
	broadcast(p.acceptors, msg)
}

func (p *proposer) decide() {
	msg := NewDecidedMessage(p.id, p.pv)
	broadcast(p.learners, msg)
}

func broadcast(peers []chan message, msg message) {
	for _, peer := range peers {
		peer <- msg
	}
}
