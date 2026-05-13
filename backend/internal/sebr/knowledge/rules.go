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
		rule("R01", "Detect fever", domain.NumberComparison(temperature, domain.GreaterThanOrEqualOperator, 38), fever, defaultPriority, "temperatura >= 38"),
		rule("R02", "Detect suspected fever case", domain.Any(domain.BoolFact(headache, true), domain.BoolFact(fatigue, true)), suspectedFeverCase, defaultPriority, "dolor_cabeza == true OR fatiga == true"),
		rule("R03", "Confirm fever case", domain.All(domain.BoolFact(fever, true), domain.BoolFact(suspectedFeverCase, true)), feverCase, defaultPriority, "fiebre AND sospecha_cuadro_febril"),
		rule("R04", "Detect suspected respiratory case", domain.Any(domain.BoolFact(cough, true), domain.BoolFact(breathingDifficulty, true)), suspectedRespiratoryCase, defaultPriority, "tos == true OR dificultad_respiratoria == true"),
		rule("R05", "Confirm respiratory case", domain.All(domain.BoolFact(feverCase, true), domain.BoolFact(suspectedRespiratoryCase, true)), respiratoryCase, defaultPriority, "cuadro_febril AND sospecha_cuadro_respiratorio"),
		rule("R06", "Mark respiratory fever case as evaluable", domain.All(domain.BoolFact(feverCase, true), domain.BoolFact(respiratoryCase, true)), evaluableCase, defaultPriority, "cuadro_febril AND cuadro_respiratorio"),
		rule("R07", "Mark vomiting fever case as evaluable", domain.All(domain.BoolFact(feverCase, true), domain.BoolFact(vomiting, true)), evaluableCase, defaultPriority, "cuadro_febril AND vomito == true"),
		rule("R08", "Mark respiratory vomiting case as evaluable", domain.All(domain.BoolFact(respiratoryCase, true), domain.BoolFact(vomiting, true)), evaluableCase, defaultPriority, "cuadro_respiratorio AND vomito == true"),
		rule("R09", "Evidence from fever", domain.BoolFact(fever, true), hasEvidence, defaultPriority, "fiebre"),
		rule("R10", "Evidence from fever case", domain.BoolFact(feverCase, true), hasEvidence, defaultPriority, "cuadro_febril"),
		rule("R11", "Evidence from suspected respiratory case", domain.BoolFact(suspectedRespiratoryCase, true), hasEvidence, defaultPriority, "sospecha_cuadro_respiratorio"),
		rule("R12", "Evidence from respiratory case", domain.BoolFact(respiratoryCase, true), hasEvidence, defaultPriority, "cuadro_respiratorio"),
		rule("R13", "Evidence from evaluable case", domain.BoolFact(evaluableCase, true), hasEvidence, defaultPriority, "caso_evaluable"),
		rule("R14", "COVID compatibility from respiratory case and smell or taste loss", domain.All(domain.BoolFact(respiratoryCase, true), domain.BoolFact(lossOfTasteOrSmell, true)), covidCompatible, defaultPriority, "cuadro_respiratorio AND perdida_gusto_u_olfato == true"),
		rule("R15", "Suspect COVID from evaluable case and close contact", domain.All(domain.BoolFact(evaluableCase, true), domain.BoolFact(closeContact, true)), suspectedCovid, defaultPriority, "caso_evaluable AND contacto_estrecho == true"),
		rule("R16", "Suspect COVID reinfection", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(hadCovidBefore, true)), suspectedCovidReinfection, defaultPriority, "sospecha_covid AND tuvo_covid_antes == true"),
		rule("R17", "COVID compatibility from reinfection and chest pain", domain.All(domain.BoolFact(suspectedCovidReinfection, true), domain.BoolFact(chestPain, true)), covidCompatible, defaultPriority, "sospecha_recontagio AND dolor_toracico == true"),
		rule("R18", "COVID compatibility from reinfection and sore throat", domain.All(domain.BoolFact(suspectedCovidReinfection, true), domain.BoolFact(soreThroat, true)), covidCompatible, defaultPriority, "sospecha_recontagio AND dolor_garganta == true"),
		rule("R19", "Suspect COVID from cough and smell or taste loss", domain.All(domain.BoolFact(cough, true), domain.BoolFact(lossOfTasteOrSmell, true)), suspectedCovid, defaultPriority, "tos == true AND perdida_gusto_u_olfato == true"),
		rule("R20", "Suspect dengue from evaluable case and bleeding", domain.All(domain.BoolFact(evaluableCase, true), domain.BoolFact(bleeding, true)), suspectedDengue, defaultPriority, "caso_evaluable AND sangrado == true"),
		rule("R21", "Suspect dengue from evaluable case and nausea", domain.All(domain.BoolFact(evaluableCase, true), domain.BoolFact(nausea, true)), suspectedDengue, defaultPriority, "caso_evaluable AND nausea == true"),
		rule("R22", "Dengue compatibility from suspected dengue and pain behind eyes", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(painBehindEyes, true)), dengueCompatible, defaultPriority, "sospecha_dengue AND dolor_detras_de_ojos == true"),
		rule("R23", "Dengue compatibility from suspected dengue and muscle or joint pain", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(muscleOrJointPain, true)), dengueCompatible, defaultPriority, "sospecha_dengue AND dolor_muscular_o_articular == true"),
		rule("R24", "Cannot conclude from low temperature", domain.NumberComparison(temperature, domain.LessThanOperator, 38), cannotConclude, defaultPriority, "temperatura < 38"),
		rule("R25", "Cannot conclude from COVID suspicion with bleeding", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(bleeding, true)), cannotConclude, defaultPriority, "sospecha_covid AND sangrado == true"),
		rule("R26", "Cannot conclude from COVID suspicion with muscle or joint pain", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(muscleOrJointPain, true)), cannotConclude, defaultPriority, "sospecha_covid AND dolor_muscular_o_articular == true"),
		rule("R27", "Cannot conclude from COVID suspicion with nausea", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(nausea, true)), cannotConclude, defaultPriority, "sospecha_covid AND nausea == true"),
		rule("R28", "Cannot conclude from COVID suspicion with pain behind eyes", domain.All(domain.BoolFact(suspectedCovid, true), domain.BoolFact(painBehindEyes, true)), cannotConclude, defaultPriority, "sospecha_covid AND dolor_detras_de_ojos == true"),
		rule("R29", "Cannot conclude from dengue suspicion with smell or taste loss", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(lossOfTasteOrSmell, true)), cannotConclude, defaultPriority, "sospecha_dengue AND perdida_gusto_u_olfato == true"),
		rule("R30", "Cannot conclude from dengue suspicion with sore throat", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(soreThroat, true)), cannotConclude, defaultPriority, "sospecha_dengue AND dolor_garganta == true"),
		rule("R31", "Cannot conclude from dengue suspicion with chest pain", domain.All(domain.BoolFact(suspectedDengue, true), domain.BoolFact(chestPain, true)), cannotConclude, defaultPriority, "sospecha_dengue AND dolor_toracico == true"),
		rule("R32", "Cannot conclude when evidence does not match COVID or dengue", domain.All(domain.BoolFact(hasEvidence, true), domain.Not(domain.BoolFact(covidCompatible, true)), domain.Not(domain.BoolFact(dengueCompatible, true))), cannotConclude, lowPriority, "hay_indicios AND NOT compatible_con_covid AND NOT compatible_con_dengue"),
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
