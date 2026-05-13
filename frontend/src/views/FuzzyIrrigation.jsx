import { useState } from "react";
import FuzzyInputCard from "../components/FuzzyInputCard.jsx";
import FuzzyInferenceModal from "../components/FuzzyInferenceModal.jsx";
import FuzzyRulesModal from "../components/FuzzyRulesModal.jsx";
import OutputMembershipCard from "../components/OutputMembershipCard.jsx";
import PageHeader from "../components/PageHeader.jsx";
import { formatTime } from "../utils/membershipFunctions.js";

const fuzzyVariables = [
  {
    id: "soilHumidity",
    label: "Humedad del suelo",
    unit: "%",
    universe: [0, 100],
    step: 1,
    initialValue: 42,
    sets: [
      { name: "Seco", range: [0, 25] },
      { name: "Poco seco", range: [20, 45] },
      { name: "Adecuado", range: [40, 70] },
      { name: "Húmedo", range: [65, 85] },
      { name: "Saturado", range: [80, 100] },
    ],
  },
  {
    id: "ambientTemperature",
    label: "Temperatura ambiente",
    unit: "°C",
    universe: [0, 45],
    step: 1,
    initialValue: 28,
    sets: [
      { name: "Muy fría", range: [0, 10] },
      { name: "Fría", range: [5, 18] },
      { name: "Templada", range: [15, 30] },
      { name: "Calurosa", range: [27, 38] },
      { name: "Muy calurosa", range: [35, 45] },
    ],
  },
  {
    id: "timeOfDay",
    label: "Horario del día",
    unit: "",
    inputUnit: "min",
    universe: [0, 1339],
    step: 1,
    initialValue: 870,
    formatValue: formatTime,
    axisFormatter: formatTime,
    sets: [
      { name: "Madrugada", range: [0, 330] },
      { name: "Mañana", range: [300, 690] },
      { name: "Mediodía", range: [660, 810] },
      { name: "Tarde", range: [780, 1140] },
      { name: "Noche", range: [1110, 1339] },
    ],
  },
];

const outputVariable = {
  id: "valveOpening",
  label: "Porcentaje de apertura de válvula",
  unit: "%",
  universe: [0, 100],
  sets: [
    { name: "Cerrada", range: [0, 2] },
    { name: "Poca apertura", range: [1.5, 37.5] },
    { name: "Apertura moderada", range: [35, 75] },
    { name: "Mucha apertura", range: [70, 92] },
    { name: "Apertura completa", range: [90, 100] },
  ],
};

const initialValues = Object.fromEntries(
  fuzzyVariables.map((variable) => [variable.id, variable.initialValue])
);

const mockRules = [
  {
    id: "R01",
    expression:
      "IF humedad ES seco\nAND temperatura ES calurosa\nTHEN apertura ES apertura_completa",
  },
  {
    id: "R02",
    expression:
      "IF humedad ES poco_seco\nAND horario ES tarde\nTHEN apertura ES mucha_apertura",
  },
  {
    id: "R03",
    expression:
      "IF humedad ES adecuado\nAND temperatura ES templada\nTHEN apertura ES apertura_moderada",
  },
  {
    id: "R04",
    expression:
      "IF humedad ES humedo\nOR humedad ES saturado\nTHEN apertura ES cerrada",
  },
];

const mockInferenceResult = {
  value: 68,
  interpretation: "Mucha apertura",
};

const mockActivatedRules = [
  {
    id: "R14",
    expression:
      "IF humedad ES poco_seco\nAND temperatura ES calurosa\nTHEN apertura ES mucha_apertura",
    activationDegree: 0.72,
  },
  {
    id: "R18",
    expression:
      "IF horario ES tarde\nAND humedad ES adecuado\nTHEN apertura ES apertura_moderada",
    activationDegree: 0.41,
  },
];

function FuzzyIrrigation() {
  const [values, setValues] = useState(initialValues);
  const [isRulesModalOpen, setIsRulesModalOpen] = useState(false);
  const [isInferenceModalOpen, setIsInferenceModalOpen] = useState(false);

  function handleValueChange(variableId, value) {
    setValues((currentValues) => ({
      ...currentValues,
      [variableId]: value,
    }));
  }

  function handleInfer() {
    // TODO: Send these fuzzy input values to the backend when inference integration is available.
    console.log(values);
  }

  return (
    <div className="space-y-8">
      <PageHeader
        title="Sistema de riego automático"
        description="Sistema de Riego automático. Basado en humedad de suelo, temperatura del ambiente, horario del día, determinar cuánta agua debe permitir que se riegue. Puede expresarse la salida en unidades o como el porcentaje de apertura de una válvula."
      />

      <section className="flex flex-col justify-between gap-3 rounded-lg border border-zinc-800 bg-[#111111] px-5 py-4 sm:flex-row sm:items-center">
        <div>
          <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-zinc-600">
            Acciones del sistema
          </p>
          <p className="mt-1 text-sm text-zinc-400">
            Consulta reglas, ejecuta una inferencia futura o revisa datos simulados.
          </p>
        </div>
        <div className="flex flex-col gap-3 sm:flex-row">
          <button
            type="button"
            onClick={() => setIsRulesModalOpen(true)}
            className="rounded-md border border-zinc-800 bg-transparent px-4 py-2.5 text-sm font-medium text-zinc-300 transition duration-200 hover:border-zinc-700 hover:bg-[#171717] hover:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
          >
            Consultar reglas
          </button>
          {/* TODO: Remove this temporary test button after backend inference integration. */}
          <button
            type="button"
            onClick={() => setIsInferenceModalOpen(true)}
            className="rounded-md border border-zinc-800 bg-transparent px-4 py-2.5 text-sm font-medium text-zinc-300 transition duration-200 hover:border-zinc-700 hover:bg-[#171717] hover:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
          >
            Ver inferencia simulada
          </button>
          <button
            type="button"
            onClick={handleInfer}
            className="rounded-md border border-zinc-700 bg-[#171717] px-4 py-2.5 text-sm font-medium text-zinc-50 transition duration-200 hover:border-zinc-600 hover:bg-[#1f1f1f] focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
          >
            Inferir
          </button>
        </div>
      </section>

      <div className="grid gap-5 xl:grid-cols-2">
        {fuzzyVariables.map((variable) => (
          <FuzzyInputCard
            key={variable.id}
            variable={variable}
            value={values[variable.id]}
            onChange={(value) => handleValueChange(variable.id, value)}
          />
        ))}
        <OutputMembershipCard variable={outputVariable} />
      </div>

      <FuzzyRulesModal
        isOpen={isRulesModalOpen}
        onClose={() => setIsRulesModalOpen(false)}
        rules={mockRules}
      />

      <FuzzyInferenceModal
        isOpen={isInferenceModalOpen}
        onClose={() => setIsInferenceModalOpen(false)}
        variables={fuzzyVariables}
        values={values}
        outputResult={mockInferenceResult}
        activatedRules={mockActivatedRules}
      />
    </div>
  );
}

export default FuzzyIrrigation;
