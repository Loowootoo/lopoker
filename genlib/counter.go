package genlib

import "time"

type TimeCounter struct {
	start time.Time
	count time.Duration
}

func NewCounter(ms int) *TimeCounter {
	tc := new(TimeCounter)
	tc.start = time.Now()
	tc.count = time.Millisecond * time.Duration(ms)
	return tc
}

func (t *TimeCounter) ResetCounter(ms int) {
	t.start = time.Now()
	t.count = time.Millisecond * time.Duration(ms)
}

func (t *TimeCounter) TimeUp() bool {
	now := time.Now()
	nextStepAt := t.start.Add(t.count)
	if now.Sub(nextStepAt) > 0 {
		t.start = now
		return true
	}
	return false
}
