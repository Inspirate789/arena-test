package arena_test

import (
	"arena"
	"runtime/debug"
	"testing"
)

type Data struct {
	i  int
	s  string
	sl []float64
}

func NewDataNoGC(mem *arena.Arena) *Data {
	data := arena.New[Data](mem)
	data.i = 11
	data.s = "hello"
	data.sl = arena.MakeSlice[float64](mem, 100, 200)
	return data
}

func NewDataGC() *Data {
	return &Data{
		i:  11,
		s:  "hello",
		sl: make([]float64, 100, 200),
	}
}

func BenchmarkNoGC(b *testing.B) {
	mem := arena.NewArena()
	var data *Data
	var gcStats debug.GCStats
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data = NewDataNoGC(mem)
		b.StopTimer()
		debug.ReadGCStats(&gcStats)
		b.Log(gcStats.NumGC, gcStats.PauseTotal)
		b.StartTimer()
	}
	b.StopTimer()
	_ = data
	b.StartTimer()
	mem.Free()
}

func BenchmarkGC(b *testing.B) {
	var data *Data
	var gcStats debug.GCStats
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data = NewDataGC()
		b.StopTimer()
		debug.ReadGCStats(&gcStats)
		b.Log(gcStats.NumGC, gcStats.PauseTotal)
		b.StartTimer()
	}
	b.StopTimer()
	_ = data
	b.StartTimer()
	// runtime.GC()
}
