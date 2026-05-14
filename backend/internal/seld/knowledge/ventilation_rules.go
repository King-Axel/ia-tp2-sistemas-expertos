package knowledge

import (
	"fmt"

	"iteraciones/backend/internal/seld/domain"
)

func VentilationRules() []domain.Rule {
	return []domain.Rule{
		ventilationRule("R1", "", "dangerous", "maximum"),
		ventilationRule("R2", "critical", "", "maximum"),
		ventilationRule("R3", "low", "safe", "minimum"),
		ventilationRule("R4", "normal", "safe", "low"),
		ventilationRule("R5", "high", "safe", "medium"),
		ventilationRule("R6", "low", "moderate", "medium"),
		ventilationRule("R7", "normal", "moderate", "medium"),
		ventilationRule("R8", "high", "moderate", "high"),
	}
}

func ventilationRule(
	id string,
	temperatureSet string,
	gasSet string,
	outputSet string,
) domain.Rule {
	antecedents := make([]domain.Antecedent, 0, 2)

	if temperatureSet != "" {
		antecedents = append(antecedents, domain.Antecedent{
			VariableKey: PlantTemperatureKey,
			SetKey:      temperatureSet,
		})
	}

	if gasSet != "" {
		antecedents = append(antecedents, domain.Antecedent{
			VariableKey: GasConcentrationKey,
			SetKey:      gasSet,
		})
	}

	return domain.Rule{
		ID:          id,
		Name:        fmt.Sprintf("Ventilation rule %s", id),
		Antecedents: antecedents,
		Consequent: domain.Consequent{
			VariableKey: MotorSpeedKey,
			SetKey:      outputSet,
		},
		Expression: ventilationExpression(antecedents, outputSet),
	}
}

func ventilationExpression(antecedents []domain.Antecedent, outputSet string) string {
	expression := "IF "

	for index, antecedent := range antecedents {
		if index > 0 {
			expression += " AND "
		}

		expression += antecedent.VariableKey + " IS " + antecedent.SetKey
	}

	return expression + " THEN " + MotorSpeedKey + " IS " + outputSet
}
