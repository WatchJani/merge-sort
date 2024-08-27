package main

import "testing"

// 312ns
func Benchmark(b *testing.B) {
	b.StopTimer()

	memTable := []Read{{"M", 5}, {"M", 152}, {"M", 547}, {"M", 4114}, {"M", 5451}}
	memSort := Sort{
		memTable, 0,
	}

	freezeTable := []Read{{"M", 2}, {"M", 100}, {"M", 647}, {"M", 3114}, {"M", 6451}}
	freezeSort := Sort{
		freezeTable, 0,
	}

	discTable := []Read{{"D", 2000}, {"C", 2020}, {"D", 3201}}
	discSort := Sort{
		discTable, 0,
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		MergeSort([]Sort{memSort, freezeSort, discSort})
	}
}
