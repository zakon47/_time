package main

import (
	"fmt"
	"github.com/zakon47/time_interval"
	"time"
)

var t0 = time.Date(2020, time.May, 7, 22, 31, 59, 0, time.UTC)

func main() {
	mxx := time_interval.NewMarker1D("15s")
	min := time_interval.NewMarker1D("10h")
	max := time_interval.NewMarkerD1("m2")
	fmt.Println(mxx, min, max)
	fmt.Printf("NEXT\n\n")
	//interval()
	pointer()
}
func pointer() {
	fmt.Println("######## Start: pointer")
	rr1 := time_interval.Pointer(t0.Unix()).Ceil(15)
	rr2 := time_interval.Pointer(t0.Unix()).Floor(15)
	fmt.Println("ceil:	", time.Unix(rr1.Int64(), 0), rr1)
	fmt.Println("~15s	", t0.Local(), t0.Unix())
	fmt.Println("floor:	", time.Unix(rr2.Int64(), 0), rr2)
	fmt.Println("")
}
func interval() {
	fmt.Println("######## Start: interval")
	t1 := t0
	t2 := t0.Add(-2 * time.Minute)

	fmt.Println("~", t1.Local(), t1.Unix())
	fmt.Println("~", t2.Local(), t2.Unix())

	interval := time_interval.NewInterval(t1, t2)
	l := interval.Linear(time_interval.NewMarker1D("30s").GetSeconds())
	for _, v := range l {
		d := time.Unix(v, 0)
		fmt.Println(d, d.Unix())
	}
	fmt.Println(interval)
}
