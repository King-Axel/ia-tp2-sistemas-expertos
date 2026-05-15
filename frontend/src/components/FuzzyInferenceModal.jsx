import { X } from "lucide-react";
import {
  createSelectedMembershipPoints,
  formatValue,
} from "../utils/membershipFunctions.js";

function FuzzyInferenceModal({
  isOpen,
  onClose,
  variables,
  values,
  outputResult,
  activatedRules,
  inferenceResult = null,
  variableLabels = {},
  inputSetLabels = {},
  outputSetLabels = {},
}) {
  if (!isOpen) {
    return null;
  }

  function handlePanelClick(event) {
    event.stopPropagation();
  }

  function getVariableKey(variable) {
    return variable.backendKey ?? variable.id;
  }

  function getVariableLabel(variable) {
    return variableLabels[getVariableKey(variable)] ?? variable.label;
  }

  function getMemberships(variable) {
    const variableKey = getVariableKey(variable);

    if (inferenceResult) {
      return inferenceResult.input_memberships?.[variableKey] ?? [];
    }

    return createSelectedMembershipPoints(variable, values[variable.id]);
  }

  function getMembershipLabel(variable, membership) {
    const variableKey = getVariableKey(variable);
    const setKey = membership.set_key ?? membership.name;

    return inputSetLabels[variableKey]?.[setKey] ?? setKey;
  }

  function getMembershipDegree(membership) {
    return membership.degree ?? membership.value?.[1] ?? 0;
  }

  function getOutputValue() {
    return inferenceResult?.crisp_output_value ?? outputResult.value;
  }

  function getOutputLabel() {
    const outputSetKey = inferenceResult?.dominant_output_set;

    if (outputSetKey) {
      return outputSetLabels[outputSetKey] ?? outputSetKey;
    }

    return outputResult.interpretation;
  }

  function getActivatedRules() {
    return inferenceResult?.activated_rules ?? activatedRules;
  }

  function getRuleId(rule) {
    return rule.rule_id ?? rule.id;
  }

  function getRuleActivationDegree(rule) {
    return rule.activation_degree ?? rule.activationDegree ?? 0;
  }

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-[var(--color-overlay)] px-4 py-6"
      onClick={onClose}
    >
      <div
        className="flex max-h-[88vh] w-full max-w-4xl flex-col rounded-lg border border-[var(--color-border)] bg-[var(--color-surface)]"
        onClick={handlePanelClick}
      >
        <div className="flex items-start justify-between gap-4 border-b border-[var(--color-border)] px-5 py-4">
          <div>
            <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-[var(--color-text-subtle)]">
              Resultado de inferencia
            </p>
            <h2 className="mt-1 font-[var(--font-display)] text-xl font-semibold text-[var(--color-text)]">
              Resultado del sistema difuso
            </h2>
          </div>

          <button
            type="button"
            onClick={onClose}
            className="rounded-md border border-[var(--color-border)] bg-[var(--color-surface-raised)] p-2 text-[var(--color-text-subtle)] transition duration-200 hover:border-[var(--color-border-hover)] hover:text-[var(--color-text-soft)] focus:outline-none focus:ring-2 focus:ring-[var(--color-accent-soft)]"
            aria-label="Cerrar modal"
          >
            <X className="h-4 w-4" strokeWidth={1.8} />
          </button>
        </div>

        <div className="space-y-6 overflow-y-auto px-5 py-5">
          <section>
            <h3 className="font-[var(--font-display)] text-sm font-semibold text-[var(--color-text)]">
              Valores de entrada
            </h3>
            <div className="mt-3 grid gap-3 md:grid-cols-3">
              {variables.map((variable) => {
                const value = values[variable.id];
                const memberships = getMemberships(variable);

                return (
                  <article
                    key={variable.id}
                    className="rounded-md border border-[var(--color-border)] bg-[var(--color-page-panel)] p-3"
                  >
                    <p className="font-[var(--font-display)] text-sm font-semibold text-[var(--color-text-soft)]">
                      {getVariableLabel(variable)}
                    </p>
                    <p className="mt-2 font-mono text-lg text-[var(--color-text)]">
                      {formatValue(value, variable.unit, variable.formatValue)}
                    </p>
                    <p className="mt-3 font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.14em] text-[var(--color-text-faint)]">
                      Pertenece a
                    </p>
                    <div className="mt-2 space-y-1">
                      {memberships.map((membership) => (
                        <p
                          key={membership.set_key ?? membership.name}
                          className="flex justify-between gap-3 font-mono text-xs text-[var(--color-text-muted)]"
                        >
                          <span>{getMembershipLabel(variable, membership)}</span>
                          <span>{Number(getMembershipDegree(membership)).toFixed(2)}</span>
                        </p>
                      ))}
                    </div>
                  </article>
                );
              })}
            </div>
          </section>

          <section className="rounded-lg border border-[var(--color-border)] bg-[var(--color-page-panel)] p-4">
            <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-[var(--color-text-faint)]">
              Resultado final
            </p>
            <div className="mt-3 flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
              <p className="font-mono text-4xl font-semibold text-[var(--color-text)]">
                {Number(getOutputValue()).toFixed(2)}%
              </p>
              <div className="text-left sm:text-right">
                <p className="text-xs text-[var(--color-text-subtle)]">Conjunto difuso predominante</p>
                <p className="mt-1 font-[var(--font-display)] text-lg font-semibold text-[var(--color-text)]">
                  "{getOutputLabel()}"
                </p>
              </div>
            </div>
          </section>

          <section>
            <h3 className="font-[var(--font-display)] text-sm font-semibold text-[var(--color-text)]">
              Reglas activadas
            </h3>
            <div className="mt-3 grid gap-2">
              {getActivatedRules().map((rule) => (
                <article
                  key={getRuleId(rule)}
                  className="rounded-md border border-[var(--color-border)] bg-[var(--color-page-panel)] px-3 py-3"
                >
                  <div className="flex flex-wrap items-center justify-between gap-3">
                    <p className="font-mono text-xs font-semibold text-[var(--color-text-subtle)]">{getRuleId(rule)}</p>
                    <p className="font-[var(--font-display)] text-xs font-semibold text-[var(--color-text-muted)]">
                      Grado de activación: {getRuleActivationDegree(rule).toFixed(2)}
                    </p>
                  </div>
                  <p className="mt-2 whitespace-pre-line rounded-md border border-[var(--color-border)] bg-[var(--color-page)] px-3 py-2 font-mono text-xs leading-5 text-[var(--color-text-muted)]">
                    {rule.expression}
                  </p>
                </article>
              ))}
            </div>
          </section>
        </div>
      </div>
    </div>
  );
}

export default FuzzyInferenceModal;
