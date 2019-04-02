package paxos

import "testing"

func TestPaxos(t *testing.T) {
	if ReturnsOne() != 1 {
		t.Errorf("You really messed up here!")
	}
}
