package domain

type Antecedent struct {
	VariableKey string
	SetKey      string
	Negated     bool
}

type Consequent struct {
	VariableKey string
	SetKey      string
}

type Rule struct {
	ID          string
	Name        string
	Antecedents []Antecedent
	Consequent  Consequent
	Expression  string
}
