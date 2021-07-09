package services

func isInterval(stack [4]string, needle string) bool {
	for _, v := range stack {
		if v == needle {
			return true
		}
	}

	return false
}
