package Utils

func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func Prepend[T string](element T, slice []T) []T {
	return append([]T{element}, slice...)
}
