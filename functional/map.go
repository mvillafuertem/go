package functional



func Map(m mapper, i interface{}) interface{} {

	return m.apply(i)
}
