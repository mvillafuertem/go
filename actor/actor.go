package main

import (
	"fmt"
	"time"
)

// Actor = goroutine
// Mailbox = channel

func main() {

	mailbox := make(chan string)
	go generator(mailbox)

	actor1 := NewActor("Pepe")
	actor2 := NewActor("Maria")

	actor1.mailbox = mailbox
	actor2.mailbox = mailbox

	// goroutine principal
	var input string
	_, _ = fmt.Scanln(&input)
}

func generator(mailbox chan string) {
	for {
		mailbox <- "hola"
	}
}

type Actor struct {
	name    string
	mailbox chan string
}

func NewActor(name string) *Actor {
	actor := Actor{name, make(chan string)}
	go actor.inbox()
	return &actor
}

func (a *Actor) inbox() {
	for {
		fmt.Printf("Hi, I am %s this is my message ~ %s\n", a.name, <-a.mailbox)
		time.Sleep(time.Second * 1)
	}
}
