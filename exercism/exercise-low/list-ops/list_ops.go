package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) (res int) {
	if len(s) == 0 {
		return initial
	}
	res = initial
	for _, v := range s {
		res = fn(res, v)
	}
	return
}

func (s IntList) Foldr(fn func(int, int) int, initial int) (res int) {
	if len(s) == 0 {
		return initial
	}
	res = initial
	for _, v := range s.Reverse() {
		res = fn(v, res)
	}
	return
}

func (s IntList) Filter(fn func(int) bool) (res IntList) {
	res = []int{}
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}

	return
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) (res IntList) {
	res = []int{}
	for _, v := range s {
		res = append(res, fn(v))
	}

	return
}

func (s IntList) Reverse() (res IntList) {
	res = []int{}
	for i := range s {
		res = append(res, s[len(s)-i-1])
	}
	return
}

func (s IntList) Append(lst IntList) (res IntList) {
	res = []int{}
	res = append(res, s...)
	res = append(res, lst...)
	return
}

func (s IntList) Concat(lists []IntList) (res IntList) {
	res = []int{}
	res = append(res, s...)
	for _, v := range lists {
		res = append(res, v...)
	}
	return
}
