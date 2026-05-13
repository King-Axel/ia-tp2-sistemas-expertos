package engine_test

import (
	"testing"

	"iteraciones/backend/internal/sebr/domain"
	"iteraciones/backend/internal/sebr/engine"
	"iteraciones/backend/internal/sebr/knowledge"
)

func TestInferenceEngineDerivesCOVIDCompatibility(t *testing.T) {
	factBase := engine.NewFactBase(
		domain.NewNumberFact("temperatura", 39),
		domain.NewBoolFact("dolor_cabeza", true),
		domain.NewBoolFact("tos", true),
		domain.NewBoolFact("perdida_gusto_u_olfato", true),
	)

	inferenceEngine := engine.NewInferenceEngine(knowledge.Rules(), knowledge.FinalConclusionKeys())
	result := inferenceEngine.Run(factBase)

	if !hasFact(result.FinalConclusions, "compatible_con_covid") {
		t.Fatal("expected compatible_con_covid conclusion")
	}

	if hasFact(result.FinalConclusions, "no_puede_concluir") {
		t.Fatal("did not expect no_puede_concluir when COVID compatibility was derived")
	}

	if len(result.ActivatedRules) == 0 {
		t.Fatal("expected activated rules trace")
	}
}

func TestInferenceEngineDerivesNoConclusionForLowTemperature(t *testing.T) {
	factBase := engine.NewFactBase(
		domain.NewNumberFact("temperatura", 37.5),
	)

	inferenceEngine := engine.NewInferenceEngine(knowledge.Rules(), knowledge.FinalConclusionKeys())
	result := inferenceEngine.Run(factBase)

	if !hasFact(result.FinalConclusions, "no_puede_concluir") {
		t.Fatal("expected no_puede_concluir conclusion")
	}
}

func hasFact(facts []domain.Fact, key string) bool {
	for _, fact := range facts {
		if fact.Key == key {
			return true
		}
	}

	return false
}
