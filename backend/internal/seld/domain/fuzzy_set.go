package domain

type FuzzySet struct {
	Key        string
	Name       string
	Membership MembershipFunction
}

func (set FuzzySet) Degree(value float64) float64 {
	return set.Membership.Degree(value)
}
