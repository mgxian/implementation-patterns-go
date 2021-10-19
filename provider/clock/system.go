package clock

import "time"

type System struct{}

func NewSystem() System {
	return System{}
}

func (s System) Now() time.Time {
	return time.Now()
}
