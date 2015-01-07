
package business


type Dimmer int
type Level float


// DimmerMap is store a map of dimmer addresses to dimmer levels
type DimmerMap map[Dimmer]Level

type Modifier interface {
	Modify(previous, current *DimmerMap, level Level)
}

type Level interface {
	GetPercent() float
}

type System interface {
	DimmersAt() DimmerMap
}

type LevelledSystem struct {
	Modifier
	Level
	System
}

func (ls *LevelledSystem) DimmersAt() (dm DimmerMap) {
	ls.Level.float
}