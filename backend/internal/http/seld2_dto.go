package httpserver

type SELD2RuleResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

type SELD2RulesResponse struct {
	Rules []SELD2RuleResponse `json:"rules"`
}

type SELD2InferenceRequest struct {
	PlantTemperature *float64 `json:"plant_temperature" binding:"required"`
	GasConcentration *float64 `json:"gas_concentration" binding:"required"`
}

type SELD2InferenceResponse struct {
	CrispOutputValue  float64                       `json:"crisp_output_value"`
	DominantOutputSet string                        `json:"dominant_output_set"`
	InputMemberships  map[string][]SetMembershipDTO `json:"input_memberships"`
	ActivatedRules    []ActivatedRuleDTO            `json:"activated_rules"`
	AggregatedOutput  []AggregatedOutputPoint       `json:"aggregated_output"`
	OutputVariableKey string                        `json:"output_variable_key"`
}
