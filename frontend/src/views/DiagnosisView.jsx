import { useState } from "react";
import { getSebrRules, inferSebr } from "../api/sebrApi.js";
import CheckboxField from "../components/CheckboxField.jsx";
import InferenceResultModal from "../components/InferenceResultModal.jsx";
import PageHeader from "../components/PageHeader.jsx";
import Panel from "../components/Panel.jsx";
import PrimaryButton from "../components/PrimaryButton.jsx";
import RulesModal from "../components/RulesModal.jsx";

const symptoms = [
  { name: "headache", label: "Dolor de cabeza" },
  { name: "cough", label: "Tos" },
  { name: "fatigue", label: "Fatiga" },
  { name: "lossOfTasteOrSmell", label: "Pérdida de gusto y/o olfato" },
  { name: "breathingDifficulty", label: "Dificultad respiratoria" },
  { name: "chestPain", label: "Dolor torácico" },
  { name: "soreThroat", label: "Dolor de garganta" },
  { name: "eyePain", label: "Dolor detrás de ojos" },
  { name: "vomiting", label: "Vómito" },
  { name: "muscleOrJointPain", label: "Dolor muscular o articular" },
  { name: "nausea", label: "Náuseas" },
  { name: "bleeding", label: "Sangrado" },
  { name: "covidContact", label: "Contacto con caso COVID" },
  { name: "hadCovidBefore", label: "Tuvo COVID antes" },
];

const initialFormData = {
  temperature: "",
  ...Object.fromEntries(symptoms.map((symptom) => [symptom.name, false])),
};

function DiagnosisView() {
  const [formData, setFormData] = useState(initialFormData);
  const [isInferenceModalOpen, setIsInferenceModalOpen] = useState(false);
  const [isRulesModalOpen, setIsRulesModalOpen] = useState(false);
  const [inferenceResult, setInferenceResult] = useState(null);
  const [isInferring, setIsInferring] = useState(false);
  const [inferenceError, setInferenceError] = useState("");
  const [rules, setRules] = useState([]);
  const [isLoadingRules, setIsLoadingRules] = useState(false);
  const [rulesError, setRulesError] = useState("");

  const temperature = Number(formData.temperature);
  const hasTemperature = formData.temperature !== "";
  const isTemperatureValid =
    hasTemperature && Number.isFinite(temperature) && temperature >= 36 && temperature <= 42;

  function handleTemperatureChange(event) {
    setFormData((currentData) => ({
      ...currentData,
      temperature: event.target.value,
    }));
  }

  function handleCheckboxChange(event) {
    const { name, checked } = event.target;

    setFormData((currentData) => ({
      ...currentData,
      [name]: checked,
    }));
  }

  async function handleSubmit(event) {
    event.preventDefault();

    if (!isTemperatureValid) {
      return;
    }

    setIsInferring(true);
    setInferenceError("");

    try {
      const response = await inferSebr(mapFormDataToInferencePayload(formData, temperature));
      setInferenceResult(response);
      setIsInferenceModalOpen(true);
    } catch (error) {
      setInferenceError(error.message);
    } finally {
      setIsInferring(false);
    }
  }

  async function handleOpenRulesModal() {
    setIsRulesModalOpen(true);
    await loadRules();
  }

  async function loadRules() {
    setIsLoadingRules(true);
    setRulesError("");

    try {
      const response = await getSebrRules();
      setRules(response.rules ?? []);
    } catch (error) {
      setRulesError(error.message);
    } finally {
      setIsLoadingRules(false);
    }
  }

  return (
    <div className="space-y-8">
      <PageHeader
        title="Sistema de diagnóstico de dengue y covid"
        description="Realizar un sistema basado en reglas donde ayude a detectar si una persona tiene dengue o covid, o en caso contrario indicar que no se puede identificar enfermedad. Las reglas deben ser claras en cada situación."
      />

      <Panel title="Registro de síntomas" eyebrow="Entrada clínica">
        <form onSubmit={handleSubmit} className="space-y-7">
          <div>
            <label
              htmlFor="temperature"
              className="mb-2 block font-[var(--font-display)] text-[13px] font-semibold text-zinc-200"
            >
              Temperatura
            </label>
            <div className="flex max-w-xs items-center rounded-md border border-zinc-800 bg-[#0d0d0d] transition duration-200 focus-within:border-[#0070f3] focus-within:ring-2 focus-within:ring-[#0070f3]/20">
              <input
                id="temperature"
                type="number"
                name="temperature"
                min="36"
                max="42"
                step="0.1"
                value={formData.temperature}
                onChange={handleTemperatureChange}
                placeholder="37.5"
                className="h-11 w-full bg-transparent px-3 font-mono text-sm text-zinc-100 outline-none placeholder:text-zinc-600"
              />
              <span className="border-l border-zinc-800 px-3 font-[var(--font-display)] text-xs font-semibold text-zinc-500">
                °C
              </span>
            </div>
            {hasTemperature && !isTemperatureValid ? (
              <p className="mt-2 text-xs text-red-300">
                La temperatura debe estar entre 36 y 42.
              </p>
            ) : null}
          </div>

          <div>
            <p className="mb-3 font-[var(--font-display)] text-[13px] font-semibold text-zinc-200">
              Indicadores clínicos
            </p>
            <div className="grid gap-2.5 sm:grid-cols-2 lg:grid-cols-3">
              {symptoms.map((symptom) => (
                <CheckboxField
                  key={symptom.name}
                  name={symptom.name}
                  label={symptom.label}
                  checked={formData[symptom.name]}
                  onChange={handleCheckboxChange}
                />
              ))}
            </div>
          </div>

          <div className="flex flex-col justify-end gap-3 border-t border-zinc-800 pt-5 sm:flex-row">
            <button
              type="button"
              onClick={handleOpenRulesModal}
              className="rounded-md border border-zinc-800 bg-transparent px-4 py-2.5 text-sm font-medium text-zinc-300 transition duration-200 hover:border-zinc-700 hover:bg-[#171717] hover:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
            >
              Consultar reglas
            </button>
            <PrimaryButton type="submit" disabled={!isTemperatureValid || isInferring}>
              {isInferring ? "Procesando..." : "Confirmar datos"}
            </PrimaryButton>
          </div>
          {inferenceError ? (
            <p className="text-right text-sm text-red-300">{inferenceError}</p>
          ) : null}
        </form>
      </Panel>

      <InferenceResultModal
        isOpen={isInferenceModalOpen}
        onClose={() => setIsInferenceModalOpen(false)}
        result={inferenceResult?.result}
        activatedRules={inferenceResult?.activated_rules ?? []}
      />

      <RulesModal
        isOpen={isRulesModalOpen}
        onClose={() => setIsRulesModalOpen(false)}
        rules={rules}
        isLoading={isLoadingRules}
        error={rulesError}
        onRuleCreated={loadRules}
      />
    </div>
  );
}

function mapFormDataToInferencePayload(formData, temperature) {
  return {
    temperature,
    headache: formData.headache,
    cough: formData.cough,
    fatigue: formData.fatigue,
    loss_of_taste_or_smell: formData.lossOfTasteOrSmell,
    breathing_difficulty: formData.breathingDifficulty,
    chest_pain: formData.chestPain,
    sore_throat: formData.soreThroat,
    pain_behind_eyes: formData.eyePain,
    vomiting: formData.vomiting,
    muscle_or_joint_pain: formData.muscleOrJointPain,
    nausea: formData.nausea,
    bleeding: formData.bleeding,
    close_contact: formData.covidContact,
    had_covid_before: formData.hadCovidBefore,
  };
}

export default DiagnosisView;
