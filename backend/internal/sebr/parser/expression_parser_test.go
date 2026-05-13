package parser

import "testing"

func TestParseRuleWithAndExpression(t *testing.T) {
	parser := NewExpressionParser()

	rule, err := parser.ParseRule("R33", "Possible flu", "IF fever AND headache THEN possible_flu")
	if err != nil {
		t.Fatalf("expected valid expression: %v", err)
	}

	if rule.Conclusion.Key != "possible_flu" {
		t.Fatalf("expected possible_flu conclusion, got %q", rule.Conclusion.Key)
	}
}

func TestParseRuleRejectsMixedAndOrExpression(t *testing.T) {
	parser := NewExpressionParser()

	_, err := parser.ParseRule("R33", "Invalid", "IF fever AND headache OR cough THEN possible_flu")
	if err == nil {
		t.Fatal("expected mixed AND/OR expression to be rejected")
	}
}

func TestParseRuleMapsEnglishAliasesToInternalFactKeys(t *testing.T) {
	parser := NewExpressionParser()

	rule, err := parser.ParseRule("R33", "COVID", "IF temperature >= 38 THEN fever")
	if err != nil {
		t.Fatalf("expected valid expression: %v", err)
	}

	if rule.Condition.FactKey != "temperatura" {
		t.Fatalf("expected temperatura fact key, got %q", rule.Condition.FactKey)
	}
	if rule.Conclusion.Key != "fiebre" {
		t.Fatalf("expected fiebre conclusion, got %q", rule.Conclusion.Key)
	}
}
