package functional

type predicate interface {
	apply(interface{}) bool
}

type PredicateOdd struct {}

func (PredicateOdd) apply (value interface{}) bool {
	return value.(int)%2 == 0
}