package services

import "time"

var intervals = [4]string{"day", "week", "month", "year"}

func isInterval(stack [4]string, needle string) bool {
	for _, v := range stack {
		if v == needle {
			return true
		}
	}

	return false
}

func identifyLimit(interval string) time.Time {
	var limit time.Time

	switch interval {
	case "day":
		limit = time.Now().AddDate(0, 0, -1)
	case "week":
		limit = time.Now().AddDate(0, 0, -7)
	case "month":
		limit = time.Now().AddDate(0, -1, 0)
	case "year":
		limit = time.Now().AddDate(-1, 0, 0)
	}

	return limit
}
