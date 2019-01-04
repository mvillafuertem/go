package main

import (
	"fmt"
	"go/application"
	"go/domain"
	"sync"
)

func main() {

	service := domain.OperationsService{}
	routines := application.GoRoutines{Operations: service}

	(&sync.WaitGroup{}).Add(1)

	go routines.AddRoutine(1, 2, 3, 4, 5)
	go routines.AddRoutine(10, 1, 2, 3, 4)
	go routines.SubtractRoutine(10, 1, 2, 3, 4)

	(&sync.WaitGroup{}).Wait()

	fmt.Println()
}
