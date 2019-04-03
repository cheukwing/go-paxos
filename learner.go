package paxos

import "fmt"

type learner struct {
	receives chan message
}

func (l *learner) run() {
	v := -1
	for v == -1 {
		msg := <-l.receives
		switch msg.t {
		case Chosen:
			v = msg.pv
		default:
		}
	}
	fmt.Printf("Chosen %v\n", v)
}
