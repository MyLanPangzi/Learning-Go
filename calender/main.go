package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	year  int
	month int
	day   int
	now   = time.Now()
)

func init() {
	flag.IntVar(&year, "year", now.Year(), "year of calender. default current year")
	flag.IntVar(&month, "month", int(now.Month()), "month of calender. default current month")
	flag.IntVar(&day, "day", now.Day(), "day of calender. default current day")
}

func main() {
	flag.Parse()

	now = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	fmt.Println(now.YearDay(), now.Day(), ChineseOfWeekday(int(now.Weekday())))
	fmt.Println(getCalenderOfMonth(now))

	for i := 1; i <= 12; i++ {
		date := time.Date(now.Year(), time.Month(i), 10, 0, 0, 0, 0, time.Local)
		fmt.Println(date, ChineseOfWeekday(int(date.Weekday())))
		//fmt.Println(getCalenderOfMonth(date))
	}

}

func getCalenderOfMonth(now time.Time) *strings.Builder {
	firstDay := firstDayOfMonth(now)
	s := &strings.Builder{}
	for i := 1; i <= 7; i++ {
		s.WriteString(fmt.Sprintf("%s\t", ChineseOfWeekday(i)))
	}
	s.WriteString("\n")
	s.WriteString(tabsOfWeekday(firstDay.Weekday()))
	for i, j := 0+1, firstDay.Weekday(); i <= DaysOfMonth(now); i++ {
		s.WriteString(fmt.Sprintf("%d\t", i))
		if j%7 == 0 {
			s.WriteString("\n")
		}
		j++
	}
	return s
}

func tabsOfWeekday(weekday time.Weekday) string {
	switch weekday {
	case time.Monday, time.Tuesday, time.Thursday, time.Wednesday, time.Friday, time.Saturday:
		return strings.Repeat("\t", int(weekday)-1)
	default:
		return strings.Repeat("\t", 6)
	}
}

func firstDayOfMonth(now time.Time) time.Time {
	return now.AddDate(0, 0, -now.Day()+1)
}

func ChineseOfWeekday(i int) string {
	switch i {
	case 1:
		return "周一"
	case 2:
		return "周二"
	case 3:
		return "周三"
	case 4:
		return "周四"
	case 5:
		return "周五"
	case 6:
		return "周六"
	case 7, 0:
		return "周日"
	default:
		panic("unknown weekday " + strconv.Itoa(i))
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
