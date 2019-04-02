package paxos

type messageType int

const (
	Prepare messageType = iota
	Promise
	Accept
	Accepted
)

type message struct {
	t    messageType
	from int
	pn   int
	pv   int
}

func NewPrepareMessage(from, pn int) message {
	return message{t: Prepare, from: from, pn: pn}
}

func NewPromiseMessage(from, apn, apv int) message {
	return message{t: Promise, from: from, pn: apn, pv: apv}
}

func NewAcceptMessage(from, pn, pv int) message {
	return message{t: Accept, from: from, pn: pn, pv: pv}
}

func NewAcceptedMessage(from, pn int) message {
	return message{t: Accepted, from: from, pn: pn}
}
