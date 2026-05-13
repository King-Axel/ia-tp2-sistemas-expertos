package httpserver

import (
	"net/http"

	"iteraciones/backend/internal/sebr/domain"
	"iteraciones/backend/internal/sebr/engine"
	"iteraciones/backend/internal/sebr/knowledge"
	"iteraciones/backend/internal/sebr/service"

	"github.com/gin-gonic/gin"
)

type SEBRHandler struct {
	ruleService *service.RuleService
}

func NewSEBRHandler(ruleService *service.RuleService) *SEBRHandler {
	return &SEBRHandler{
		ruleService: ruleService,
	}
}

func (handler *SEBRHandler) ListRules(context *gin.Context) {
	context.JSON(http.StatusOK, SEBRRulesResponse{
		Rules: mapRulesResponse(handler.ruleService.ListRules()),
	})
}

func (handler *SEBRHandler) CreateRule(context *gin.Context) {
	var request CreateSEBRRuleRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	rule, err := handler.ruleService.CreateRule(request.Name, request.Expression)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, CreateSEBRRuleResponse{
		Rule: mapRuleResponse(rule),
	})
}

func (handler *SEBRHandler) Infer(context *gin.Context) {
	var request SEBRInferenceRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	factBase := engine.NewFactBase(mapInferenceRequestToFacts(request)...)
	inferenceEngine := engine.NewInferenceEngine(handler.ruleService.ListRules(), knowledge.FinalConclusionKeys())
	result := inferenceEngine.Run(factBase)

	context.JSON(http.StatusOK, mapInferenceResponse(result))
}

func mapRulesResponse(rules []domain.Rule) []SEBRRuleResponse {
	response := make([]SEBRRuleResponse, 0, len(rules))
	for _, rule := range rules {
		response = append(response, mapRuleResponse(rule))
	}

	return response
}

func mapRuleResponse(rule domain.Rule) SEBRRuleResponse {
	return SEBRRuleResponse{
		ID:         rule.ID,
		Name:       rule.Name,
		Expression: rule.Expression,
	}
}

func mapInferenceRequestToFacts(request SEBRInferenceRequest) []domain.Fact {
	return []domain.Fact{
		domain.NewNumberFact("temperatura", request.Temperature),
		domain.NewBoolFact("dolor_cabeza", request.Headache),
		domain.NewBoolFact("tos", request.Cough),
		domain.NewBoolFact("fatiga", request.Fatigue),
		domain.NewBoolFact("perdida_gusto_u_olfato", request.LossOfTasteOrSmell),
		domain.NewBoolFact("dificultad_respiratoria", request.BreathingDifficulty),
		domain.NewBoolFact("dolor_toracico", request.ChestPain),
		domain.NewBoolFact("dolor_garganta", request.SoreThroat),
		domain.NewBoolFact("dolor_detras_de_ojos", request.PainBehindEyes),
		domain.NewBoolFact("vomito", request.Vomiting),
		domain.NewBoolFact("dolor_muscular_o_articular", request.MuscleOrJointPain),
		domain.NewBoolFact("nausea", request.Nausea),
		domain.NewBoolFact("sangrado", request.Bleeding),
		domain.NewBoolFact("contacto_estrecho", request.CloseContact),
		domain.NewBoolFact("tuvo_covid_antes", request.HadCovidBefore),
	}
}

func mapInferenceResponse(result domain.InferenceResult) SEBRInferenceResponse {
	return SEBRInferenceResponse{
		Result:           selectAPIResult(result.FinalConclusions),
		ActivatedRules:   mapActivatedRulesResponse(result.ActivatedRules),
		FinalConclusions: mapFinalConclusionsResponse(result.FinalConclusions),
		FinalFacts:       mapFactsResponse(result.Facts),
	}
}

func mapActivatedRulesResponse(activatedRules []domain.ActivatedRule) []ActivatedRuleResponse {
	response := make([]ActivatedRuleResponse, 0, len(activatedRules))
	for _, activatedRule := range activatedRules {
		response = append(response, ActivatedRuleResponse{
			ID:         activatedRule.RuleID,
			Name:       activatedRule.RuleName,
			Expression: activatedRule.Expression,
			Cycle:      activatedRule.Cycle,
			Conclusion: activatedRule.Conclusion.Key,
		})
	}

	return response
}

func mapFinalConclusionsResponse(facts []domain.Fact) []string {
	response := make([]string, 0, len(facts))
	for _, fact := range facts {
		response = append(response, fact.Key)
	}

	return response
}

func mapFactsResponse(facts []domain.Fact) map[string]any {
	response := make(map[string]any, len(facts))
	for _, fact := range facts {
		if value, ok := fact.Value.Bool(); ok {
			response[fact.Key] = value
			continue
		}

		if value, ok := fact.Value.Number(); ok {
			response[fact.Key] = value
		}
	}

	return response
}

func selectAPIResult(finalConclusions []domain.Fact) string {
	conclusions := make(map[string]struct{}, len(finalConclusions))
	for _, fact := range finalConclusions {
		conclusions[fact.Key] = struct{}{}
	}

	// API response selection policy only; inference remains fully rule-driven.
	for _, key := range []string{"no_puede_concluir", "compatible_con_covid", "compatible_con_dengue"} {
		if _, exists := conclusions[key]; exists {
			return key
		}
	}

	return ""
}
