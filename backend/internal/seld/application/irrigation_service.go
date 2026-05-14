package application

import (
	"iteraciones/backend/internal/seld/domain"
	"iteraciones/backend/internal/seld/engine"
	"iteraciones/backend/internal/seld/knowledge"
)

type IrrigationInput struct {
	SoilHumidity       float64
	AmbientTemperature float64
	TimeOfDay          float64
}

type IrrigationResult struct {
	domain.InferenceResult
}

type IrrigationService struct {
	inputVariables []domain.Variable
	outputVariable domain.Variable
	rules          []domain.Rule
	step           float64
}

func NewIrrigationService() *IrrigationService {
	return &IrrigationService{
		inputVariables: knowledge.IrrigationInputVariables(),
		outputVariable: knowledge.IrrigationOutputVariable(),
		rules:          knowledge.IrrigationRules(),
		step:           0.5,
	}
}

func (service *IrrigationService) Rules() []domain.Rule {
	return append([]domain.Rule(nil), service.rules...)
}

func (service *IrrigationService) InferIrrigation(input IrrigationInput) (IrrigationResult, error) {
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

	return IrrigationResult{
		InferenceResult: domain.InferenceResult{
			CrispOutputValue:  crispOutput,
			DominantOutputSet: dominantSet,
			InputMemberships:  inputMemberships,
			ActivatedRules:    activatedRules,
			AggregatedOutput:  aggregatedOutput,
			OutputVariableKey: knowledge.ValveOpeningKey,
		},
	}, nil
}

func (service *IrrigationService) clampedInputs(input IrrigationInput) map[string]float64 {
	values := map[string]float64{
		knowledge.SoilHumidityKey:       input.SoilHumidity,
		knowledge.AmbientTemperatureKey: input.AmbientTemperature,
		knowledge.TimeOfDayKey:          input.TimeOfDay,
	}

	for _, variable := range service.inputVariables {
		values[variable.Key] = variable.Universe.Clamp(values[variable.Key])
	}

	return values
}
