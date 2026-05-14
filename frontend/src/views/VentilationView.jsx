import { useState } from "react";
import { getSeld2Rules, inferSeld2 } from "../api/seld2Api.js";
import FuzzyInferenceModal from "../components/FuzzyInferenceModal.jsx";
import FuzzyInputCard from "../components/FuzzyInputCard.jsx";
import FuzzyRulesModal from "../components/FuzzyRulesModal.jsx";
import OutputMembershipCard from "../components/OutputMembershipCard.jsx";
import PageHeader from "../components/PageHeader.jsx";

const fuzzyVariables = [
  {
    id: "ambientTemperature",
    backendKey: "plant_temperature",
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
    backendKey: "gas_concentration",
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

const variableLabels = {
  plant_temperature: "Temperatura ambiente",
  gas_concentration: "Concentración de gases/humo",
};

const inputSetLabels = {
  plant_temperature: {
    low: "Baja",
    normal: "Normal",
    high: "Alta",
    critical: "Crítica",
  },
  gas_concentration: {
    safe: "Seguro",
    moderate: "Moderado",
    dangerous: "Peligroso",
  },
};

const outputSetLabels = {
  minimum: "Mínima / Reposo",
  low: "Baja",
  medium: "Media",
  high: "Alta",
  maximum: "Máxima",
};

function VentilationView() {
  const [values, setValues] = useState(initialValues);
  const [isRulesModalOpen, setIsRulesModalOpen] = useState(false);
  const [isInferenceModalOpen, setIsInferenceModalOpen] = useState(false);
  const [rules, setRules] = useState([]);
  const [inferenceResult, setInferenceResult] = useState(null);
  const [isLoadingRules, setIsLoadingRules] = useState(false);
  const [isInferring, setIsInferring] = useState(false);
  const [rulesError, setRulesError] = useState("");
  const [inferenceError, setInferenceError] = useState("");

  function handleValueChange(variableId, value) {
    setValues((currentValues) => ({
      ...currentValues,
      [variableId]: value,
    }));
  }

  async function handleOpenRulesModal() {
    setIsRulesModalOpen(true);
    setRulesError("");
    setIsLoadingRules(true);

    try {
      const response = await getSeld2Rules();
      setRules(response.rules ?? []);
    } catch (error) {
      setRulesError(error.message);
    } finally {
      setIsLoadingRules(false);
    }
  }

  async function handleInfer() {
    setInferenceError("");
    setIsInferring(true);

    try {
      const response = await inferSeld2(mapValuesToInferencePayload(values));
      setInferenceResult(response);
      setIsInferenceModalOpen(true);
    } catch (error) {
      setInferenceError(error.message);
    } finally {
      setIsInferring(false);
    }
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
            Consulta reglas o ejecuta una inferencia con los valores actuales.
          </p>
        </div>
        <div className="flex flex-col gap-3 sm:flex-row">
          <button
            type="button"
            onClick={handleOpenRulesModal}
            disabled={isLoadingRules}
            className="rounded-md border border-zinc-800 bg-transparent px-4 py-2.5 text-sm font-medium text-zinc-300 transition duration-200 hover:border-zinc-700 hover:bg-[#171717] hover:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
          >
            {isLoadingRules ? "Cargando..." : "Consultar reglas"}
          </button>
          <button
            type="button"
            onClick={handleInfer}
            disabled={isInferring}
            className="rounded-md border border-zinc-700 bg-[#171717] px-4 py-2.5 text-sm font-medium text-zinc-50 transition duration-200 hover:border-zinc-600 hover:bg-[#1f1f1f] focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
          >
            {isInferring ? "Infiriendo..." : "Inferir"}
          </button>
        </div>
      </section>

      {inferenceError ? (
        <p className="rounded-md border border-red-900/60 bg-red-950/20 px-4 py-3 text-sm text-red-300">
          {inferenceError}
        </p>
      ) : null}

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
        rules={rules}
        title="Reglas de ventilación industrial"
        description="Consulta las reglas cargadas para el sistema difuso de ventilación industrial."
        isLoading={isLoadingRules}
        error={rulesError}
      />

      {inferenceResult ? (
        <FuzzyInferenceModal
          isOpen={isInferenceModalOpen}
          onClose={() => setIsInferenceModalOpen(false)}
          variables={fuzzyVariables}
          values={values}
          outputResult={{
            value: inferenceResult.crisp_output_value,
            interpretation: outputSetLabels[inferenceResult.dominant_output_set],
          }}
          activatedRules={inferenceResult.activated_rules ?? []}
          inferenceResult={inferenceResult}
          variableLabels={variableLabels}
          inputSetLabels={inputSetLabels}
          outputSetLabels={outputSetLabels}
        />
      ) : null}
    </div>
  );
}

function mapValuesToInferencePayload(values) {
  return {
    plant_temperature: values.ambientTemperature,
    gas_concentration: values.gasConcentration,
  };
}

export default VentilationView;
