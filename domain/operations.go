package domain

type Operations interface {
	Add(*int, ... int)
	Subtract(*int, ... int)
}


type OperationsService struct {}

func (o OperationsService) Add(r *int, nums... int){

	for n := range nums {
		*r += nums[n]
	}
}

func (o OperationsService) Subtract(r *int, nums... int) {

	for n := range nums {
		*r -= nums[n]
	}
}