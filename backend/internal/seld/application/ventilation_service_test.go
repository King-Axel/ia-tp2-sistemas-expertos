package application

import "testing"

func TestVentilationServiceInferVentilationDangerousGas(t *testing.T) {
	service := NewVentilationService()

	result, err := service.InferVentilation(VentilationInput{
		PlantTemperature: 42,
		GasConcentration: 70,
	})
	if err != nil {
		t.Fatalf("InferVentilation returned error: %v", err)
	}

	if len(result.ActivatedRules) == 0 {
		t.Fatal("expected activated rules")
	}

	if result.OutputVariableKey != "motor_speed" {
		t.Fatalf("expected motor_speed output key, got %q", result.OutputVariableKey)
	}

	if result.CrispOutputValue <= 0 {
		t.Fatalf("expected positive crisp output, got %.2f", result.CrispOutputValue)
	}
}
