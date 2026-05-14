package engine

import "iteraciones/backend/internal/seld/domain"

type Fuzzifier struct {
	variables map[string]domain.Variable
}

func NewFuzzifier(variables []domain.Variable) *Fuzzifier {
	index := make(map[string]domain.Variable, len(variables))
	for _, variable := range variables {
		index[variable.Key] = variable
	}

	return &Fuzzifier{
		variables: index,
	}
}

func (fuzzifier *Fuzzifier) Fuzzify(inputs map[string]float64) domain.VariableMemberships {
	memberships := make(domain.VariableMemberships, len(inputs))

	for variableKey, value := range inputs {
		variable, exists := fuzzifier.variables[variableKey]
		if !exists {
			continue
		}

		clampedValue := variable.Universe.Clamp(value)
		setMemberships := make([]domain.SetMembership, 0, len(variable.Sets))

		for _, set := range variable.Sets {
			setMemberships = append(setMemberships, domain.SetMembership{
				SetKey: set.Key,
				Degree: set.Degree(clampedValue),
			})
		}

		memberships[variableKey] = setMemberships
	}

	return memberships
}
