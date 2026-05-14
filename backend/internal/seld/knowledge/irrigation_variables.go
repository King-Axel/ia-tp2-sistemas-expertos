package knowledge

import "iteraciones/backend/internal/seld/domain"

const (
	SoilHumidityKey       = "soil_humidity"
	AmbientTemperatureKey = "ambient_temperature"
	TimeOfDayKey          = "time_of_day"
	ValveOpeningKey       = "valve_opening"
)

func IrrigationInputVariables() []domain.Variable {
	return []domain.Variable{
		{
			Key:      SoilHumidityKey,
			Name:     "Soil humidity",
			Universe: domain.Universe{Min: 0, Max: 100},
			Sets: []domain.FuzzySet{
				fuzzySet("very_low", "Very low", domain.LeftShoulder, 0, 25),
				fuzzySet("low", "Low", domain.Triangular, 20, 44),
				fuzzySet("adequate", "Adequate", domain.Triangular, 40, 70),
				fuzzySet("somewhat_high", "Somewhat high", domain.Triangular, 65, 85),
				fuzzySet("elevated", "Elevated", domain.RightShoulder, 80, 100),
			},
		},
		{
			Key:      AmbientTemperatureKey,
			Name:     "Ambient temperature",
			Universe: domain.Universe{Min: 0, Max: 45},
			Sets: []domain.FuzzySet{
				fuzzySet("very_cold", "Very cold", domain.LeftShoulder, 0, 10),
				fuzzySet("cold", "Cold", domain.Triangular, 6, 18),
				fuzzySet("mild", "Mild", domain.Triangular, 15, 30),
				fuzzySet("warm", "Warm", domain.Triangular, 27, 38),
				fuzzySet("sweltering", "Sweltering", domain.RightShoulder, 35, 45),
			},
		},
		{
			Key:      TimeOfDayKey,
			Name:     "Time of day",
			Universe: domain.Universe{Min: 0, Max: 1439},
			Sets: []domain.FuzzySet{
				fuzzySet("dawn", "Dawn", domain.LeftShoulder, 0, 330),
				fuzzySet("morning", "Morning", domain.Triangular, 300, 690),
				fuzzySet("midday", "Midday", domain.Triangular, 660, 810),
				fuzzySet("afternoon", "Afternoon", domain.Triangular, 780, 1140),
				fuzzySet("night", "Night", domain.RightShoulder, 1111, 1439),
			},
		},
	}
}

func IrrigationOutputVariable() domain.Variable {
	return domain.Variable{
		Key:      ValveOpeningKey,
		Name:     "Valve opening percentage",
		Universe: domain.Universe{Min: 0, Max: 100},
		Sets: []domain.FuzzySet{
			fuzzySet("none", "None", domain.LeftShoulder, 0, 2),
			fuzzySet("low", "Low", domain.Triangular, 1.5, 37.5),
			fuzzySet("moderate", "Moderate", domain.Triangular, 35, 75),
			fuzzySet("high", "High", domain.Triangular, 70, 92),
			fuzzySet("complete", "Complete", domain.RightShoulder, 90, 100),
		},
	}
}

func fuzzySet(
	key string,
	name string,
	functionType domain.MembershipFunctionType,
	start float64,
	end float64,
) domain.FuzzySet {
	return domain.FuzzySet{
		Key:        key,
		Name:       name,
		Membership: domain.NewMembershipFunction(functionType, start, end),
	}
}
