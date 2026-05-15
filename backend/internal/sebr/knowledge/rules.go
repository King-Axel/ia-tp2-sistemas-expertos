package knowledge

import "iteraciones/backend/internal/sebr/domain"

const (
	temperature               = "temperatura"
	headache                  = "dolor_cabeza"
	cough                     = "tos"
	fatigue                   = "fatiga"
	lossOfTasteOrSmell        = "perdida_gusto_u_olfato"
	breathingDifficulty       = "dificultad_respiratoria"
	chestPain                 = "dolor_toracico"
	soreThroat                = "dolor_garganta"
	painBehindEyes            = "dolor_detras_de_ojos"
	vomiting                  = "vomito"
	muscleOrJointPain         = "dolor_muscular_o_articular"
	nausea                    = "nausea"
	bleeding                  = "sangrado"
	closeContact              = "contacto_estrecho"
	hadCovidBefore            = "tuvo_covid_antes"
	fever                     = "fiebre"
	suspectedFeverCase        = "sospecha_cuadro_febril"
	feverCase                 = "cuadro_febril"
	suspectedRespiratoryCase  = "sospecha_cuadro_respiratorio"
	respiratoryCase           = "cuadro_respiratorio"
	evaluableCase             = "caso_evaluable"
	hasEvidence               = "hay_indicios"
	covidCompatible           = "compatible_con_covid"
	suspectedCovid            = "sospecha_covid"
	suspectedCovidReinfection = "sospecha_recontagio"
	suspectedDengue           = "sospecha_dengue"
	dengueCompatible          = "compatible_con_dengue"
	cannotConclude            = "no_puede_concluir"
	defaultPriority           = 10
	lowPriority               = 0
)

