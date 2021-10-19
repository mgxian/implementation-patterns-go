package clock

import "time"

type Stub struct {
	now time.Time
}

func NewStub() Stub {
	return Stub{now: time.Now()}
}

func (s Stub) Now() time.Time {
	return s.now
}

func (s *Stub) SetNow(t time.Time) {
	s.now = t
}
