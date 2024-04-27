package apostro

import (
	"strings"
	"unicode"
)

func Apos(data []string) []string{
	toJoin := strings.Join(data, " ")
	toRune := []rune(toJoin)

	// Front Apostrophy
	for i:=0; i <len(toRune)-1; i++{
		if unicode.IsPunct(toRune[i]) && toRune[i+1] == ' ' && toRune[i] == '\''{
			toRune[i+1], toRune[i] = toRune[i], toRune[i+1]
		}
	}
	toSplit := strings.Fields(string(toRune))

	return toSplit
}