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
	n := NewNetwork(1, 1, 1, []int{1})
	go n.acceptors[0].run()
	go n.proposers[0].run()
	if n.learners[0].run() != 1 {
		t.Errorf("Did not receive the proposed value!")
	}
}
