package knowledge

import (
	"fmt"

	"iteraciones/backend/internal/seld/domain"
)

func IrrigationRules() []domain.Rule {
	return []domain.Rule{
		rule("R01", "elevated", "", "", false, "none"),
		rule("R02", "very_low", "very_cold", "", false, "low"),
		rule("R03", "very_low", "cold", "", false, "moderate"),
		rule("R04", "very_low", "mild", "", false, "high"),
		rule("R05", "very_low", "warm", "", false, "complete"),
		rule("R06", "very_low", "sweltering", "", false, "complete"),
		rule("R07", "adequate", "very_cold", "", false, "none"),
		rule("R08", "adequate", "cold", "", false, "low"),
		rule("R09", "adequate", "mild", "", false, "low"),
		rule("R10", "adequate", "warm", "", false, "moderate"),
		rule("R11", "adequate", "sweltering", "", false, "moderate"),
		rule("R12", "somewhat_high", "sweltering", "", true, "none"),
		rule("R13", "somewhat_high", "sweltering", "", false, "low"),
		rule("R14", "low", "sweltering", "", false, "complete"),
		rule("R15", "low", "very_cold", "dawn", false, "low"),
		rule("R16", "low", "very_cold", "morning", false, "moderate"),
		rule("R17", "low", "very_cold", "midday", false, "high"),
		rule("R18", "low", "very_cold", "afternoon", false, "moderate"),
		rule("R19", "low", "very_cold", "night", false, "low"),
		rule("R20", "low", "cold", "dawn", false, "low"),
		rule("R21", "low", "cold", "morning", false, "moderate"),
		rule("R22", "low", "cold", "midday", false, "high"),
		rule("R23", "low", "cold", "afternoon", false, "moderate"),
		rule("R24", "low", "cold", "night", false, "low"),
		rule("R25", "low", "mild", "dawn", false, "low"),
		rule("R26", "low", "mild", "morning", false, "moderate"),
		rule("R27", "low", "mild", "midday", false, "complete"),
		rule("R28", "low", "mild", "afternoon", false, "moderate"),
		rule("R29", "low", "mild", "night", false, "low"),
		rule("R30", "low", "warm", "dawn", false, "moderate"),
		rule("R31", "low", "warm", "morning", false, "high"),
		rule("R32", "low", "warm", "midday", false, "complete"),
		rule("R33", "low", "warm", "afternoon", false, "high"),
		rule("R34", "low", "warm", "night", false, "moderate"),
	}
}

func rule(
	id string,
	soilHumiditySet string,
	temperatureSet string,
	timeSet string,
	temperatureNegated bool,
	outputSet string,
) domain.Rule {
	antecedents := []domain.Antecedent{
		{VariableKey: SoilHumidityKey, SetKey: soilHumiditySet},
	}

	if temperatureSet != "" {
		antecedents = append(antecedents, domain.Antecedent{
			VariableKey: AmbientTemperatureKey,
			SetKey:      temperatureSet,
			Negated:     temperatureNegated,
		})
	}

	if timeSet != "" {
		antecedents = append(antecedents, domain.Antecedent{
			VariableKey: TimeOfDayKey,
			SetKey:      timeSet,
		})
	}

	return domain.Rule{
		ID:          id,
		Name:        fmt.Sprintf("Irrigation rule %s", id),
		Antecedents: antecedents,
		Consequent: domain.Consequent{
			VariableKey: ValveOpeningKey,
			SetKey:      outputSet,
		},
		Expression: expression(antecedents, outputSet),
	}
}

func expression(antecedents []domain.Antecedent, outputSet string) string {
	expression := "IF "

	for index, antecedent := range antecedents {
		if index > 0 {
			expression += " AND "
		}

		expression += antecedent.VariableKey + " IS "
		if antecedent.Negated {
			expression += "NOT "
		}
		expression += antecedent.SetKey
	}

	return expression + " THEN " + ValveOpeningKey + " IS " + outputSet
}
