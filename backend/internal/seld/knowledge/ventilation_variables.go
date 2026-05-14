package knowledge

import "iteraciones/backend/internal/seld/domain"

const (
	PlantTemperatureKey = "plant_temperature"
	GasConcentrationKey = "gas_concentration"
	MotorSpeedKey       = "motor_speed"
)

func VentilationInputVariables() []domain.Variable {
	return []domain.Variable{
		{
			Key:      PlantTemperatureKey,
			Name:     "Plant temperature",
			Universe: domain.Universe{Min: 0, Max: 60},
			Sets: []domain.FuzzySet{
				fuzzySet("low", "Low", domain.LeftShoulder, 0, 20),
				fuzzySet("normal", "Normal", domain.Triangular, 15, 30),
				fuzzySet("high", "High", domain.Triangular, 25, 45),
				fuzzySet("critical", "Critical", domain.RightShoulder, 40, 60),
			},
		},
		{
			Key:      GasConcentrationKey,
			Name:     "Gas/smoke concentration",
			Universe: domain.Universe{Min: 0, Max: 100},
			Sets: []domain.FuzzySet{
				fuzzySet("safe", "Safe", domain.LeftShoulder, 0, 25),
				fuzzySet("moderate", "Moderate", domain.Triangular, 20, 60),
				fuzzySet("dangerous", "Dangerous", domain.RightShoulder, 50, 100),
			},
		},
	}
}

func VentilationOutputVariable() domain.Variable {
	return domain.Variable{
		Key:      MotorSpeedKey,
		Name:     "Extractor motor speed",
		Universe: domain.Universe{Min: 0, Max: 100},
		Sets: []domain.FuzzySet{
			fuzzySet("minimum", "Minimum", domain.LeftShoulder, 0, 15),
			fuzzySet("low", "Low", domain.Triangular, 10, 40),
			fuzzySet("medium", "Medium", domain.Triangular, 30, 70),
			fuzzySet("high", "High", domain.Triangular, 60, 90),
			fuzzySet("maximum", "Maximum", domain.RightShoulder, 80, 100),
		},
	}
}
