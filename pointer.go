package _time

import (
	"time"
)

type Pointer int64

func (s Pointer) Floor(metka uint32) Pointer {
	if metka <= 0 {
		metka = 1
	}
	del := s % Pointer(metka)
	return s - del
}
func (s Pointer) Ceil(metka uint32) Pointer {
	if metka < 0 {
		metka = 0
	}
	del := s % Pointer(metka)
	return s + (Pointer(metka) - del)
}
func (s Pointer) String() string {
	return time.Unix(int64(s), 0).String()
}
func (s Pointer) Time() time.Time {
	return time.Unix(int64(s), 0)
}
func (s Pointer) Int64() int64 {
	return int64(s)
}