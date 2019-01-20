package main

import (
	"fmt"
	"time"
)

func main() {
	mailbox := make(chan string)

	go printPing(mailbox)
	go printPong(mailbox)

	var input string
	_, _ = fmt.Scanln(&input)
}

func printPong(mailbox chan string) {

	for {
		fmt.Println("Pong ~> ", <-mailbox)
		mailbox <- "Pong"
		time.Sleep(time.Second * 1)
	}

}

func printPing(mailbox chan string) {

	for {
		mailbox <- "Ping"
		fmt.Println("Ping ~> ", <-mailbox)
	}

}
