package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go/domain"
	"math/rand"
	"time"
)

func main() {

	channel1 := make(chan int, 1)
	channel2 := make(chan int, 1)

	go Producer(channel1)
	go Producer(channel2)


	go ConsumerAdd(channel1)
	go ConsumerAdd(channel1)

	go ConsumerSubtract(channel2)
	go ConsumerSubtract(channel2)

	//for {
	//	select {
	//	case <- channel1:
	//		go ConsumerAdd(channel1)
	//		time.Sleep(time.Second * 1)
	//	case <-channel2:
	//		go ConsumerSubtract(channel2)
	//		time.Sleep(time.Second * 1)
	//	}
	//}

	var input string
	fmt.Scanln(&input)

}

func Producer(producer chan<- int) {
	for {
		n := rand.Intn(100)
		log.Info("[PRODUCER] Se producen datos ", n)
		producer <- n
	}
}

func ConsumerAdd(consumer <-chan int) {
	var r int
	for {
		n := <-consumer
		domain.OperationsService{}.Add(&r, n)
		log.Info("[CONSUMER ADD] Se consumen datos ", n, r)
		time.Sleep(time.Second * 1)

	}
}

func ConsumerSubtract(consumer <-chan int) {
	var r int
	for {
		n := <-consumer
		domain.OperationsService{}.Subtract(&r, n)
		log.Info("[CONSUMER SUBTRACT] Se consumen datos ", n, r)
		time.Sleep(time.Second * 1)
	}
}
