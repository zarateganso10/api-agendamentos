package format

import "strconv"

func QueryParamsToInt(str string) int {
	number, _ := strconv.Atoi(str)

	return number
}
