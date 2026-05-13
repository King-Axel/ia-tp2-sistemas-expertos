package store

import (
	"fmt"
	"sync"

	"iteraciones/backend/internal/sebr/domain"
)

type RuleStore struct {
	mutex  sync.RWMutex
	rules  []domain.Rule
	nextID int
}

func NewRuleStore(defaultRules []domain.Rule) *RuleStore {
	rules := append([]domain.Rule(nil), defaultRules...)

	return &RuleStore{
		rules:  rules,
		nextID: len(rules) + 1,
	}
}

func (store *RuleStore) All() []domain.Rule {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	return append([]domain.Rule(nil), store.rules...)
}

func (store *RuleStore) Append(rule domain.Rule) domain.Rule {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	rule.ID = fmt.Sprintf("R%02d", store.nextID)
	store.nextID++
	store.rules = append(store.rules, rule)

	return rule
}
