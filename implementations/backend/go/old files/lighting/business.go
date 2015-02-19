package business

type Dimmer int
type Level float

const DefaultLevel = 1

type ColorLabel string

type Tag string
type Value string
type Query map[Tag]Value

// DimmerMap is store a map of dimmer addresses to dimmer levels
type DimmerMap map[Dimmer]Level

type Levelable interface {
	GetLevel() Level
}

type DirectLevel struct {
	Level Level
}

func (dl *DirectLevel) Level() Level {
	return dl.Level
}

type ColorResolver interface {
	GetLevel(ColorLabel) Level
}

type ColorLevel struct {
	Label ColorLabel
	ColorResolver
}

func (cl *ColorLevel) Level() Level {
	return ColorResolver.GetLevel(cl.Label)
}

type System interface {
	DimmersAt() DimmerMap
}

type LevelledSystem struct {
	Levelable
	System
}

func (ls *LevelledSystem) DimmersAt() (dm DimmerMap) {
	dm = make(DimmerMap)
	l_modifier := Levelable.GetLevel()
	for d, l_original := range ls.System.DimmersAt() {
		dm[d] = l_original * l_modifier
	}
}

type DimmerSystem struct {
	Dimmer Dimmer
}

func (ds *DimmerSystem) DimmersAt() {
	dm := make(DimmerMap)
	dm[ds.Dimmer] = DefaultLevel
}

type QueryResolver interface {
	GetDimmers(Query) []Dimmer
}

type FilterSystem struct {
	Query Query
	QueryResolver
}

func (fs *FilterSystem) DimmersAt() (dm DimmerMap) {
	dm = make(DimmerMap)

}
