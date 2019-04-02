package paxos

type proposer struct {
	id        int
	pv        int
	pn        int
	receives  chan message
	acceptors []chan Message
}

func NewProposer(id int, acceptors []chan message) *proposer {
	p := new(proposer)
	p.id = id
	p.pv = 0
	p.pn = 0
	p.acceptors = acceptors
}

func (p *proposer) run() {
	p.prepare()
	responded := make(map[int]int)
}

func (p *proposer) prepare() {
	p.pn++
	msg := NewPrepareMessage(p.id, p.pn)
	p.broadcast(msg)
}

func (p *proposer) accept() {
	msg := NewAcceptMessage(p.id, p.pn, p.pv)
	p.broadcast(msg)
}

func (p *proposer) broadcast(msg message) {
	for _, acceptor := range p.acceptors {
		acceptor <- message
	}
}
