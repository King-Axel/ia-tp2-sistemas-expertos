package engine

import "iteraciones/backend/internal/seld/domain"

type Defuzzifier struct{}

func NewDefuzzifier() *Defuzzifier {
	return &Defuzzifier{}
}

func (defuzzifier *Defuzzifier) Centroid(aggregatedMemberships map[float64]float64) float64 {
	numerator := 0.0
	denominator := 0.0

	for value, membership := range aggregatedMemberships {
		numerator += value * membership
		denominator += membership
	}

	if denominator == 0 {
		// No rule contributed to the output surface, so the safest deterministic fallback is zero.
		return 0
	}

	return numerator / denominator
}

func (defuzzifier *Defuzzifier) DominantSet(outputVariable domain.Variable, crispValue float64) string {
	dominantSet := ""
	dominantDegree := 0.0

	for _, set := range outputVariable.Sets {
		degree := set.Degree(crispValue)
		if degree > dominantDegree {
			dominantDegree = degree
			dominantSet = set.Key
		}
	}

	return dominantSet
}
