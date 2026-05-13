import { useState } from "react";
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

const mockResult = "compatible_con_dengue";

const mockActivatedRules = [
  {
    id: "R1",
    name: "Temperatura elevada",
    expression: "IF temperatura >= 38 THEN fiebre",
  },
  {
    id: "R2",
    name: "Cuadro febril",
    expression: "IF fiebre AND posible_cuadro_febril THEN cuadro_febril",
  },
  {
    id: "R3",
    name: "Sospecha dengue",
    expression: "IF caso_evaluable AND sangrado == true THEN sospecha_dengue",
  },
  {
    id: "R4",
    name: "Compatible con dengue",
    expression:
      "IF sospecha_dengue AND dolor_muscular_o_articular == true THEN compatible_con_dengue",
  },
];

const mockRules = [
  {
    id: "R1",
    name: "Detección de fiebre",
    expression: "IF temperatura >= 38 THEN fiebre",
  },
  {
    id: "R2",
    name: "Cuadro febril",
    expression: "IF fiebre AND posible_cuadro_febril THEN cuadro_febril",
  },
  {
    id: "R3",
    name: "Compatible con COVID-19",
    expression:
      "IF cuadro_respiratorio AND perdida_gusto_u_olfato == true THEN compatible_con_covid",
  },
  {
    id: "R4",
    name: "Compatible con dengue",
    expression:
      "IF sospecha_dengue AND dolor_muscular_o_articular == true THEN compatible_con_dengue",
  },
];

function DiagnosisView() {
  const [formData, setFormData] = useState(initialFormData);
  const [isInferenceModalOpen, setIsInferenceModalOpen] = useState(false);
  const [isRulesModalOpen, setIsRulesModalOpen] = useState(false);

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

  function handleSubmit(event) {
    event.preventDefault();

    if (!isTemperatureValid) {
      return;
    }

    // TODO: Send this data to the backend when the diagnosis endpoint is available.
    const submittedData = {
      ...formData,
      temperature,
    };

    console.log(submittedData);
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
              onClick={() => setIsRulesModalOpen(true)}
              className="rounded-md border border-zinc-800 bg-transparent px-4 py-2.5 text-sm font-medium text-zinc-300 transition duration-200 hover:border-zinc-700 hover:bg-[#171717] hover:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
            >
              Consultar reglas
            </button>
            {/* TODO: Remove this temporary test button after backend integration. */}
            <button
              type="button"
              onClick={() => setIsInferenceModalOpen(true)}
              className="rounded-md border border-zinc-800 bg-transparent px-4 py-2.5 text-sm font-medium text-zinc-300 transition duration-200 hover:border-zinc-700 hover:bg-[#171717] hover:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
            >
              Ver resultado simulado
            </button>
            <PrimaryButton type="submit" disabled={!isTemperatureValid}>
              Confirmar datos
            </PrimaryButton>
          </div>
        </form>
      </Panel>

      <InferenceResultModal
        isOpen={isInferenceModalOpen}
        onClose={() => setIsInferenceModalOpen(false)}
        result={mockResult}
        activatedRules={mockActivatedRules}
      />

      <RulesModal
        isOpen={isRulesModalOpen}
        onClose={() => setIsRulesModalOpen(false)}
        rules={mockRules}
      />
    </div>
  );
}

export default DiagnosisView;
