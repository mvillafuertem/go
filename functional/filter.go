package functional

func Filter(p predicate, i interface{}) bool {
	return p.apply(i)
}