package wordnet

func strIn(bag []string, word string) bool {
	for _, w := range bag {
		if word == w {
			return true
		} else {
			continue
		}
	}
	return false
}
