package slice_map

import "testing"

func BenchmarkSliceMapAdd(b *testing.B) {
	lm := NewLMap()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lm.Add(initDummy(i, i))
	}
}

func BenchmarkSliceMapDel(b *testing.B) {
	lm := NewLMap()
	for i := 0; i < 1e6; i++ {
		lm.Add(initDummy(i, i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lm.Del(i)
	}
}

func BenchmarkSliceMapIter(b *testing.B) {
	lm := NewLMap()
	for i := 0; i < 1e6; i++ {
		lm.Add(initDummy(i, i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lm.Iter(func(_ LMapObj) {})
	}
}

func BenchmarkSliceMapFastIter(b *testing.B) {
	lm := NewLMap()
	for i := 0; i < 1e6; i++ {
		lm.Add(initDummy(i, i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lm.FastIter(func(_ LMapObj) {})
	}
}

func BenchmarkMapAdd(b *testing.B) {
	mm := make(map[int]*Dummy)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mm[i] = initDummy(i, i)
	}
}

func BenchmarkMapDel(b *testing.B) {
	mm := make(map[int]*Dummy)
	for i := 0; i < 1e6; i++ {
		mm[i] = initDummy(i, i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		delete(mm, i)
	}
}

func BenchmarkMapIter(b *testing.B) {
	mm := make(map[int]*Dummy)
	for i := 0; i < 1e6; i++ {
		mm[i] = initDummy(i, i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, _ = range mm {
		}
	}
}
