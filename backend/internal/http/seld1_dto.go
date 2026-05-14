package httpserver

type SELD1RuleResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

type SELD1RulesResponse struct {
	Rules []SELD1RuleResponse `json:"rules"`
}

type SELD1InferenceRequest struct {
	SoilHumidity       *float64 `json:"soil_humidity" binding:"required"`
	AmbientTemperature *float64 `json:"ambient_temperature" binding:"required"`
	TimeOfDay          *float64 `json:"time_of_day" binding:"required"`
}

type SELD1InferenceResponse struct {
	CrispOutputValue  float64                       `json:"crisp_output_value"`
	DominantOutputSet string                        `json:"dominant_output_set"`
	InputMemberships  map[string][]SetMembershipDTO `json:"input_memberships"`
	ActivatedRules    []SELD1ActivatedRuleResponse  `json:"activated_rules"`
	AggregatedOutput  []AggregatedOutputPoint       `json:"aggregated_output"`
	OutputVariableKey string                        `json:"output_variable_key"`
}

type SetMembershipDTO struct {
	SetKey string  `json:"set_key"`
	Degree float64 `json:"degree"`
}

type SELD1ActivatedRuleResponse struct {
	RuleID            string  `json:"rule_id"`
	RuleName          string  `json:"rule_name"`
	Expression        string  `json:"expression"`
	ActivationDegree  float64 `json:"activation_degree"`
	OutputVariableKey string  `json:"output_variable_key"`
	OutputSetKey      string  `json:"output_set_key"`
}

type AggregatedOutputPoint struct {
	X          float64 `json:"x"`
	Membership float64 `json:"membership"`
}
