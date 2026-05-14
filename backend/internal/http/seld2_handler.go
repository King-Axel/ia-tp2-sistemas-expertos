package httpserver

import (
	"errors"
	"net/http"

	"iteraciones/backend/internal/seld/application"
	"iteraciones/backend/internal/seld/domain"

	"github.com/gin-gonic/gin"
)

type SELD2Handler struct {
	ventilationService *application.VentilationService
}

func NewSELD2Handler(ventilationService *application.VentilationService) *SELD2Handler {
	return &SELD2Handler{
		ventilationService: ventilationService,
	}
}

func (handler *SELD2Handler) ListRules(context *gin.Context) {
	context.JSON(http.StatusOK, SELD2RulesResponse{
		Rules: mapSELD2RulesResponse(handler.ventilationService.Rules()),
	})
}

func (handler *SELD2Handler) Infer(context *gin.Context) {
	var request SELD2InferenceRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := validateSELD2InferenceRequest(request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := handler.ventilationService.InferVentilation(application.VentilationInput{
		PlantTemperature: *request.PlantTemperature,
		GasConcentration: *request.GasConcentration,
	})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, mapSELD2InferenceResponse(result.InferenceResult))
}

func validateSELD2InferenceRequest(request SELD2InferenceRequest) error {
	if !isInRange(*request.PlantTemperature, 0, 60) {
		return errors.New("plant_temperature must be between 0 and 60")
	}

	if !isInRange(*request.GasConcentration, 0, 100) {
		return errors.New("gas_concentration must be between 0 and 100")
	}

	return nil
}

func mapSELD2RulesResponse(rules []domain.Rule) []SELD2RuleResponse {
	response := make([]SELD2RuleResponse, 0, len(rules))
	for _, rule := range rules {
		response = append(response, SELD2RuleResponse{
			ID:         rule.ID,
			Name:       rule.Name,
			Expression: rule.Expression,
		})
	}

	return response
}

func mapSELD2InferenceResponse(result domain.InferenceResult) SELD2InferenceResponse {
	return SELD2InferenceResponse{
		CrispOutputValue:  result.CrispOutputValue,
		DominantOutputSet: result.DominantOutputSet,
		InputMemberships:  mapSELD1InputMembershipsResponse(result.InputMemberships),
		ActivatedRules:    mapSELD1ActivatedRulesResponse(result.ActivatedRules),
		AggregatedOutput:  mapSELD1AggregatedOutputResponse(result.AggregatedOutput),
		OutputVariableKey: result.OutputVariableKey,
	}
}