func Rules() []domain.Rule {
	return []domain.Rule{
		rule("R01", "Detectada fiebre", domain.NumberComparison(temperature, domain.GreaterThanOrEqualOperator, 38), fever, defaultPriority, "temperatura >= 38"),
		rule("R02", "Sospecha de cuadro febril", domain.Any(domain.BoolFact(headache, true), domain.BoolFact(fatigue, true)), suspectedFeverCase, defaultPriority, "dolor_cabeza == true OR fatiga == true"),
		rule("R03", "Confirma cuadro febril", domain.All(domain.BoolFact(fever, true), domain.BoolFact(suspectedFeverCase, true)), feverCase, defaultPriority, "fiebre AND sospecha_cuadro_febril"),
		rule("R04", "Sospecha cuadro respiratorio", domain.Any(domain.BoolFact(cough, true), domain.BoolFact(breathingDifficulty, true)), suspectedRespiratoryCase, defaultPriority, "tos == true OR dificultad_respiratoria == true"),
		rule("R05", "Confirma cuadro respiratorio", domain.All(domain.BoolFact(feverCase, true), domain.BoolFact(suspectedRespiratoryCase, true)), respiratoryCase, defaultPriority, "cuadro_febril AND sospecha_cuadro_respiratorio"),
		rule("R06", "Caso evaluable a partir de cuadro respiratorio y febril", domain.All(domain.BoolFact(feverCase, true), domain.BoolFact(respiratoryCase, true)), evaluableCase, defaultPriority, "cuadro_febril AND cuadro_respiratorio"),
		rule("R07", "Caso evaluable a partir de cuadro febril y vómito", domain.All(domain.BoolFact(feverCase, true), domain.BoolFact(vomiting, true)), evaluableCase, defaultPriority, "cuadro_febril AND vomito == true"),
		rule("R08", "Caso evaluable a partir de cuadro respiratorio y vómito", domain.All(domain.BoolFact(respiratoryCase, true), domain.BoolFact(vomiting, true)), evaluableCase, defaultPriority, "cuadro_respiratorio AND vomito == true"),
		rule("R09", "Tiene fiebre, hay indicios", domain.BoolFact(fever, true), hasEvidence, defaultPriority, "fiebre"),
		// rule("R10", "Evidence from fever case", domain.BoolFact(feverCase, true), hasEvidence, defaultPriority, "cuadro_febril"),
		// rule("R11", "Evidence from suspected respiratory case", domain.BoolFact(suspectedRespiratoryCase, true), hasEvidence, defaultPriority, "sospecha_cuadro_respiratorio"),
		// rule("R12", "Evidence from respiratory case", domain.BoolFact(respiratoryCase, true), hasEvidence, defaultPriority, "cuadro_respiratorio"),
		// rule("R13", "Evidence from evaluable case", domain.BoolFact(evaluableCase, true), hasEvidence, defaultPriority, "caso_evaluable"),
		rule("R14", "Sospecha COVID a partir de cuadro respiratorio y pérdida de gusto y olfato", domain.All(domain.BoolFact(respiratoryCase, true), domain.BoolFact(lossOfTasteOrSmell, true)), covidCompatible, defaultPriority, "cuadro_respiratorio AND perdida_gusto_u_olfato == true"),
		rule("R15", "Sospecha COVID a partir de caso evaluable y contacto estrecho", domain.All(domain.BoolFact(evaluableCase, true), domain.BoolFact(closeContact, true)), suspectedCovid, defaultPriority, "caso_evaluable AND contacto_estrecho == true"),
		rule("R16", "Sospecha recontagio de COVID", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(hadCovidBefore, true)), suspectedCovidReinfection, defaultPriority, "sospecha_covid AND tuvo_covid_antes == true"),
		rule("R17", "Compatibilidad con COVID a partir de sospecha de recontagio y dolor torácico", domain.All(domain.BoolFact(suspectedCovidReinfection, true), domain.BoolFact(chestPain, true)), covidCompatible, defaultPriority, "sospecha_recontagio AND dolor_toracico == true"),
		rule("R18", "Compatibilidad con COVID a partir de sospecha de recontagio y dolor de garganta", domain.All(domain.BoolFact(suspectedCovidReinfection, true), domain.BoolFact(soreThroat, true)), covidCompatible, defaultPriority, "sospecha_recontagio AND dolor_garganta == true"),
		rule("R19", "Sospecha de COVID a partir de tos y pérdida de gusto u olfato", domain.All(domain.BoolFact(cough, true), domain.BoolFact(lossOfTasteOrSmell, true)), suspectedCovid, defaultPriority, "tos == true AND perdida_gusto_u_olfato == true"),
		rule("R20", "Sospecha Dengue a partir de caso evaluable y sangrado", domain.All(domain.BoolFact(evaluableCase, true), domain.BoolFact(bleeding, true)), suspectedDengue, defaultPriority, "caso_evaluable AND sangrado == true"),
		rule("R21", "Sospecha Dengue a partir de caso evaluable y náuseas", domain.All(domain.BoolFact(evaluableCase, true), domain.BoolFact(nausea, true)), suspectedDengue, defaultPriority, "caso_evaluable AND nausea == true"),
		rule("R22", "Compatibilidad con Dengue a partir de sospecha de Dengue y dolor detrás de ojos", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(painBehindEyes, true)), dengueCompatible, defaultPriority, "sospecha_dengue AND dolor_detras_de_ojos == true"),
		rule("R23", "Compatibilidad con Dengue a partir de sospecha de Dengue y dolor muscular o articular", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(muscleOrJointPain, true)), dengueCompatible, defaultPriority, "sospecha_dengue AND dolor_muscular_o_articular == true"),
		rule("R24", "No puede concluir por baja temperatura", domain.NumberComparison(temperature, domain.LessThanOperator, 38), cannotConclude, defaultPriority, "temperatura < 38"),
		rule("R25", "No puede concluir a partir de sospecha de COVID con sangrado", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(bleeding, true)), cannotConclude, defaultPriority, "sospecha_covid AND sangrado == true"),
		rule("R26", "No puede concluir a partir de sospecha de COVID con dolor muscular o articular", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(muscleOrJointPain, true)), cannotConclude, defaultPriority, "sospecha_covid AND dolor_muscular_o_articular == true"),
		rule("R27", "No puede concluir a partir de sospecha de COVID con náuseas", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(nausea, true)), cannotConclude, defaultPriority, "sospecha_covid AND nausea == true"),
		rule("R28", "No puede concluir a partir de sospecha de COVID con dolor detrás de ojos", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(painBehindEyes, true)), cannotConclude, defaultPriority, "sospecha_covid AND dolor_detras_de_ojos == true"),
		rule("R29", "No puede concluir a partir de sospecha de Dengue con pérdida de gusto u olfato", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(lossOfTasteOrSmell, true)), cannotConclude, defaultPriority, "sospecha_dengue AND perdida_gusto_u_olfato == true"),
		rule("R30", "No puede concluir a partir de sospecha de Dengue con dolor de garganta", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(soreThroat, true)), cannotConclude, defaultPriority, "sospecha_dengue AND dolor_garganta == true"),
		rule("R31", "No puede concluir a partir de sospecha de Dengue con dolor torácico", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(chestPain, true)), cannotConclude, defaultPriority, "sospecha_dengue AND dolor_toracico == true"),
		rule("R32", "No puede concluir a partir de indicios que no coinciden con COVID o Dengue", domain.All(domain.BoolFact(hasEvidence, true), domain.Not(domain.BoolFact(covidCompatible, true)), domain.Not(domain.BoolFact(dengueCompatible, true))), cannotConclude, lowPriority, "hay_indicios AND NOT compatible_con_covid AND NOT compatible_con_dengue"),
	}
}

func FinalConclusionKeys() []string {
	return []string{
		covidCompatible,
		dengueCompatible,
		cannotConclude,
	}
}

func rule(id string, name string, condition domain.Condition, conclusionKey string, priority int, expression string) domain.Rule {
	return domain.Rule{
		ID:         id,
		Name:       name,
		Condition:  condition,
		Conclusion: domain.NewBoolFact(conclusionKey, true),
		Priority:   priority,
		Expression: expression,
	}
}
