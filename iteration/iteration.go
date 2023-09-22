package iteration

func Repeat(word string, count int) string {

	var repeated string

	if count <= 0 {
		return "Input a number above 0"
	}

	for i := 0; i < count; i++ {
		repeated += word
	}
	return repeated
}
