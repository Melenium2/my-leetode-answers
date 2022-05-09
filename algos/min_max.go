package algos

type numbers interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~byte
}

func min[T numbers](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func max[T numbers](a, b T) T {
	if a < b {
		return b
	}

	return a
}
