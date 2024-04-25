package vowel

func Vow(data []string) []string{
	myVowels := []string {"a","e","i","o","u","h"}

	for i, cha := range data{
		for _, cha2 := range myVowels{
			if cha == "a" && string(data[i+1][0]) == cha2{
				data[i] = "an"
			}else if cha == "A" && string(data[i+1][0]) == cha2{
				data[i] = "An"
			}else if cha == "an" && string(data[i+1][0]) != cha2{
				data[i] = "a"
			}else if cha == "An" && string(data[i+1][0]) != cha2{
				data[i] = "A"
			}
		}
		
	}
	return data
}