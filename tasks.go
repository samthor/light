package light

import (
	"log"
	"time"
)

type Task struct {
	id       int // internal ID for ref
	Priority int
	Start    time.Time     // optional start time, if blank, always/now
	Duration time.Duration // optional duration (<=0 is forever)
	Color    HasColor
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

// Cancel this task. Returns true if a change was made. Safe on nil or zero Ref.
func (r *Ref) Cancel() bool {
	if r == nil || r.ts == nil {
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

// Debug dumps the light task list to console.
func Debug() {
	log.Printf("got tasks:")
	for _, t := range masterSystem.tasks {
		log.Printf("\t%+v", t)
	}
	log.Printf("-")
}
