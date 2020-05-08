package main

import (
	"fmt"
	"github.com/zakon47/_time"
	"time"
)

var t0 = time.Date(2020, time.May, 7, 22, 31, 59, 0, time.UTC)

func main() {
	//pointer()
	interval()
}
func pointer() {
	fmt.Println("Start: pointer")
	fmt.Println(t0, t0.Unix())
	fmt.Println("=================")
	rr := _time.Pointer(t0.Unix()).Floor(0)
	fmt.Println(rr)
	fmt.Println("")
}
func interval() {
	fmt.Println("Start: interval")
	t1 := t0
	t2 := t0.Add(-5 * time.Minute)

	fmt.Println(t1, t1.Unix())
	fmt.Println(t2, t2.Unix())
	fmt.Println("=================")

	interval := _time.Interval{t1, t2}
	l := interval.Linear(0)
	for _, v := range l {
		d := time.Unix(v, 0)
		fmt.Println(d, d.Unix())
	}
	fmt.Println("")
}
