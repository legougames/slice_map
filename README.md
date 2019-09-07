# slice_map

The `slice_map` is designed to impove the performance for the case that you need to hold a large map in memory.
It uses a slice and an extra index map to replace the go-standard map such as `map[int]interface` for better `add/iter` performance.

If you have a large map (more than 500k keys and use complex struct as map value), `slice_map` saves lots of cpu time when every GC excuting.

Adding/deleting elements:

```go
type OBJ struct {
	// ...
    int id
}

// Need to implement a Id interface.
func (o *Obj) LMapId() int {
	return o.id
}

lm := NewLMap()
newObj := OBJ{ id: 1024 }
lm.Add()
lm.Del(newObj.id)
lm.Len()
```

Iter the slice_map:
When use `FastIter`, you can't delete keys in the map. And get 10 times faster than `Iter` which you can delete keys safely during the iterting.

```go
lm.FastIter(func(tmp LMapObj) {
	obj, ok := tmp.(*OBJ)
	...
})

lm.Iter(func(tmp LMapObj) {
	obj, ok := tmp.(*OBJ)
	...
})
```

