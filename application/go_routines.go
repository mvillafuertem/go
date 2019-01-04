package application

import (
	log "github.com/sirupsen/logrus"
	"go/domain"
	"sync"
)

type GoRoutines struct { Operations domain.Operations }

func (g GoRoutines) AddRoutine(nums... int) (r int) {

	(&sync.WaitGroup{}).Wait()

	log.Warn("IN ADD")

	r = g.Operations.Add(nums...)

	log.Warn("OUT ADD")

	return
}

func (g GoRoutines) SubtractRoutine(nums... int) (r int) {

	log.Warn("IN SUBTRACT")

	r = g.Operations.Subtract(nums...)

	log.Warn("OUT SUBTRACT")

	return
}
