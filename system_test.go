package light

import (
	"testing"
	"time"
)

func TestPriority(t *testing.T) {
	system := &taskSystem{}
	now := time.Now()

	if actual := system.step(now); !actual.Equal(nil) {
		t.Errorf("got %+v, expected initial nil", actual)
	}

	system.Add(Task{
		Color:    &Red,
		Priority: 0,
		Start:    now,
	})

	if expected, actual := &Red, system.step(now); !expected.Equal(actual) {
		t.Errorf("got %+v, expected %+v", actual, expected)
	}

	greenRef := system.Add(Task{
		Color:    &Green,
		Priority: 100,
		Start:    now,
	})

	if expected, actual := &Green, system.step(now); !expected.Equal(actual) {
		t.Errorf("got %+v, expected %+v", actual, expected)
	}
	if !greenRef.Cancel() {
		t.Errorf("expected change on Cancel")
	}
	if expected, actual := &Red, system.step(now); !expected.Equal(actual) {
		t.Errorf("got %+v, expected %+v", actual, expected)
	}
}

func TestExpiry(t *testing.T) {
	system := &taskSystem{}
	now := time.Now()

	system.Add(Task{
		Color:    &Blue,
		Priority: 0,
		Start:    now,
	})
	ref := system.Add(Task{
		Color:    &Red,
		Priority: 100,
		Start:    now,
		Duration: time.Minute,
	})

	if expected, actual := &Blue, system.step(now.Add(time.Hour)); !expected.Equal(actual) {
		t.Errorf("got %+v, expected %+v after one hour", actual, expected)
	}
	if ref.Cancel() {
		t.Errorf("expected no change on Cancel due to step side-effects")
	}
}
