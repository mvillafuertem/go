package functional

func AddCurrying(n1 int) func(n2 int) func(n3 int) int {

	return func(n2 int) func(n3 int) int {

		return func(n3 int) int {

			return n1 + n2 + n3
		}
	}
}

func FilterOddAndAddXCurrying(p predicate, nums ...int) func(int) interface{} {

	filter := FilterOdd(p, nums)

	return func(n int) interface{} {

		AddX(filter, n)

		return filter

	}
}

func AddX(filter map[int]int, n int) {
	for e := range filter {
		filter[e] = e + n
	}
}

func FilterOdd(p predicate,nums []int) map[int]int {

	numsWith := make(map[int]int)

	for e := range nums {
		if p.apply(e) {
			numsWith[e] = e
		}
	}

	return numsWith
}
