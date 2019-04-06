package paxos

import "fmt"

type learner struct {
	id       int
	receives chan message
}

// NewLearner makes a new learner component with identifier id and
// receival channel receives.
func NewLearner(id int, receives chan message) *learner {
	l := new(learner)
	l.id = id
	l.receives = receives
	return l
}

// Run starts the learner to listen to chosen values.
func (l *learner) Run() int {
	v := -1
	for v == -1 {
		msg := <-l.receives
		switch msg.t {
		case Chosen:
			v = msg.pv
		default:
		}
	}
	fmt.Printf("Learner %v: Chosen %v\n", l.id, v)
	return v
}
