package slice_map

const (
	_BASE_NO_SHRINK_SIZE = 1024
)

type LMapObj interface {
	LMapId() int
}

type LMap struct {
	objSlice []LMapObj
	idxMap   map[int]int
	maxIdx   int
}

func NewLMap() *LMap {
	return &LMap{
		objSlice: make([]LMapObj, 0),
		idxMap:   make(map[int]int),
		maxIdx:   0,
	}
}

func (lmap *LMap) Len() int {
	return lmap.maxIdx
}

func (lmap *LMap) Get(id int) LMapObj {
	if idx, present := lmap.idxMap[id]; present {
		return lmap.objSlice[idx]
	}
	return nil
}

func (lmap *LMap) Del(id int) {
	if curIdx, present := lmap.idxMap[id]; present {
		delete(lmap.idxMap, id)
		if curIdx == lmap.maxIdx-1 {
			lmap.objSlice[curIdx] = nil
			lmap.maxIdx--
		} else {
			lmap.maxIdx--
			obj := lmap.objSlice[lmap.maxIdx]
			lmap.objSlice[curIdx] = obj
			lmap.idxMap[obj.LMapId()] = curIdx
			lmap.objSlice[lmap.maxIdx] = nil
		}
	} else {
		// errors.New("Id not found.")
		return
	}

	// Shrink to prevent slice increasing with no limit.
	fvalue := len(lmap.objSlice) - lmap.maxIdx
	if lmap.maxIdx > _BASE_NO_SHRINK_SIZE && fvalue > 0 &&
		(float32(fvalue)/float32(lmap.maxIdx) > 0.1) {
		lmap.shrink()
	}
}

func (lmap *LMap) Add(obj LMapObj) {
	id := obj.LMapId()
	if idx, exist := lmap.idxMap[id]; exist {
		lmap.objSlice[idx] = obj
		return
	}

	lmap.idxMap[id] = lmap.maxIdx
	if lmap.maxIdx < len(lmap.objSlice) {
		lmap.objSlice[lmap.maxIdx] = obj
	} else {
		lmap.objSlice = append(lmap.objSlice, obj)
	}
	lmap.maxIdx++
}

func (lmap *LMap) shrink() {
	if lmap.maxIdx < len(lmap.objSlice) {
		lmap.objSlice = lmap.objSlice[0:lmap.maxIdx]
	}
}

func (lmap *LMap) Shrink() {
	lmap.shrink()
}

// Iterting map, as you must not delete any element in user-specified function.
func (lmap *LMap) FastIter(f func(LMapObj)) {
	for _, c := range lmap.objSlice {
		if c == nil {
			break
		}
		f(c)
	}
}

// Iterating map with the ability of deleting element.
// But some element would not be travelled, if you deleted keys.
func (lmap *LMap) Iter(f func(LMapObj)) {
	i := 0
	for _, c := range lmap.objSlice {
		if c == nil {
			break
		}

		f(c)
		if i < len(lmap.objSlice) {
			newC := lmap.objSlice[i]
			for newC != nil && newC.LMapId() != c.LMapId() {
				f(newC)
				c = newC
				if i < len(lmap.objSlice) {
					newC = lmap.objSlice[i]
				}
			}
		}
		i++
	}
}
