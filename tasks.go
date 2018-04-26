package light

import (
	"time"
)

type Task struct {
	id       int // internal ID for ref
	Priority uint
	Color    *Color                     // color
	Get      func(time.Duration) *Color // callback to fetch color since start
	Start    time.Time                  // optional start time, if blank, always/now
	Duration time.Duration              // optional duration (<=0 is forever)
}

func (t Task) colorAt(at time.Time) *Color {
	var out *Color
	if t.Get != nil {
		since := at.Sub(t.Start)
		out = t.Get(since)
	}
	if out == nil {
		out = t.Color
	}
	return out
}

func (t Task) statusAt(at time.Time) (run, delete bool) {
	if t.Duration <= 0 {
		return true, false // always valid
	}

	sinceStart := at.Sub(t.Start)
	if sinceStart <= 0 {
		// it's coming up in the future
		return false, false
	} else if sinceStart < t.Duration {
		// yes it's now!
		return true, false
	} else {
		// done, remove me
		return false, true
	}
}

// Ref is a reference to a created Task.
type Ref struct {
	ts *taskSystem
	id int
}

// Cancel this task. Returns true if a change was made. Safe on nil Ref.
func (r *Ref) Cancel() bool {
	if r == nil {
		return false
	}
	return r.ts.Cancel(r.id)
}

// Run adds the task to the light system.
func Add(t Task) Ref {
	return masterSystem.Add(t)
}

// Update enacts any current state at the current time.
func Update() (*Color, error) {
	now := time.Now()
	color := masterSystem.step(now)
	ret := color
	if color == nil {
		color = &zeroColor
	}
	err := device.Set(*color)
	return ret, err
}

var (
	device weblightDevice
)
