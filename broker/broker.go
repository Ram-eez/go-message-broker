package broker

// Message that will be passed in the Subscriber Channel
type Message struct {
	Topic   string
	Payload interface{}
}

// This represents the cilent(us only, no multiple clients). This client will sub to various topics.
type Subscriber struct {
	Channel     chan interface{}
	Unsubscribe chan bool
}
