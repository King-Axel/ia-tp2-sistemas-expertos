package domain

type MembershipFunctionType string

const (
	LeftShoulder  MembershipFunctionType = "left_shoulder"
	Triangular    MembershipFunctionType = "triangular"
	RightShoulder MembershipFunctionType = "right_shoulder"
)

type MembershipFunction struct {
	Type  MembershipFunctionType
	Start float64
	End   float64
}

func NewMembershipFunction(functionType MembershipFunctionType, start, end float64) MembershipFunction {
	return MembershipFunction{
		Type:  functionType,
		Start: start,
		End:   end,
	}
}

func (function MembershipFunction) Degree(value float64) float64 {
	midpoint := (function.Start + function.End) / 2

	switch function.Type {
	case LeftShoulder:
		if value < function.Start || value > function.End {
			return 0
		}
		if value <= midpoint {
			return 1
		}
		return (function.End - value) / (function.End - midpoint)
	case RightShoulder:
		if value < function.Start || value > function.End {
			return 0
		}
		if value >= midpoint {
			return 1
		}
		return (value - function.Start) / (midpoint - function.Start)
	case Triangular:
		if value < function.Start || value > function.End {
			return 0
		}
		if value == midpoint {
			return 1
		}
		if value < midpoint {
			return (value - function.Start) / (midpoint - function.Start)
		}
		return (function.End - value) / (function.End - midpoint)
	default:
		return 0
	}
}
