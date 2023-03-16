package iteration

func Repeat(a string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += a
	}
	return result
}
