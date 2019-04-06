package paxos

import "testing"

func TestPaxos(t *testing.T) {
	if ReturnsOne() != 1 {
		t.Errorf("You really messed up here!")
	}
}

func TestBroadcast(t *testing.T) {
	acceptors := []chan message{make(chan message, 100), make(chan message, 100)}
	msg := message{}
	broadcast(acceptors, msg)
	for _, acceptor := range acceptors {
		if <-acceptor != msg {
			t.Errorf("Received message was different to sent message!")
		}
	}
}

func TestSingleProposer(t *testing.T) {
	n := NewNetwork(1, 1, 2, []int{1})
	go n.acceptors[0].run()
	go n.proposers[0].run()
	if n.learners[0].run() != n.learners[1].run() {
		t.Errorf("Did not receive the proposed value!")
	}
}

func TestTwoProposersSameValue(t *testing.T) {
	n := NewNetwork(2, 1, 2, []int{2, 2})
	go n.acceptors[0].run()
	go n.proposers[0].run()
	go n.proposers[1].run()
	if n.learners[0].run() != n.learners[1].run() {
		t.Errorf("Did not receive the proposed value!")
	}
}

func TestTwoProposersDifferentValue(t *testing.T) {
	n := NewNetwork(2, 1, 2, []int{1, 2})
	go n.acceptors[0].run()
	go n.proposers[0].run()
	go n.proposers[1].run()
	if n.learners[0].run() != n.learners[1].run() {
		t.Errorf("Did not receive the proposed value!")
	}
}

func TestManyProposersDifferentValues(t *testing.T) {
	n := NewNetwork(5, 1, 2, []int{1, 2, 3, 4, 5})
	go n.acceptors[0].run()
	for _, p := range n.proposers {
		go p.run()
	}
	if n.learners[0].run() != n.learners[1].run() {
		t.Errorf("Did not receive the proposed value!")
	}
}

func TestTwoAcceptors(t *testing.T) {
	n := NewNetwork(1, 2, 2, []int{3})
	go n.acceptors[0].run()
	go n.acceptors[1].run()
	go n.proposers[0].run()
	if n.learners[0].run() != n.learners[1].run() {
		t.Errorf("Did not receive the proposed value!")
	}
}

func TestManyProposersManyAcceptorsSameValue(t *testing.T) {
	n := NewNetwork(5, 5, 2, []int{1, 1, 1, 1, 1})
	for _, a := range n.acceptors {
		go a.run()
	}
	for _, p := range n.proposers {
		go p.run()
	}
	if n.learners[0].run() != n.learners[1].run() {
		t.Errorf("Did not receive the proposed value!")
	}
}

func TestManyProposersManyAcceptorsDifferentValues(t *testing.T) {
	n := NewNetwork(5, 5, 2, []int{1, 2, 3, 4, 5})
	for _, a := range n.acceptors {
		go a.run()
	}
	for _, p := range n.proposers {
		go p.run()
	}
	if n.learners[0].run() != n.learners[1].run() {
		t.Errorf("Did not receive the proposed value!")
	}
}

func TestManyProposersManyAcceptorsSemiSameValues(t *testing.T) {
	n := NewNetwork(5, 5, 2, []int{1, 1, 1, 2, 2})
	for _, a := range n.acceptors {
		go a.run()
	}
	for _, p := range n.proposers {
		go p.run()
	}
	if n.learners[0].run() != n.learners[1].run() {
		t.Errorf("Did not receive the proposed value!")
	}
}
