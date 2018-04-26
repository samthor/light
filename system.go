package light

import (
	"sort"
	"sync"
	"time"
)

var (
	masterSystem = &taskSystem{}
)

// taskList sorts by priority (high first) and then time (low first).
type taskList []Task

func (tl taskList) Len() int      { return len(tl) }
func (tl taskList) Swap(i, j int) { tl[i], tl[j] = tl[j], tl[i] }
func (tl taskList) Less(i, j int) bool {
	a, b := tl[i], tl[j]
	if a.Priority > b.Priority {
		return true
	} else if a.Priority == b.Priority && a.Start.Before(b.Start) {
		return true
	}
	return false
}

type taskSystem struct {
	lock     sync.Mutex
	nextTask int
	tasks    taskList
}

func (ts *taskSystem) Cancel(id int) bool {
	ts.lock.Lock()
	defer ts.lock.Unlock()

	for i := range ts.tasks {
		if ts.tasks[i].id == id {
			ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (ts *taskSystem) Add(t Task) Ref {
	ts.lock.Lock()
	defer ts.lock.Unlock()

	if t.Start.IsZero() {
		// we always need a real start to calculate end time
		t.Start = time.Now()
	}
	ts.nextTask++
	t.id = ts.nextTask

	ts.tasks = append(ts.tasks, t)
	sort.Sort(ts.tasks)
	return Ref{ts, t.id}
}

func (ts *taskSystem) step(at time.Time) *Color {
	ts.lock.Lock()
	defer ts.lock.Unlock()

retry:
	for i, t := range ts.tasks {
		run, delete := t.statusAt(at)
		if delete {
			// rinse and repeat to pop off dead tasks
			ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
			goto retry
		}

		if run && t.Color != nil {
			since := at.Sub(t.Start)
			out := t.Color.colorAt(since)
			if out != nil {
				return out
			}
		}
	}

	return nil
}
