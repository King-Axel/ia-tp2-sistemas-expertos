package engine

import "iteraciones/backend/internal/seld/domain"

type OutputActivation struct {
	SetKey string
	Degree float64
}

type Aggregator struct {
	step float64
}

func NewAggregator(step float64) *Aggregator {
	return &Aggregator{
		step: step,
	}
}

func (aggregator *Aggregator) Aggregate(
	outputVariable domain.Variable,
	activations []OutputActivation,
) map[float64]float64 {
	aggregated := make(map[float64]float64)

	for x := outputVariable.Universe.Min; x <= outputVariable.Universe.Max; x += aggregator.step {
		aggregated[x] = aggregator.aggregatedDegreeAt(outputVariable, activations, x)
	}

	return aggregated
}

func (aggregator *Aggregator) aggregatedDegreeAt(
	outputVariable domain.Variable,
	activations []OutputActivation,
	value float64,
) float64 {
	maxDegree := 0.0

	for _, activation := range activations {
		set, exists := outputVariable.SetByKey(activation.SetKey)
		if !exists {
			continue
		}

		clippedDegree := minFloat(set.Degree(value), activation.Degree)
		if clippedDegree > maxDegree {
			maxDegree = clippedDegree
		}
	}

	return maxDegree
}
