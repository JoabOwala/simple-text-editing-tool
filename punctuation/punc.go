package punctuation

import (
	"strings"
	"unicode"
)
func Punc(data []string) []string {
	myJoin := strings.Join(data, " ")
	runeData := []rune(myJoin)
	for i := 0; i < len(runeData)-1; i++ {
		if unicode.IsPunct(runeData[i+1]) && runeData[i] == ' ' {
			runeData[i+1], runeData[i] = runeData[i], runeData[i+1]
		}
	}
	// Back apostrophy
	for i := 0; i < len(runeData)-1; i++ {
		if unicode.IsPunct(runeData[i+1]) && runeData[i] == ' ' && runeData[i+1] == '\'' {
			runeData[i+1], runeData[i] = runeData[i], runeData[i+1]
		}
	}
	toString := string(runeData)
	toSplit := strings.Fields(toString)
	return toSplit
}