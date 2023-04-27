package format

import "strconv"

func Time(input int) string {

	if input == -1 {
		return "86400"
	}

	return strconv.Itoa(input)
}
