package application

import "testing"

func TestIrrigationServiceInferIrrigation(t *testing.T) {
	service := NewIrrigationService()

	result, err := service.InferIrrigation(IrrigationInput{
		SoilHumidity:       42,
		AmbientTemperature: 28,
		TimeOfDay:          870,
	})
	if err != nil {
		t.Fatalf("InferIrrigation returned error: %v", err)
	}

	if len(result.InputMemberships) == 0 {
		t.Fatal("expected input memberships")
	}

	if len(result.ActivatedRules) == 0 {
		t.Fatal("expected activated rules")
	}

	if result.CrispOutputValue <= 0 {
		t.Fatalf("expected positive crisp output, got %.2f", result.CrispOutputValue)
	}

	t.Logf("crisp valve opening: %.2f", result.CrispOutputValue)
	for _, rule := range result.ActivatedRules {
		t.Logf("%s activated at %.2f -> %s", rule.RuleID, rule.ActivationDegree, rule.OutputSetKey)
	}
}
