package main

import (
	"fmt"
	"go-message-broker/broker"
	"time"
)

func main() {
	broker := broker.NewBroker()
	subscriber := broker.Subscribe("thenews")

	go func() {
		for {
			select {
			case msg, ok := <-subscriber.Channel:
				if !ok {
					fmt.Println("Subscriber Channel is closed")
					return
				}
				fmt.Printf("Received the message: %v\n", msg)
			case <-subscriber.Unsubscribe:
				fmt.Println("Unsubscribed from the channel")
				return
			}
		}
	}()

	broker.Publish("thenews", "hello world news")
	broker.Publish("thenews", "doge is gay")

	time.Sleep(time.Second)
	broker.Unsubscribe("thenews", subscriber)

	broker.Publish("thenews", "doge is not gay")
	time.Sleep(time.Second)

}
