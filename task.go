package tlog

import "time"

type Task string

func (t Task) Extend(task string) Task {
	return Task(string(t) + "-" + task)
}

func (t Task) Start() (Task, time.Time) {
	return t, time.Now()
}
