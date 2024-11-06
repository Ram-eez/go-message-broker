package broker

import (
	"fmt"
	"sync"
	"time"
)

// Message that will be passed in the Subscriber Channel
type Message struct {
	Topic   string
	Payload interface{}
}

// This represents the cilent(us only, no multiple clients). This client will sub to various topics.
type Subscriber struct {
	Channel     chan interface{} // The channel here will have all the messages (payloads) which the user is subbed to.
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

// This method locks the mutex then makes the subscriber instance and finally appends the subscriber to the topic and unlocks the mutex.
func (b *Broker) Subscribe(topic string) *Subscriber {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	subscriber := &Subscriber{
		Channel:     make(chan interface{}, 1),
		Unsubscribe: make(chan bool),
	}

	b.subscribers[topic] = append(b.subscribers[topic], subscriber)

	return subscriber
}

// This method takes the slice of subscribers of x topic and closes its Channel and removes it form the slice of subs.
func (b *Broker) Unsubscribe(topic string, subscriber *Subscriber) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if subscribers, found := b.subscribers[topic]; found {
		for i, sub := range subscribers {
			if sub == subscriber {
				close(sub.Channel)
				b.subscribers[topic] = append(subscribers[:i], subscribers[i:]...)
				return
			}
		}
	}
}

func (b *Broker) Publish(topic string, payload interface{}) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if subscribers, found := b.subscribers[topic]; found {
		for _, sub := range subscribers {
			select {
			case sub.Channel <- payload:
			case <-time.After(time.Second):
				fmt.Printf("Subscriber slow. Unsubscribing from topic: %s\n", topic)
			}
		}
	}
}
