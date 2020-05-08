package _time

import (
	"time"
)

type Interval struct {
	From time.Time
	To   time.Time
}

//получить линейный интервал по точкам
func (s *Interval) Linear(step uint32) []int64 {
	t1 := s.From.Unix()
	t2 := s.To.Unix()
	if step <= 0 {
		return []int64{}
	}
	if t1 < t2 {
		t1 = Pointer(s.From.Unix()).Ceil(step).Int64()
		t2 = Pointer(s.To.Unix()).Floor(step).Int64()
		seg := (t2 - t1) / int64(step)
		arr := make([]int64, seg+1)
		for i := 0; i < cap(arr); i++ {
			n := t1 + int64(step)*int64(i)
			arr[i] = n
		}
		return arr
	} else {
		t1 = Pointer(s.From.Unix()).Floor(step).Int64()
		t2 = Pointer(s.To.Unix()).Ceil(step).Int64()
		seg := (t1 - t2) / int64(step)
		arr := make([]int64, seg+1)
		for i := cap(arr) - 1; i >= 0; i-- {
			n := t1 - int64(step)*int64(i)
			arr[i] = n
		}
		return arr
	}
}

func (s *Interval) Double(step uint32) []int64 {
	return nil
}
