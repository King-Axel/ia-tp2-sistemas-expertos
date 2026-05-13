package domain

type Rule struct {
	ID         string
	Name       string
	Condition  Condition
	Conclusion Fact
	Priority   int
	Expression string
}

func (rule Rule) Matches(facts FactLookup) bool {
	return rule.Condition.Matches(facts)
}
