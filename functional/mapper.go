package functional

type mapper interface {
	apply(interface{}) interface{}
}

type Add2 struct {}

func (Add2) apply(i interface{}) interface{}{
	return i.(int) + 2
}