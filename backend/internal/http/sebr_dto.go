package httpserver

type SEBRRuleResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

type SEBRRulesResponse struct {
	Rules []SEBRRuleResponse `json:"rules"`
}

type CreateSEBRRuleRequest struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

type CreateSEBRRuleResponse struct {
	Rule SEBRRuleResponse `json:"rule"`
}

type SEBRInferenceRequest struct {
	Temperature         float64 `json:"temperature" binding:"required"`
	Headache            bool    `json:"headache"`
	Cough               bool    `json:"cough"`
	Fatigue             bool    `json:"fatigue"`
	LossOfTasteOrSmell  bool    `json:"loss_of_taste_or_smell"`
	BreathingDifficulty bool    `json:"breathing_difficulty"`
	ChestPain           bool    `json:"chest_pain"`
	SoreThroat          bool    `json:"sore_throat"`
	PainBehindEyes      bool    `json:"pain_behind_eyes"`
	Vomiting            bool    `json:"vomiting"`
	MuscleOrJointPain   bool    `json:"muscle_or_joint_pain"`
	Nausea              bool    `json:"nausea"`
	Bleeding            bool    `json:"bleeding"`
	CloseContact        bool    `json:"close_contact"`
	HadCovidBefore      bool    `json:"had_covid_before"`
}

type SEBRInferenceResponse struct {
	Result           string                  `json:"result"`
	ActivatedRules   []ActivatedRuleResponse `json:"activated_rules"`
	FinalConclusions []string                `json:"final_conclusions"`
	FinalFacts       map[string]any          `json:"final_facts"`
}

type ActivatedRuleResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Expression string `json:"expression"`
	Cycle      int    `json:"cycle"`
	Conclusion string `json:"conclusion"`
}
