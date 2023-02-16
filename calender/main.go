package main

import (
	"time"
)

func main() {
	now := time.Now()
	for i := 0 + 1; i <= DaysOfMonth(now); i++ {

	}
}

func DaysOfMonth(t time.Time) int {
	switch t.Month() {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	default:
		if isLeap(t.Year()) {
			return 29
		}
		return 28
	}

}

func isLeap(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
