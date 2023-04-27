package sessions

func count(name string) int {
	i := 0
	for _, s := range sessions {
		if s.Name == name {
			i++
		}
	}
	return i
}
