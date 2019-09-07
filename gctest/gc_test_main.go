package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/legougames/slice_map"
)

type Dummy struct {
	id  int
	val int

	p1 [16]*Dummy
	p2 string
	p3 []*Dummy
	p4 map[int]string
}

func (d Dummy) LMapId() int {
	return d.id
}

func initDummy(id, val int) *Dummy {
	return &Dummy{
		id:  id,
		val: val,
		p2:  fmt.Sprintf("%d_%d", id, val),
		p3:  []*Dummy{&Dummy{id: 1, val: 0}, &Dummy{id: 1, val: 0}},
		p4: map[int]string{
			id + val + 1:     fmt.Sprintf("%d_%d", id+1, val),
			id + val + 10086: fmt.Sprintf("%d_%d", id+10086, val),
			id + val - 1e6:   fmt.Sprintf("%d_%d", id-1e6, val),
			id + val + 1e7:   fmt.Sprintf("%d_%d", id+1e7, val),
		},
	}
}

func showGCStats(info debug.GCStats) string {
	return fmt.Sprintf("GC NumGC %v PauseTotal %v.", info.NumGC, info.PauseTotal)
}

func sliceMapGCTest() {
	lm := slice_map.NewLMap()
	debug.SetGCPercent(0)

	for i := 0; i < 500000; i++ {
		lm.Add(initDummy(i, i))
	}

	fmt.Println("Init slice-map done.")
	t := time.Now()
	debug.SetGCPercent(100)

	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Printf("Slice map gc x 10 cost: %v.\n", time.Now().Sub(t))
}

func mapGCTest() {
	mm := make(map[int]*Dummy)
	debug.SetGCPercent(0)

	for i := 0; i < 500000; i++ {
		mm[i] = initDummy(i, i)
	}

	fmt.Println("Init map done.")
	t := time.Now()
	debug.SetGCPercent(100)

	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Printf("Go map gc x 10 cost: %v.\n", time.Now().Sub(t))
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "lm" {
		fmt.Println("Do test using slice-map.")
		sliceMapGCTest()
	} else if len(os.Args) > 1 && os.Args[1] == "map" {
		fmt.Println("Do test using go map.")
		mapGCTest()
	} else {
		fmt.Println("Usage ./gc_test_main.go [lm|map]")
	}
}
