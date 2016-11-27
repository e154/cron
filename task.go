package cron

import (
	"strings"
)

type Task struct {
	_time	map[int][]int
	_func	func()
	cron	*Cron
	enabled bool
}

func (t *Task) Enable() *Task {
	t.enabled = true
	return t
}

func (t *Task) Disable() *Task {
	t.enabled = false
	return t
}

func (t *Task) Enabled() bool {
	return t.enabled
}

func (t *Task) SetTime(time string) {
	args := strings.Split(time, " ")
	switch len(args) {
	case 1:
		if args[0] == "*" {

		}

	}
}

func (t *Task) GetTime() (time string) {

	return
}

func (t *Task) exec() {

	//log.Println("exec")

	// WEEKDAY
	exist := false
	for _, weekday := range t._time[WEEKDAY] {
		if weekday == int(t.cron.weekday) {
			exist = true
			break
		}
	}

	if !exist {
		return
	}

	//log.Println("weekday")

	// MONTH
	exist = false
	for _, month := range t._time[MONTH] {
		if month == int(t.cron.month) {
			exist = true
			break
		}
	}
	if !exist {
		return
	}

	//log.Println("month")

	// DAY
	exist = false
	for _, day := range t._time[DAY] {
		if day == t.cron.day {
			exist = true
			break
		}
	}
	if !exist {
		return
	}

	//log.Println("day")

	// HOUR
	exist = false
	for _, hour := range t._time[HOUR] {
		if hour == t.cron.hour {
			exist = true
			break
		}
	}
	if !exist {
		return
	}

	//log.Println("hour")

	// MINUTES
	exist = false
	for _, min := range t._time[MINUTE] {
		if min == t.cron.min {
			exist = true
			break
		}
	}
	if !exist {
		return
	}

	//log.Println("minutes")

	// SECONDS
	exist = false
	for _, sec := range t._time[SECOND] {
		if sec == t.cron.second {
			exist = true
			break
		}
	}
	if !exist {
		return
	}

	//log.Println("seconds")

	t.Run()
}

func (t *Task) Run() {
	t._func()
}