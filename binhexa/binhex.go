package binhexa

import "strconv"

func Bin(data []string) []string{
	for i:=0; i<len(data);i++{
		if data[i] == "(bin)"{
			toBin, _ :=strconv.ParseInt(data[i-1], 2, 16)

			data[i-1] = strconv.Itoa(int(toBin))
			data = append(data[:i],data[i+1:]...)
		}
	}
	return data
}