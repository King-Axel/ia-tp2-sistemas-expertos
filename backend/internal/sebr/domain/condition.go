package domain

type FactLookup interface {
	Get(key string) (FactValue, bool)
}

type ConditionKind string

const (
	FactCondition ConditionKind = "fact"
	NotCondition  ConditionKind = "not"
	AndCondition  ConditionKind = "and"
	OrCondition   ConditionKind = "or"
)

type ComparisonOperator string

const (
	EqualsOperator             ComparisonOperator = "=="
	NotEqualsOperator          ComparisonOperator = "!="
	GreaterThanOperator        ComparisonOperator = ">"
	GreaterThanOrEqualOperator ComparisonOperator = ">="
	LessThanOperator           ComparisonOperator = "<"
	LessThanOrEqualOperator    ComparisonOperator = "<="
)

type Condition struct {
	Kind     ConditionKind
	FactKey  string
	Operator ComparisonOperator
	Value    FactValue
	Children []Condition
}

func BoolFact(key string, expected bool) Condition {
	return Condition{
		Kind:     FactCondition,
		FactKey:  key,
		Operator: EqualsOperator,
		Value:    NewBoolValue(expected),
	}
}

func NumberComparison(key string, operator ComparisonOperator, expected float64) Condition {
	return Condition{
		Kind:     FactCondition,
		FactKey:  key,
		Operator: operator,
		Value:    NewNumberValue(expected),
	}
}

func Not(condition Condition) Condition {
	return Condition{
		Kind:     NotCondition,
		Children: []Condition{condition},
	}
}

func All(conditions ...Condition) Condition {
	return Condition{
		Kind:     AndCondition,
		Children: conditions,
	}
}

func Any(conditions ...Condition) Condition {
	return Condition{
		Kind:     OrCondition,
		Children: conditions,
	}
}

func (condition Condition) Matches(facts FactLookup) bool {
	switch condition.Kind {
	case FactCondition:
		return condition.matchesFact(facts)
	case NotCondition:
		if len(condition.Children) != 1 {
			return false
		}
		return !condition.Children[0].Matches(facts)
	case AndCondition:
		for _, child := range condition.Children {
			if !child.Matches(facts) {
				return false
			}
		}
		return true
	case OrCondition:
		for _, child := range condition.Children {
			if child.Matches(facts) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func (condition Condition) matchesFact(facts FactLookup) bool {
	actual, exists := facts.Get(condition.FactKey)
	if !exists {
		return false
	}

	if actual.Kind() != condition.Value.Kind() {
		return false
	}

	switch actual.Kind() {
	case BoolFactValue:
		if condition.Operator == EqualsOperator {
			return actual.Equal(condition.Value)
		}
		if condition.Operator == NotEqualsOperator {
			return !actual.Equal(condition.Value)
		}
		return false
	case NumberFactValue:
		return compareNumbers(actual, condition.Operator, condition.Value)
	default:
		return false
	}
}

func compareNumbers(actual FactValue, operator ComparisonOperator, expected FactValue) bool {
	actualNumber, ok := actual.Number()
	if !ok {
		return false
	}

	expectedNumber, ok := expected.Number()
	if !ok {
		return false
	}

	switch operator {
	case EqualsOperator:
		return actualNumber == expectedNumber
	case NotEqualsOperator:
		return actualNumber != expectedNumber
	case GreaterThanOperator:
		return actualNumber > expectedNumber
	case GreaterThanOrEqualOperator:
		return actualNumber >= expectedNumber
	case LessThanOperator:
		return actualNumber < expectedNumber
	case LessThanOrEqualOperator:
		return actualNumber <= expectedNumber
	default:
		return false
	}
}
