package iteration

func Repeat(a string, repeatCount int) string {
	var result string
	for i := 0; i < repeatCount; i++ {
		result += a
	}
	return result
}
