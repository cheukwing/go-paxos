package paxos

type proposer struct {
	id        int
	pv        int
	pn        int
	acceptors []chan Message
}

func NewProposer(id int, acceptors []chan message) *proposer {
	p := new(proposer)
	p.id = id
	p.pv = 0
	p.pn = 0
	p.acceptors = acceptors
}

func (p *proposer) prepare() {
	p.pn++

}
