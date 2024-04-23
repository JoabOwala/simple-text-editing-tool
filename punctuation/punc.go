package punctuation

import (
	"strings"
	"unicode"
)

func Punc(data []string)[]string{
	toJoin := strings.Join(data, " ")
	toRune := []rune(toJoin)

	for i:=0; i<len(toRune);i++{
		// Check if rune is a punctuation 
		// and if privious rune is a space
		if unicode.IsPunct(toRune[i]) && toRune[i-1] == ' '{
			toRune[i-1], toRune[i] = toRune[i], toRune[i-1]
		}

		// check for caseses of last punctuation
		if unicode.IsPunct(toRune[i]) && toRune[i-1] == ' ' && toRune[i] == '\''{
			toRune[i-1], toRune[i] = toRune[i], toRune[i-1]
		}
	}
	toSplit := strings.Fields(string(toRune))
	return toSplit
}