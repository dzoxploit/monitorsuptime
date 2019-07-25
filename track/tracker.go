package track

import (
	"time"
)

type TimeTracker struct {
	delayer Delayer
	nextTime time.Time

	counter int
}

func (t *TimeTracker) IsReady bool {
	if time.Now().After(t.nextTime){
		return true
	}
	return false
}

func (t *TimeTracker) SetNext() (time.Duration, time.Time) {
	t.counter++
	nextDelay := t.delayer.Delay()
	t.nextTime = time.Now().Add(nextDelay)
	return nextDelay, t.nextTime
}

func NewTracker(delayer Delayer) *TimeTracker {
	return &TimeTracker{delayer: delayer}
}

func (t *TimeTracker) HasBeenRan() bool {
	if t.counter > 0 {
		return true
	}
	return false
}

