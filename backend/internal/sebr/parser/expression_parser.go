package parser

import (
	"fmt"
	"strconv"
	"strings"

	"iteraciones/backend/internal/sebr/domain"
)

type ExpressionParser struct{}

func NewExpressionParser() *ExpressionParser {
	return &ExpressionParser{}
}

func (parser *ExpressionParser) ParseRule(id string, name string, expression string) (domain.Rule, error) {
	conditionsText, conclusionText, err := splitExpression(expression)
	if err != nil {
		return domain.Rule{}, err
	}

	condition, err := parseConditions(conditionsText)
	if err != nil {
		return domain.Rule{}, err
	}

	conclusion := normalizeFactKey(conclusionText)
	if conclusion == "" {
		return domain.Rule{}, fmt.Errorf("empty conclusion")
	}

	return domain.Rule{
		ID:         id,
		Name:       name,
		Condition:  condition,
		Conclusion: domain.NewBoolFact(conclusion, true),
		Priority:   10,
		Expression: expression,
	}, nil
}

func splitExpression(expression string) (string, string, error) {
	trimmed := strings.TrimSpace(expression)
	if trimmed == "" {
		return "", "", fmt.Errorf("missing expression")
	}
	if strings.ContainsAny(trimmed, "()") {
		return "", "", fmt.Errorf("parentheses are not supported")
	}

	tokens := strings.Fields(trimmed)
	if len(tokens) < 4 {
		return "", "", fmt.Errorf("expression must use IF <conditions> THEN <conclusion>")
	}
	if !equalKeyword(tokens[0], "IF") {
		return "", "", fmt.Errorf("expression must start with IF")
	}

	thenIndex := -1
	for index, token := range tokens {
		if equalKeyword(token, "THEN") {
			if thenIndex != -1 {
				return "", "", fmt.Errorf("expression must contain only one THEN")
			}
			thenIndex = index
		}
	}
	if thenIndex == -1 {
		return "", "", fmt.Errorf("expression must contain THEN")
	}
	if thenIndex == 1 {
		return "", "", fmt.Errorf("missing conditions")
	}
	if thenIndex == len(tokens)-1 {
		return "", "", fmt.Errorf("empty conclusion")
	}
	if len(tokens[thenIndex+1:]) != 1 {
		return "", "", fmt.Errorf("conclusion must be a single fact")
	}

	return strings.Join(tokens[1:thenIndex], " "), tokens[thenIndex+1], nil
}

func parseConditions(conditionsText string) (domain.Condition, error) {
	tokens := strings.Fields(conditionsText)
	if len(tokens) == 0 {
		return domain.Condition{}, fmt.Errorf("missing conditions")
	}

	operator := ""
	for _, token := range tokens {
		if equalKeyword(token, "AND") || equalKeyword(token, "OR") {
			current := strings.ToUpper(token)
			if operator != "" && operator != current {
				return domain.Condition{}, fmt.Errorf("mixed AND/OR expressions are not supported without parentheses")
			}
			operator = current
		}
	}

	parts, err := splitConditionParts(tokens, operator)
	if err != nil {
		return domain.Condition{}, err
	}

	conditions := make([]domain.Condition, 0, len(parts))
	for _, part := range parts {
		condition, err := parseConditionPart(part)
		if err != nil {
			return domain.Condition{}, err
		}
		conditions = append(conditions, condition)
	}

	if len(conditions) == 1 {
		return conditions[0], nil
	}
	if operator == "OR" {
		return domain.Any(conditions...), nil
	}
	return domain.All(conditions...), nil
}

func splitConditionParts(tokens []string, operator string) ([][]string, error) {
	parts := make([][]string, 0)
	current := make([]string, 0)

	for _, token := range tokens {
		if equalKeyword(token, "AND") || equalKeyword(token, "OR") {
			if operator == "" {
				return nil, fmt.Errorf("invalid logical operator")
			}
			if len(current) == 0 {
				return nil, fmt.Errorf("empty condition near %s", token)
			}
			parts = append(parts, current)
			current = make([]string, 0)
			continue
		}
		current = append(current, token)
	}

	if len(current) == 0 {
		return nil, fmt.Errorf("expression cannot end with a logical operator")
	}

	parts = append(parts, current)
	return parts, nil
}

