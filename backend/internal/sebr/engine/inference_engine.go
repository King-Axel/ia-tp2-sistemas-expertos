package engine

import (
	"sort"

	"iteraciones/backend/internal/sebr/domain"
)

type InferenceEngine struct {
	rules          []domain.Rule
	conclusionKeys map[string]struct{}
	maxCycles      int
}

func NewInferenceEngine(rules []domain.Rule, conclusionKeys []string) *InferenceEngine {
	orderedRules := append([]domain.Rule(nil), rules...)
	sort.SliceStable(orderedRules, func(left, right int) bool {
		return orderedRules[left].Priority > orderedRules[right].Priority
	})

	keys := make(map[string]struct{}, len(conclusionKeys))
	for _, key := range conclusionKeys {
		keys[key] = struct{}{}
	}

	return &InferenceEngine{
		rules:          orderedRules,
		conclusionKeys: keys,
		maxCycles:      100,
	}
}

func (engine *InferenceEngine) Run(factBase *FactBase) domain.InferenceResult {
	firedRules := make(map[string]struct{})
	activatedRules := make([]domain.ActivatedRule, 0)

	for cycle := 1; cycle <= engine.maxCycles; cycle++ {
		matchingRules := engine.nextMatchingRules(factBase, firedRules)
		if len(matchingRules) == 0 {
			break
		}

		addedFact := false
		for _, rule := range matchingRules {
			firedRules[rule.ID] = struct{}{}
			if factBase.Add(rule.Conclusion) {
				addedFact = true
			}

			activatedRules = append(activatedRules, domain.ActivatedRule{
				RuleID:     rule.ID,
				RuleName:   rule.Name,
				Conclusion: rule.Conclusion,
				Expression: rule.Expression,
				Cycle:      cycle,
			})
		}

		if !addedFact {
			break
		}
	}

	return domain.InferenceResult{
		Facts:            factBase.Facts(),
		ActivatedRules:   activatedRules,
		FinalConclusions: engine.finalConclusions(factBase),
	}
}

func (engine *InferenceEngine) nextMatchingRules(factBase *FactBase, firedRules map[string]struct{}) []domain.Rule {
	matchingRules := make([]domain.Rule, 0)
	activePriority := 0
	hasActivePriority := false

	for _, rule := range engine.rules {
		if _, fired := firedRules[rule.ID]; fired {
			continue
		}

		if hasActivePriority && rule.Priority < activePriority {
			break
		}

		if rule.Matches(factBase) {
			if !hasActivePriority {
				activePriority = rule.Priority
				hasActivePriority = true
			}
			matchingRules = append(matchingRules, rule)
		}
	}

	return matchingRules
}

func (engine *InferenceEngine) finalConclusions(factBase *FactBase) []domain.Fact {
	conclusions := make([]domain.Fact, 0, len(engine.conclusionKeys))
	for _, fact := range factBase.Facts() {
		if _, isConclusion := engine.conclusionKeys[fact.Key]; isConclusion {
			conclusions = append(conclusions, fact)
		}
	}

	return conclusions
}
