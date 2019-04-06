package paxos

func ReturnsOne() int {
	return 1
}

type network struct {
	proposers []*proposer
	acceptors []*acceptor
	learners  []*learner
}

func NewNetwork(nProposers, nAcceptors, nLearners int, vs []int) *network {
	cProposers := makeChannels(nProposers)
	cAcceptors := makeChannels(nAcceptors)
	cLearners := makeChannels(nLearners)

	n := new(network)
	n.proposers = make([]*proposer, nProposers)
	n.acceptors = make([]*acceptor, nAcceptors)
	n.learners = make([]*learner, nLearners)

	for i := range n.proposers {
		n.proposers[i] = NewProposer(i, vs[i], cProposers[i], cAcceptors, cLearners)
	}

	for i := range n.acceptors {
		n.acceptors[i] = NewAcceptor(i, cAcceptors[i], cProposers)
	}

	for i := range n.learners {
		n.learners[i] = NewLearner(i, cLearners[i])
	}

	return n
}

func makeChannels(n int) []chan message {
	chans := make([]chan message, n)
	for i := range chans {
		chans[i] = make(chan message, 1024)
	}
	return chans
}

func (n *network) start() {
	for _, l := range n.learners {
		go l.run()
	}

	for _, a := range n.acceptors {
		go a.run()
	}

	for _, p := range n.proposers {
		go p.run()
	}
}
