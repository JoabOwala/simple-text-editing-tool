package upper

import (
	"strconv"
	"strings"
)

func Up(data []string) []string{
	for i:=0; i<len(data);i++{
		if strings.Contains(data[i], "(up"){
			if strings.Contains(data[i], "(up,"){
				//Convert the string to int
				// Remove trailing closing bracket
				toInt, _ := strconv.Atoi(data[i+1][:len(data[i+1])-1])

				for j := i-toInt; j<i;j++{
					data[j] = strings.ToUpper(data[j])
				}
				data = append(data[:i],data[i+2:]...)
				
				}else if strings.Contains(data[i], "(up)"){
					data[i-1] = strings.ToUpper(data[i-1])
					data = append(data[:i],data[i+1:]...)
				}
			}
	}
	return data
}