package iteration

//Repeat : repeats a character x times
func Repeat(character string, times int) (repeated string) {
	for index := 0; index < times; index++ {
		repeated += character
	}
	return
}
