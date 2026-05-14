import { useState } from "react";
import { getSeld1Rules, inferSeld1 } from "../api/seld1Api.js";
import FuzzyInputCard from "../components/FuzzyInputCard.jsx";
import FuzzyInferenceModal from "../components/FuzzyInferenceModal.jsx";
import FuzzyRulesModal from "../components/FuzzyRulesModal.jsx";
import OutputMembershipCard from "../components/OutputMembershipCard.jsx";
import PageHeader from "../components/PageHeader.jsx";
import { formatTime } from "../utils/membershipFunctions.js";

const fuzzyVariables = [
  {
    id: "soilHumidity",
    backendKey: "soil_humidity",
    label: "Humedad del suelo",
    unit: "%",
    universe: [0, 100],
    step: 1,
    initialValue: 42,
    sets: [
      { name: "Muy baja", range: [0, 25] },
      { name: "Baja", range: [20, 44] },
      { name: "Adecuada", range: [40, 70] },
      { name: "Algo alta", range: [65, 85] },
      { name: "Elevada", range: [80, 100] },
    ],
  },
  {
    id: "ambientTemperature",
    backendKey: "ambient_temperature",
    label: "Temperatura ambiente",
    unit: "°C",
    universe: [0, 45],
    step: 1,
    initialValue: 28,
    sets: [
      { name: "Muy fría", range: [0, 10] },
      { name: "Fría", range: [6, 18] },
      { name: "Templada", range: [15, 30] },
      { name: "Cálida", range: [27, 38] },
      { name: "Sofocante", range: [35, 45] },
    ],
  },
  {
    id: "timeOfDay",
    backendKey: "time_of_day",
    label: "Horario del día",
    unit: "",
    inputUnit: "min",
    universe: [0, 1439],
    step: 1,
    initialValue: 870,
    formatValue: formatTime,
    axisFormatter: formatTime,
    sets: [
      { name: "Madrugada", range: [0, 330] },
      { name: "Mañana", range: [300, 690] },
      { name: "Mediodía", range: [660, 810] },
      { name: "Tarde", range: [780, 1140] },
      { name: "Noche", range: [1111, 1439] },
    ],
  },
];

const outputVariable = {
  id: "valveOpening",
  label: "Porcentaje de apertura de válvula",
  unit: "%",
  universe: [0, 100],
  sets: [
    { name: "Nada", range: [0, 2] },
    { name: "Baja", range: [1.5, 37.5] },
    { name: "Moderada", range: [35, 75] },
    { name: "Mucha", range: [70, 92] },
    { name: "Completa", range: [90, 100] },
  ],
};

const initialValues = Object.fromEntries(
  fuzzyVariables.map((variable) => [variable.id, variable.initialValue])
);

const variableLabels = {
  soil_humidity: "Humedad del suelo",
  ambient_temperature: "Temperatura ambiente",
  time_of_day: "Horario del día",
};

const inputSetLabels = {
  soil_humidity: {
    very_low: "Muy baja",
    low: "Baja",
    adequate: "Adecuada",
    somewhat_high: "Algo alta",
    elevated: "Elevada",
  },
  ambient_temperature: {
    very_cold: "Muy fría",
    cold: "Fría",
    mild: "Templada",
    warm: "Cálida",
    sweltering: "Sofocante",
  },
  time_of_day: {
    dawn: "Madrugada",
    morning: "Mañana",
    midday: "Mediodía",
    afternoon: "Tarde",
    night: "Noche",
  },
};

const outputSetLabels = {
  none: "Nada",
  low: "Baja",
  moderate: "Moderada",
  high: "Mucha",
  complete: "Completa",
};

function FuzzyIrrigation() {
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
      const response = await getSeld1Rules();
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
      const response = await inferSeld1(mapValuesToInferencePayload(values));
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
        title="Sistema de riego automático"
        description="Sistema de Riego automático. Basado en humedad de suelo, temperatura del ambiente, horario del día, determinar cuánta agua debe permitir que se riegue. Puede expresarse la salida en unidades o como el porcentaje de apertura de una válvula."
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
    soil_humidity: values.soilHumidity,
    ambient_temperature: values.ambientTemperature,
    time_of_day: values.timeOfDay,
  };
}

export default FuzzyIrrigation;
