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
