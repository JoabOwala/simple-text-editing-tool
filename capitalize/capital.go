package capitalize

import (
	"strconv"
	"strings"
)

func Cap(data []string)[]string{
	for i:=0; i<len(data); i++{
		if strings.Contains(data[i], "(cap"){
			if strings.Contains(data[i], "(cap,"){
				// Convert next index to int and remove the (cap, _)
				toInt, _ := strconv.Atoi(data[i+1][:len(data[i+1])-1])

				// using the int obtained Title the previous Strings
				for j:= i-toInt; j<i;j++{
					data[j] = strings.Title(data[j])
				}
				data = append(data[:i], data[i+2:]...)

			}else if strings.Contains(data[i], "(cap)"){
				data[i-1] = strings.Title(data[i-1])
				data = append(data[:i], data[i+1:]...)
			}
			
		}
	}
	return data
}