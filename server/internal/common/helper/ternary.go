package helper

func Ternary[T any](condition bool, right, left T) T {
	if condition {
		return right
	}

	return left
}
