package iteration


func Repeat(character string, repititions int) string {
	var repeated string
	for i := 0; i < repititions; i++ {
		repeated += character
	}
	return repeated
}
