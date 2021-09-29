package main

import (
	"fmt"
	"runtime"
	"sort"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type PairList []Suggestion
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].score < p[j].score }

func firstSuggestions(m Suggestions, size int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Suggestion{k, v}
		i += 1
	}
	sort.Sort(sort.Reverse(p))
	return p[:size]
}