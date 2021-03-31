package slice_map

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dummy struct {
	id  int
	val int

	p1 [16]*Dummy
	p2 string
	p3 []*Dummy
	p4 map[int]string
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

func (d *Dummy) LMapId() int {
	return d.id
}

func TestBaseFunction(t *testing.T) {
	lm := NewLMap()
	lm.Add(initDummy(1, 0))
	assert.Equal(t, 1, lm.Len())
	lm.Add(initDummy(1, 3))
	assert.Equal(t, 1, lm.Len())
	lm.Del(1)
	assert.Equal(t, 0, lm.Len())
	lm.Del(1)
	assert.Equal(t, 0, lm.Len())

	lm.Add(initDummy(1e8, 2e9))
	assert.Equal(t, 1, lm.Len())
	dObj, ok := lm.Get(1e8).(*Dummy)
	assert.Equal(t, true, ok)
	assert.Equal(t, int(1e8), dObj.id)
	assert.Equal(t, int(2e9), dObj.val)
}

func TestLargeKeys(t *testing.T) {
	lm := NewLMap()
	for i := 0; i < 1e6; i++ {
		lm.Add(initDummy(i, i*2-1))
	}
	assert.Equal(t, int(1e6), lm.Len())

	for i := int(1e6 - 1); i >= 0; i-- {
		lm.Del(i)
	}
	assert.Equal(t, int(0), lm.Len())
}

func TestIter(t *testing.T) {
	lm := NewLMap()
	seq := []int{9, 5, 2, 7, 1970, 1, 3, 654}
	for _, v := range seq {
		lm.Add(initDummy(v, v))
	}

	iterIdx := 0
	lm.FastIter(func(tmp LMapObj) {
		obj, ok := tmp.(*Dummy)
		assert.Equal(t, true, ok)
		assert.Equal(t, seq[iterIdx], obj.val)
		assert.Equal(t, seq[iterIdx], obj.id)
		iterIdx++
	})

	iterIdx = 0
	lm.Iter(func(tmp LMapObj) {
		obj, ok := tmp.(*Dummy)
		assert.Equal(t, true, ok)
		assert.Equal(t, seq[iterIdx], obj.val)
		assert.Equal(t, seq[iterIdx], obj.id)
		iterIdx++
	})

	assert.Equal(t, iterIdx, len(seq))
}
