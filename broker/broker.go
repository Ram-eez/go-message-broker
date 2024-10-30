package broker

import "sync"

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

// The broker struct keeps track of all the subscribers for each topic.
type Broker struct {
	subscribers map[string][]*Subscriber
	mutex       sync.Mutex
}

// This function creates a new broker by returning an empty map of topic: {subscribers}
func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]*Subscriber),
	}
}

// func (b *Broker) Subscribe(topic string) *Subscriber {

// }
