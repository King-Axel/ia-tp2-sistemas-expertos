package service

import (
	"fmt"
	"strings"

	"iteraciones/backend/internal/sebr/domain"
	"iteraciones/backend/internal/sebr/parser"
	"iteraciones/backend/internal/sebr/store"
)

type RuleService struct {
	store  *store.RuleStore
	parser *parser.ExpressionParser
}

func NewRuleService(store *store.RuleStore) *RuleService {
	return &RuleService{
		store:  store,
		parser: parser.NewExpressionParser(),
	}
}

func (service *RuleService) ListRules() []domain.Rule {
	return service.store.All()
}

func (service *RuleService) CreateRule(name string, expression string) (domain.Rule, error) {
	name = strings.TrimSpace(name)
	expression = strings.TrimSpace(expression)

	if name == "" {
		return domain.Rule{}, fmt.Errorf("missing name")
	}
	if expression == "" {
		return domain.Rule{}, fmt.Errorf("missing expression")
	}

	rule, err := service.parser.ParseRule("", name, expression)
	if err != nil {
		return domain.Rule{}, err
	}

	return service.store.Append(rule), nil
}
