package engine

import "iteraciones/backend/internal/seld/domain"

type InferenceEngine struct {
	rules []domain.Rule
}

func NewInferenceEngine(rules []domain.Rule) *InferenceEngine {
	return &InferenceEngine{
		rules: append([]domain.Rule(nil), rules...),
	}
}

func (engine *InferenceEngine) Evaluate(
	memberships domain.VariableMemberships,
) ([]domain.ActivatedRule, []OutputActivation) {
	activatedRules := make([]domain.ActivatedRule, 0)
	outputActivations := make([]OutputActivation, 0)

	for _, rule := range engine.rules {
		activationDegree := engine.evaluateRule(rule, memberships)
		if activationDegree <= 0 {
			continue
		}

		activatedRules = append(activatedRules, domain.ActivatedRule{
			RuleID:            rule.ID,
			RuleName:          rule.Name,
			Expression:        rule.Expression,
			ActivationDegree:  activationDegree,
			OutputVariableKey: rule.Consequent.VariableKey,
			OutputSetKey:      rule.Consequent.SetKey,
		})
		outputActivations = append(outputActivations, OutputActivation{
			SetKey: rule.Consequent.SetKey,
			Degree: activationDegree,
		})
	}

	return activatedRules, outputActivations
}

func (engine *InferenceEngine) evaluateRule(
	rule domain.Rule,
	memberships domain.VariableMemberships,
) float64 {
	if len(rule.Antecedents) == 0 {
		return 0
	}

	activationDegree := 1.0

	for _, antecedent := range rule.Antecedents {
		degree := membershipDegree(memberships, antecedent.VariableKey, antecedent.SetKey)
		if antecedent.Negated {
			degree = 1 - degree
		}

		activationDegree = minFloat(activationDegree, degree)
	}

	return activationDegree
}

func membershipDegree(
	memberships domain.VariableMemberships,
	variableKey string,
	setKey string,
) float64 {
	for _, membership := range memberships[variableKey] {
		if membership.SetKey == setKey {
			return membership.Degree
		}
	}

	return 0
}

func minFloat(left, right float64) float64 {
	if left < right {
		return left
	}

	return right
}
