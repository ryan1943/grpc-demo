package main

import (
	"errors"
	"fmt"
	"time"
)

type TimeStamp struct {
	StartTimestamp int64
	EndTimestamp   int64
}

//当天的0点和23:59:59的毫秒时间戳
func TodayNanoTime() TimeStamp {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)

	return TimeStamp{
		StartTimestamp: tm1.UnixNano() / 1e6,
		EndTimestamp:   tm2.UnixNano()/1e6 - 1,
	}
}

func TheMonthNanoTime() TimeStamp {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 1, 0)

	return TimeStamp{
		StartTimestamp: tm1.UnixNano() / 1e6,
		EndTimestamp:   tm2.UnixNano()/1e6 - 1,
	}
}

func MonthNanoTime(year int, month int) (TimeStamp, error) {
	if year <= 0 {
		return TimeStamp{}, errors.New("The year is greater than 0 ")
	}
	if month < 1 || month > 12 {
		return TimeStamp{}, errors.New("The month is between 1 and 12 ")
	}

	sMonth := time.Month(month)
	tm1 := time.Date(year, sMonth, 1, 0, 0, 0, 0, time.Local)
	tm2 := tm1.AddDate(0, 1, 0)

	return TimeStamp{
		StartTimestamp: tm1.UnixNano() / 1e6,
		EndTimestamp:   tm2.UnixNano()/1e6 - 1,
	}, nil
}

func main() {
	t, _ := MonthNanoTime(2018, 6)
	fmt.Println(t.StartTimestamp, t.EndTimestamp)
}
