package domain

import log "github.com/sirupsen/logrus"

type Operations interface {
	Add(... int) int
	Subtract(... int) int
}


type OperationsService struct {}

func (o OperationsService) Add(nums... int) (r int){

	for n := range nums {
		log.Debugf("%d add %d", r, n)
		r += nums[n]
	}

	return
}

func (o OperationsService) Subtract(nums... int) (r int) {

	for n := range nums {
		log.Debugf("%d subtract %d", r, n)
		r -= nums[n]
	}

	return
}