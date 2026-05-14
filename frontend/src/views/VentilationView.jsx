import { useState } from "react";
import FuzzyInferenceModal from "../components/FuzzyInferenceModal.jsx";
import FuzzyInputCard from "../components/FuzzyInputCard.jsx";
import FuzzyRulesModal from "../components/FuzzyRulesModal.jsx";
import OutputMembershipCard from "../components/OutputMembershipCard.jsx";
import PageHeader from "../components/PageHeader.jsx";

const fuzzyVariables = [
  {
    id: "ambientTemperature",
    label: "Temperatura ambiente",
    unit: "°C",
    universe: [0, 60],
    step: 1,
    initialValue: 25,
    sets: [
      { name: "Baja", range: [0, 20] },
      { name: "Normal", range: [15, 30] },
      { name: "Alta", range: [25, 45] },
      { name: "Crítica", range: [40, 60] },
    ],
  },
  {
    id: "gasConcentration",
    label: "Concentración de gases/humo",
    unit: "%",
    universe: [0, 100],
    step: 1,
    initialValue: 35,
    sets: [
      { name: "Seguro", range: [0, 25] },
      { name: "Moderado", range: [20, 60] },
      { name: "Peligroso", range: [50, 100] },
    ],
  },
];

const outputVariable = {
  id: "extractorMotorSpeed",
  label: "Velocidad del motor del extractor",
  unit: "%",
  universe: [0, 100],
  sets: [
    { name: "Mínima / Reposo", range: [0, 15] },
    { name: "Baja", range: [10, 40] },
    { name: "Media", range: [30, 70] },
    { name: "Alta", range: [60, 90] },
    { name: "Máxima", range: [80, 100] },
  ],
};

const initialValues = Object.fromEntries(
  fuzzyVariables.map((variable) => [variable.id, variable.initialValue])
);

const mockRules = [
  {
    id: "R1",
    expression: "IF temperatura ES baja AND gases ES seguro THEN velocidad ES mínima",
  },
  {
    id: "R2",
    expression: "IF temperatura ES normal AND gases ES seguro THEN velocidad ES baja",
  },
  {
    id: "R3",
    expression: "IF temperatura ES alta AND gases ES moderado THEN velocidad ES media",
  },
  {
    id: "R4",
    expression: "IF temperatura ES crítica OR gases ES peligroso THEN velocidad ES máxima",
  },
  {
    id: "R5",
    expression: "IF gases ES peligroso THEN velocidad ES máxima",
  },
];

const mockInferenceResult = {
  value: 62,
  interpretation: "Media",
};

const mockActivatedRules = [
  {
    id: "R1",
    expression: "IF temperatura ES normal AND gases ES seguro THEN velocidad ES baja",
    activationDegree: 0.45,
  },
  {
    id: "R2",
    expression: "IF temperatura ES alta OR gases ES peligroso THEN velocidad ES alta",
    activationDegree: 0.72,
  },
];

function VentilationView() {
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
        title="Sistema de ventilación industrial"
        description="Se requiere automatizar el sistema de extracción de aire y ventilación. La velocidad del motor del extractor debe regularse de forma suave y continua para mantener un ambiente seguro y optimizar el consumo eléctrico. El sistema difuso debe tomar decisiones basadas en dos sensores principales: temperatura del ambiente en la planta y nivel de concentración de gases o humo."
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
        title="Reglas de ventilación industrial"
        description="Consulta las reglas simuladas para el sistema difuso de ventilación industrial."
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

export default VentilationView;
