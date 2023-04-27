package format

import "strconv"

func Devices(input int) string {

	if input == -1 {
		return "All"
	}

	return strconv.Itoa(input)
}
