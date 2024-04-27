package lower

import (
	"strconv"
	"strings"
)

func Low(data []string) []string{
	for i:=0;i<len(data);i++{
		if strings.Contains(data[i], "(low"){
			if strings.Contains(data[i], "(low,"){
				// Convert the sting to int and remove trailing bracket
				toInt, _ := strconv.Atoi(data[i+1][:len(data[i+1])-1])

				// Use the int to convert previous value into lower case
				for j:=i-toInt; j<i;j++{
					data[j] =strings.ToLower(data[j])
				}
				data = append(data[:i],data[i+2:]...)
				}else if strings.Contains(data[i], "(low)"){
					data[i-1] = strings.ToLower(data[i-1])
					data = append(data[:i],data[i+1:]...)
				}
			}
	}
	return data
}