package format

func Bool(input bool) string {

	if input == false {
		return "\033[91mfalse\033[0m"
	}

	return "\033[32mtrue\033[0m"
}
