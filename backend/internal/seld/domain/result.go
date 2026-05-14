package domain

type SetMembership struct {
	SetKey string
	Degree float64
}

type VariableMemberships map[string][]SetMembership

type ActivatedRule struct {
	RuleID            string
	RuleName          string
	Expression        string
	ActivationDegree  float64
	OutputVariableKey string
	OutputSetKey      string
}

type InferenceResult struct {
	CrispOutputValue  float64
	DominantOutputSet string
	InputMemberships  VariableMemberships
	ActivatedRules    []ActivatedRule
	AggregatedOutput  map[float64]float64
	OutputVariableKey string
}
