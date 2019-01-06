package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go/application"
	"go/domain"
	"go/functional"
	"math/rand"
	"time"
)



func main() {

	filter := functional.Filter(functional.PredicateOdd{}, 2)
	mapper := functional.Map(functional.Add2{}, 2)
	currying := functional.AddCurrying(2)( 6)(6)
	currying2 := functional.FilterOddAndAddXCurrying(functional.PredicateOdd{}, 3, 4, 5)(6)

	fmt.Println(filter)
	fmt.Println(mapper)
	fmt.Println(currying)
	fmt.Println(currying2)
	time.Sleep(time.Second * 5)


	channel1 := make(chan int, 1)
	channel2 := make(chan int, 1)

	go Producer(channel1)
	go Producer(channel2)

	go ConsumerAdd(channel1)
	go ConsumerAdd(channel1)

	go ConsumerSubtract(channel2)
	go ConsumerSubtract(channel2)

	for {
		select {
		case <-channel1:
			go ConsumerAdd(channel1)
			time.Sleep(time.Second * 1)
		case <-channel2:
			go ConsumerSubtract(channel2)
			time.Sleep(time.Second * 1)
		}
	}

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

		service := domain.OperationsService{}
		functions := application.HighOrderFunctionsCallBacks{}
		functions.Print("[CONSUMER ADD] Se consumen datos", service.Add, &r, n)(&r, n)

		time.Sleep(time.Second * 1)

	}
}

func ConsumerSubtract(consumer <-chan int) {

	var r int
	for {
		n := <-consumer

		service := domain.OperationsService{}
		functions := application.HighOrderFunctionsCallBacks{}
		functions.Print("[CONSUMER SUBTRACT] Se consumen datos", service.Subtract, &r, n)(&r, n)

		time.Sleep(time.Second * 1)
	}
}
