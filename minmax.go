package intutils

// MinOf возвращает минимальное значение из входного набора целых чисел.
func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}
