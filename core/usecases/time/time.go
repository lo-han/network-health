package time_usecase

import "time"

type Time interface {
	Now() time.Time
}

type GoTime struct{}

func (*GoTime) Now() time.Time {
	return time.Now()
}
