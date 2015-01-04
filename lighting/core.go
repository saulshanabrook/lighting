package core

type DimmerMap map[int]int

type State interface {
	DimmersAt() (DimmerMap, error)
}

type StateStack struct {
	states []State
}

func (s StateStack) DimmersAt() (mergedDimmers DimmerMap, err error) {
	mergedDimmers = make(DimmerMap)
	for _, state := range s.states {
		var dimmers DimmerMap
		if dimmers, err = state.DimmersAt(); err != nil {
			return
		}
		for k, v := range dimmers {
			mergedDimmers[k] = v
		}
	}
	return
}
