package clock

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (realClocker RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

func (fixedClocker FixedClocker) Now() time.Time {
	return time.Date(2022, 5, 10, 12, 34, 56, 0, time.UTC)
}
