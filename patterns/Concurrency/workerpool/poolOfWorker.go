package main

import (
	"context"
	"math/rand"
	"time"

	"mazxcv.github.com/learn/patterns/Concurrency/workerpool/wp"
)

type Message struct {
	Id    int
	Title string
	Body  string
}

type IPool interface {
	Create()
	Handle(Message)
	Wait()
	Stats()
}

func processMessage(workerId int, message Message) {
	time.Sleep(100 * time.Millisecond)
	println("Worker", workerId, "Processing message", message.Id, message.Title, message.Body)
}

var messCounter int

func getMessages() []Message {
	messCount := rand.Intn(200)

	messages := make([]Message, 0, messCount)
	for range messCount {
		messCounter++
		messages = append(messages, Message{messCounter, "Title" + string(messCounter), "Body" + string(messCounter)})
	}

	return messages
}

func main() {

	var pool IPool

	pool = wp.New(processMessage)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

l:
	for {
		select {
		case <-ctx.Done():
			break l
		default:
		}
		mess := getMessages()

		pool.Create()

		for _, message := range mess {
			pool.Handle(message)
		}

		pool.Wait()
	}

	pool.Stats()
}
