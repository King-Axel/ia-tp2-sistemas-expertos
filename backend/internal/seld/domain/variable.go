package domain

type Universe struct {
	Min float64
	Max float64
}

func (universe Universe) Clamp(value float64) float64 {
	if value < universe.Min {
		return universe.Min
	}

	if value > universe.Max {
		return universe.Max
	}

	return value
}

type Variable struct {
	Key      string
	Name     string
	Universe Universe
	Sets     []FuzzySet
}

func (variable Variable) SetByKey(setKey string) (FuzzySet, bool) {
	for _, set := range variable.Sets {
		if set.Key == setKey {
			return set, true
		}
	}

	return FuzzySet{}, false
}
