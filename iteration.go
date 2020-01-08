package iteration

//Repeat : repeats a character x times
func Repeat(character string) (repeated string) {
	for index := 0; index < 5; index++ {
		repeated += character
	}
	return
}
