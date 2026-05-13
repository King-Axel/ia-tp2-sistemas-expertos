package engine

import (
	"sort"

	"iteraciones/backend/internal/sebr/domain"
)

type FactBase struct {
	facts map[string]domain.FactValue
}

func NewFactBase(initialFacts ...domain.Fact) *FactBase {
	base := &FactBase{
		facts: make(map[string]domain.FactValue),
	}

	for _, fact := range initialFacts {
		base.Add(fact)
	}

	return base
}

func (base *FactBase) Add(fact domain.Fact) bool {
	current, exists := base.facts[fact.Key]
	if exists && current.Equal(fact.Value) {
		return false
	}

	base.facts[fact.Key] = fact.Value
	return true
}

func (base *FactBase) Get(key string) (domain.FactValue, bool) {
	value, exists := base.facts[key]
	return value, exists
}

func (base *FactBase) Has(key string) bool {
	_, exists := base.facts[key]
	return exists
}

func (base *FactBase) Facts() []domain.Fact {
	keys := make([]string, 0, len(base.facts))
	for key := range base.facts {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	facts := make([]domain.Fact, 0, len(keys))
	for _, key := range keys {
		facts = append(facts, domain.Fact{
			Key:   key,
			Value: base.facts[key],
		})
	}

	return facts
}
