package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DayLeft() int64 {

	d := time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	return int64(dur.Minutes())
}