func parseConditionPart(tokens []string) (domain.Condition, error) {
	if len(tokens) == 0 {
		return domain.Condition{}, fmt.Errorf("empty condition")
	}

	negated := false
	if equalKeyword(tokens[0], "NOT") {
		negated = true
		tokens = tokens[1:]
	}

	if len(tokens) == 0 {
		return domain.Condition{}, fmt.Errorf("NOT must be followed by a fact")
	}

	if len(tokens) == 1 {
		condition := domain.BoolFact(normalizeFactKey(tokens[0]), true)
		if negated {
			return domain.Not(condition), nil
		}
		return condition, nil
	}

	if negated {
		return domain.Condition{}, fmt.Errorf("NOT is only supported for standalone facts")
	}
	if len(tokens) != 3 {
		return domain.Condition{}, fmt.Errorf("condition must be a standalone fact or fact operator value")
	}

	key := normalizeFactKey(tokens[0])
	operator, err := parseOperator(tokens[1])
	if err != nil {
		return domain.Condition{}, err
	}

	value, err := parseValue(tokens[2])
	if err != nil {
		return domain.Condition{}, err
	}

	if value.Kind() == domain.BoolFactValue && operator != domain.EqualsOperator && operator != domain.NotEqualsOperator {
		return domain.Condition{}, fmt.Errorf("operator %s is not supported for boolean values", operator)
	}

	return domain.Condition{
		Kind:     domain.FactCondition,
		FactKey:  key,
		Operator: operator,
		Value:    value,
	}, nil
}

func parseOperator(token string) (domain.ComparisonOperator, error) {
	switch token {
	case "==":
		return domain.EqualsOperator, nil
	case "!=":
		return domain.NotEqualsOperator, nil
	case ">":
		return domain.GreaterThanOperator, nil
	case ">=":
		return domain.GreaterThanOrEqualOperator, nil
	case "<":
		return domain.LessThanOperator, nil
	case "<=":
		return domain.LessThanOrEqualOperator, nil
	default:
		return "", fmt.Errorf("unsupported operator %q", token)
	}
}

func parseValue(token string) (domain.FactValue, error) {
	if equalKeyword(token, "true") {
		return domain.NewBoolValue(true), nil
	}
	if equalKeyword(token, "false") {
		return domain.NewBoolValue(false), nil
	}

	number, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return domain.FactValue{}, fmt.Errorf("unsupported value %q", token)
	}

	return domain.NewNumberValue(number), nil
}

func equalKeyword(value string, keyword string) bool {
	return strings.EqualFold(value, keyword)
}

func normalizeFactKey(value string) string {
	key := strings.ToLower(strings.TrimSpace(value))
	aliases := map[string]string{
		"temperature":                 "temperatura",
		"headache":                    "dolor_cabeza",
		"cough":                       "tos",
		"fatigue":                     "fatiga",
		"loss_of_taste_or_smell":      "perdida_gusto_u_olfato",
		"breathing_difficulty":        "dificultad_respiratoria",
		"chest_pain":                  "dolor_toracico",
		"sore_throat":                 "dolor_garganta",
		"pain_behind_eyes":            "dolor_detras_de_ojos",
		"vomiting":                    "vomito",
		"muscle_or_joint_pain":        "dolor_muscular_o_articular",
		"bleeding":                    "sangrado",
		"close_contact":               "contacto_estrecho",
		"had_covid_before":            "tuvo_covid_antes",
		"fever":                       "fiebre",
		"suspected_fever_condition":   "sospecha_cuadro_febril",
		"fever_condition":             "cuadro_febril",
		"suspected_respiratory_case":  "sospecha_cuadro_respiratorio",
		"respiratory_case":            "cuadro_respiratorio",
		"evaluable_case":              "caso_evaluable",
		"has_evidence":                "hay_indicios",
		"covid_compatible":            "compatible_con_covid",
		"suspected_covid":             "sospecha_covid",
		"suspected_covid_reinfection": "sospecha_recontagio",
		"suspected_dengue":            "sospecha_dengue",
		"dengue_compatible":           "compatible_con_dengue",
		"cannot_conclude":             "no_puede_concluir",
	}

	if alias, exists := aliases[key]; exists {
		return alias
	}

	return key
}
