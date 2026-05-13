package domain

type ActivatedRule struct {
	RuleID     string
	RuleName   string
	Conclusion Fact
	Expression string
	Cycle      int
}

type InferenceResult struct {
	Facts            []Fact
	ActivatedRules   []ActivatedRule
	FinalConclusions []Fact
}
