import { useState } from "react";
import { Plus, X } from "lucide-react";
import { createSebrRule } from "../api/sebrApi.js";

function RulesModal({ isOpen, onClose, rules, isLoading, error, onRuleCreated }) {
  const [isAddingRule, setIsAddingRule] = useState(false);
  const [ruleForm, setRuleForm] = useState({
    name: "",
    expression: "",
  });
  const [isCreatingRule, setIsCreatingRule] = useState(false);
  const [createError, setCreateError] = useState("");

  if (!isOpen) {
    return null;
  }

  function handlePanelClick(event) {
    event.stopPropagation();
  }

  function handleInputChange(event) {
    const { name, value } = event.target;

    setRuleForm((currentForm) => ({
      ...currentForm,
      [name]: value,
    }));
  }

  function handleCancelAddRule() {
    setIsAddingRule(false);
    setRuleForm({ name: "", expression: "" });
    setCreateError("");
  }

  async function handleSubmitRule(event) {
    event.preventDefault();
    setIsCreatingRule(true);
    setCreateError("");

    try {
      await createSebrRule({
        name: ruleForm.name,
        expression: ruleForm.expression,
      });
      handleCancelAddRule();
      await onRuleCreated();
    } catch (error) {
      setCreateError(error.message);
    } finally {
      setIsCreatingRule(false);
    }
  }

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 px-4 py-6"
      onClick={onClose}
    >
      <div
        className="flex max-h-[88vh] w-full max-w-3xl flex-col rounded-lg border border-zinc-800 bg-[#111111]"
        onClick={handlePanelClick}
      >
        <div className="flex items-start justify-between gap-4 border-b border-zinc-800 px-5 py-4">
          <div>
            <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-zinc-500">
              Base de conocimiento
            </p>
            <h2 className="mt-1 font-[var(--font-display)] text-xl font-semibold text-zinc-50">
              Reglas del sistema
            </h2>
            <p className="mt-2 max-w-xl text-sm leading-6 text-zinc-400">
              Consulta las reglas utilizadas por el sistema experto basado en reglas.
            </p>
          </div>

          <button
            type="button"
            onClick={onClose}
            className="rounded-md border border-zinc-800 bg-[#171717] p-2 text-zinc-500 transition duration-200 hover:border-zinc-700 hover:text-zinc-200 focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
            aria-label="Cerrar modal"
          >
            <X className="h-4 w-4" strokeWidth={1.8} />
          </button>
        </div>

        <div className="space-y-6 overflow-y-auto px-5 py-5">
          {isLoading ? (
            <p className="rounded-md border border-zinc-800 bg-[#0d0d0d] px-3 py-3 text-sm text-zinc-500">
              Cargando reglas...
            </p>
          ) : null}

          {error ? (
            <p className="rounded-md border border-red-900/60 bg-red-950/20 px-3 py-3 text-sm text-red-300">
              {error}
            </p>
          ) : null}

          {!isLoading && !error ? (
            <div className="grid gap-2">
              {rules.map((rule) => (
              <article
                key={rule.id}
                className="rounded-md border border-zinc-800 bg-[#0d0d0d] px-3 py-3"
              >
                <div className="flex flex-wrap items-center gap-2">
                  <span className="rounded border border-zinc-800 bg-[#171717] px-2 py-0.5 font-mono text-[11px] text-zinc-400">
                    {rule.id}
                  </span>
                  <h3 className="font-[var(--font-display)] text-sm font-semibold text-zinc-200">
                    {rule.name}
                  </h3>
                </div>
                <p className="mt-2 rounded-md border border-zinc-800 bg-[#090909] px-3 py-2 font-mono text-xs leading-5 text-zinc-500">
                  {rule.expression}
                </p>
              </article>
              ))}
            </div>
          ) : null}

          <div className="rounded-lg border border-zinc-800 bg-[#0d0d0d]">
            <div className="flex items-center justify-between gap-4 border-b border-zinc-800 px-4 py-3">
              <div>
                <h3 className="font-[var(--font-display)] text-sm font-semibold text-zinc-200">
                  Agregar regla
                </h3>
                <p className="mt-1 text-xs text-zinc-500">
                  Registra una nueva regla para pruebas locales.
                </p>
              </div>

              {!isAddingRule && (
                <button
                  type="button"
                  onClick={() => setIsAddingRule(true)}
                  className="inline-flex items-center gap-2 rounded-md border border-zinc-700 bg-[#171717] px-3 py-2 text-sm font-medium text-zinc-100 transition duration-200 hover:border-zinc-600 hover:bg-[#1f1f1f] focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
                >
                  <Plus className="h-4 w-4" strokeWidth={1.8} />
                  Agregar nueva regla
                </button>
              )}
            </div>

            {isAddingRule && (
              <form onSubmit={handleSubmitRule} className="space-y-4 p-4">
                <div>
                  <label
                    htmlFor="ruleName"
                    className="mb-2 block font-[var(--font-display)] text-[13px] font-semibold text-zinc-200"
                  >
                    Nombre de la regla
                  </label>
                  <input
                    id="ruleName"
                    name="name"
                    type="text"
                    value={ruleForm.name}
                    onChange={handleInputChange}
                    className="h-10 w-full rounded-md border border-zinc-800 bg-[#090909] px-3 text-sm text-zinc-100 outline-none transition duration-200 placeholder:text-zinc-600 focus:border-[#0070f3] focus:ring-2 focus:ring-[#0070f3]/20"
                    placeholder="Ej: Detección de fiebre"
                  />
                </div>

                <div>
                  <label
                    htmlFor="ruleExpression"
                    className="mb-2 block font-[var(--font-display)] text-[13px] font-semibold text-zinc-200"
                  >
                    Expresión de la regla
                  </label>
                  <textarea
                    id="ruleExpression"
                    name="expression"
                    value={ruleForm.expression}
                    onChange={handleInputChange}
                    rows="3"
                    className="w-full resize-y rounded-md border border-zinc-800 bg-[#090909] px-3 py-2 font-mono text-xs leading-5 text-zinc-100 outline-none transition duration-200 placeholder:text-zinc-600 focus:border-[#0070f3] focus:ring-2 focus:ring-[#0070f3]/20"
                    placeholder="IF temperatura >= 38 THEN fiebre"
                  />
                </div>

                <div className="flex justify-end gap-3">
                  <button
                    type="button"
                    onClick={handleCancelAddRule}
                    className="rounded-md border border-zinc-800 bg-transparent px-4 py-2 text-sm font-medium text-zinc-300 transition duration-200 hover:border-zinc-700 hover:bg-[#171717] hover:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
                  >
                    Cancelar
                  </button>
                  <button
                    type="submit"
                    disabled={isCreatingRule}
                    className="rounded-md border border-zinc-700 bg-[#171717] px-4 py-2 text-sm font-medium text-zinc-50 transition duration-200 hover:border-zinc-600 hover:bg-[#1f1f1f] focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
                  >
                    {isCreatingRule ? "Guardando..." : "Guardar regla"}
                  </button>
                </div>
                {createError ? (
                  <p className="text-sm text-red-300">{createError}</p>
                ) : null}
              </form>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default RulesModal;
