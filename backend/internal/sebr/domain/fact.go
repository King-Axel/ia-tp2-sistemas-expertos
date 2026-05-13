package domain

import "fmt"

type FactValueKind string

const (
	BoolFactValue   FactValueKind = "bool"
	NumberFactValue FactValueKind = "number"
)

type FactValue struct {
	kind   FactValueKind
	bool   bool
	number float64
}

func NewBoolValue(value bool) FactValue {
	return FactValue{kind: BoolFactValue, bool: value}
}

func NewNumberValue(value float64) FactValue {
	return FactValue{kind: NumberFactValue, number: value}
}

func (value FactValue) Kind() FactValueKind {
	return value.kind
}

func (value FactValue) Bool() (bool, bool) {
	return value.bool, value.kind == BoolFactValue
}

func (value FactValue) Number() (float64, bool) {
	return value.number, value.kind == NumberFactValue
}

func (value FactValue) Equal(other FactValue) bool {
	if value.kind != other.kind {
		return false
	}

	switch value.kind {
	case BoolFactValue:
		return value.bool == other.bool
	case NumberFactValue:
		return value.number == other.number
	default:
		return false
	}
}

func (value FactValue) String() string {
	switch value.kind {
	case BoolFactValue:
		return fmt.Sprintf("%t", value.bool)
	case NumberFactValue:
		return fmt.Sprintf("%.2f", value.number)
	default:
		return "<unknown>"
	}
}

type Fact struct {
	Key   string
	Value FactValue
}

func NewBoolFact(key string, value bool) Fact {
	return Fact{Key: key, Value: NewBoolValue(value)}
}

func NewNumberFact(key string, value float64) Fact {
	return Fact{Key: key, Value: NewNumberValue(value)}
}
