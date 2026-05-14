package httpserver

import (
	"errors"
	"net/http"
	"sort"

	"iteraciones/backend/internal/seld/application"
	"iteraciones/backend/internal/seld/domain"

	"github.com/gin-gonic/gin"
)

type SELD1Handler struct {
	irrigationService *application.IrrigationService
}

func NewSELD1Handler(irrigationService *application.IrrigationService) *SELD1Handler {
	return &SELD1Handler{
		irrigationService: irrigationService,
	}
}

func (handler *SELD1Handler) ListRules(context *gin.Context) {
	context.JSON(http.StatusOK, SELD1RulesResponse{
		Rules: mapSELD1RulesResponse(handler.irrigationService.Rules()),
	})
}

func (handler *SELD1Handler) Infer(context *gin.Context) {
	var request SELD1InferenceRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := validateSELD1InferenceRequest(request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := handler.irrigationService.InferIrrigation(application.IrrigationInput{
		SoilHumidity:       *request.SoilHumidity,
		AmbientTemperature: *request.AmbientTemperature,
		TimeOfDay:          *request.TimeOfDay,
	})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, mapSELD1InferenceResponse(result.InferenceResult))
}

func validateSELD1InferenceRequest(request SELD1InferenceRequest) error {
	if !isInRange(*request.SoilHumidity, 0, 100) {
		return errors.New("soil_humidity must be between 0 and 100")
	}

	if !isInRange(*request.AmbientTemperature, 0, 45) {
		return errors.New("ambient_temperature must be between 0 and 45")
	}

	if !isInRange(*request.TimeOfDay, 0, 1439) {
		return errors.New("time_of_day must be between 0 and 1439")
	}

	return nil
}

func isInRange(value, minimum, maximum float64) bool {
	return value >= minimum && value <= maximum
}

func mapSELD1RulesResponse(rules []domain.Rule) []SELD1RuleResponse {
	response := make([]SELD1RuleResponse, 0, len(rules))
	for _, rule := range rules {
		response = append(response, SELD1RuleResponse{
			ID:         rule.ID,
			Name:       rule.Name,
			Expression: rule.Expression,
		})
	}

	return response
}

func mapSELD1InferenceResponse(result domain.InferenceResult) SELD1InferenceResponse {
	return SELD1InferenceResponse{
		CrispOutputValue:  result.CrispOutputValue,
		DominantOutputSet: result.DominantOutputSet,
		InputMemberships:  mapSELD1InputMembershipsResponse(result.InputMemberships),
		ActivatedRules:    mapSELD1ActivatedRulesResponse(result.ActivatedRules),
		AggregatedOutput:  mapSELD1AggregatedOutputResponse(result.AggregatedOutput),
		OutputVariableKey: result.OutputVariableKey,
	}
}

func mapSELD1InputMembershipsResponse(
	memberships domain.VariableMemberships,
) map[string][]SetMembershipDTO {
	response := make(map[string][]SetMembershipDTO, len(memberships))

	for variableKey, setMemberships := range memberships {
		response[variableKey] = make([]SetMembershipDTO, 0, len(setMemberships))
		for _, membership := range setMemberships {
			response[variableKey] = append(response[variableKey], SetMembershipDTO{
				SetKey: membership.SetKey,
				Degree: membership.Degree,
			})
		}
	}

	return response
}

func mapSELD1ActivatedRulesResponse(
	activatedRules []domain.ActivatedRule,
) []ActivatedRuleDTO {
	response := make([]ActivatedRuleDTO, 0, len(activatedRules))

	for _, rule := range activatedRules {
		response = append(response, ActivatedRuleDTO{
			RuleID:            rule.RuleID,
			RuleName:          rule.RuleName,
			Expression:        rule.Expression,
			ActivationDegree:  rule.ActivationDegree,
			OutputVariableKey: rule.OutputVariableKey,
			OutputSetKey:      rule.OutputSetKey,
		})
	}

	return response
}

func mapSELD1AggregatedOutputResponse(aggregatedOutput map[float64]float64) []AggregatedOutputPoint {
	values := make([]float64, 0, len(aggregatedOutput))
	for value := range aggregatedOutput {
		values = append(values, value)
	}
	sort.Float64s(values)

	response := make([]AggregatedOutputPoint, 0, len(values))
	for _, value := range values {
		response = append(response, AggregatedOutputPoint{
			X:          value,
			Membership: aggregatedOutput[value],
		})
	}

	return response
}
