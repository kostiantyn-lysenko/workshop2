package services

import (
	"time"
	"workshop2/errs"
)

func identifyLimit(interval string, now time.Time) (time.Time, error) {
	var limit time.Time

	switch interval {
	case "day":
		limit = now.AddDate(0, 0, -1)
	case "week":
		limit = now.AddDate(0, 0, -7)
	case "month":
		limit = now.AddDate(0, -1, 0)
	case "year":
		limit = now.AddDate(-1, 0, 0)
	default:
		return limit, errs.NewBadIntervalError()
	}

	return limit, nil
}
