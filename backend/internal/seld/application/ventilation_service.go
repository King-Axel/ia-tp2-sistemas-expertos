package application

import (
	"iteraciones/backend/internal/seld/domain"
	"iteraciones/backend/internal/seld/engine"
	"iteraciones/backend/internal/seld/knowledge"
)

type VentilationInput struct {
	PlantTemperature float64
	GasConcentration float64
}

type VentilationResult struct {
	domain.InferenceResult
}

type VentilationService struct {
	inputVariables []domain.Variable
	outputVariable domain.Variable
	rules          []domain.Rule
	step           float64
}

func NewVentilationService() *VentilationService {
	return &VentilationService{
		inputVariables: knowledge.VentilationInputVariables(),
		outputVariable: knowledge.VentilationOutputVariable(),
		rules:          knowledge.VentilationRules(),
		step:           0.5,
	}
}

func (service *VentilationService) Rules() []domain.Rule {
	return append([]domain.Rule(nil), service.rules...)
}

func (service *VentilationService) InferVentilation(input VentilationInput) (VentilationResult, error) {
	crispInputs := service.clampedInputs(input)

	fuzzifier := engine.NewFuzzifier(service.inputVariables)
	inputMemberships := fuzzifier.Fuzzify(crispInputs)

	inferenceEngine := engine.NewInferenceEngine(service.rules)
	activatedRules, outputActivations := inferenceEngine.Evaluate(inputMemberships)

	aggregator := engine.NewAggregator(service.step)
	aggregatedOutput := aggregator.Aggregate(service.outputVariable, outputActivations)

	defuzzifier := engine.NewDefuzzifier()
	crispOutput := defuzzifier.Centroid(aggregatedOutput)
	dominantSet := defuzzifier.DominantSet(service.outputVariable, crispOutput)

	return VentilationResult{
		InferenceResult: domain.InferenceResult{
			CrispOutputValue:  crispOutput,
			DominantOutputSet: dominantSet,
			InputMemberships:  inputMemberships,
			ActivatedRules:    activatedRules,
			AggregatedOutput:  aggregatedOutput,
			OutputVariableKey: knowledge.MotorSpeedKey,
		},
	}, nil
}

func (service *VentilationService) clampedInputs(input VentilationInput) map[string]float64 {
	values := map[string]float64{
		knowledge.PlantTemperatureKey: input.PlantTemperature,
		knowledge.GasConcentrationKey: input.GasConcentration,
	}

	for _, variable := range service.inputVariables {
		values[variable.Key] = variable.Universe.Clamp(values[variable.Key])
	}

	return values
}
