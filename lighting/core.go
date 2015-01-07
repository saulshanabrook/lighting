package core

import "fmt"

const (
	// NumDimmers is the maximum number of dimmer addresses possible
	NumDimmers = 512
)

// A Dimmer is the address for a single DMX output
type Dimmer int

// Level is the type used for a dimmer level.
type Level int

// DimmerMap is store a map of dimmer addresses to dimmer levels
type DimmerMap map[Dimmer]Level

// A State is a representation of light levels, that can be reduced down to
// a DimmerMap
type State interface {
	DimmersAt() (DimmerMap, error) // DimmersAt returns the DimmersMap of this state
}

// A StateStack is the combination of many States. They are held in a Slice
// that prioritizes the last State
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

// SingleDimmer holds one dimmer at a certain percent
type SingleDimmer struct {
	dimmer Dimmer
	level  Level
}

func checkDimmerAddress(address Dimmer) error {
	if address < 1 || address > NumDimmers {
		return fmt.Errorf("Dimmer addresses must be between 1 and %v", NumDimmers)
	}
	return nil
}

// DimmersAt will return that single dimmer at that percent
func (d SingleDimmer) DimmersAt() (dm DimmerMap, err error) {
	if err = checkDimmerAddress(d.dimmer); err != nil {
		return
	}
	dm = make(DimmerMap)
	dm[d.dimmer] = d.level
	return
}

func (d SingleDimmer) String() string {
	return fmt.Sprintf("Dimmer %v@%v", d.dimmer, d.level)
}

type Tags map[string]string
type Patch map[Dimmer]Tags

type Filter struct {
	Patch *Patch
	Tags  Tags
	level int
}

func isSubset(src, other map[string]string) bool {
	for k, v := range other {
		if other[k] != v {
			return false
		}
	}
	return true
}

func (f Filter) DimmersAt() (dm DimmerMap, err error) {
	dm = make(DimmerMap)
	p := *f.Patch

	// iterate through all dimmers in the patch
	for d, t := range p {
		// if the dimmer is tagged with all of the tags specified, then we want that one
		if isSubset(t, f.Tags) {
			dm[d] = f.level
		}
	}
	return
}

type Scene struct {
	Name string
	StateStack
}
