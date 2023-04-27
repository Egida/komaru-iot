package format

import "strconv"

func Cooldown(input int) string {

	if input == 0 {
		return "None"
	}

	return strconv.Itoa(input)
}
