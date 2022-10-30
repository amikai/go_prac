package switchex

import "time"

func typeSwtich(data interface{}) string {
	switch data.(type) {
	case bool:
		return "bool"
	case int:
		return "int"
	default:
		return "other"
	}
}

func switchWithWeekDay(t time.Time) string {
	// the underlying type of weekday is int
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return "weekend"
	default:
		return "weekday"
	}
}

func switchWithAbs(x int) int {
	switch {
	case x >= 0:
		return x
	default:
		return -x
	}
}
