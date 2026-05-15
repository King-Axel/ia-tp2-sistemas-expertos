import { X } from "lucide-react";

function FuzzyRulesModal({
  isOpen,
  onClose,
  rules,
  title = "Reglas de riego difuso",
  description = "Consulta las reglas cargadas para el sistema difuso de riego.",
  isLoading = false,
  error = "",
}) {
  if (!isOpen) {
    return null;
  }

  function handlePanelClick(event) {
    event.stopPropagation();
  }

  return (
    <div
      className="scrollbar-dark fixed inset-0 z-50 flex items-center justify-center bg-[var(--color-overlay)] px-4 py-6"
      onClick={onClose}
    >
      <div
        className="flex max-h-[86vh] w-full max-w-3xl flex-col rounded-lg border border-[var(--color-border)] bg-[var(--color-surface)]"
        onClick={handlePanelClick}
      >
        <div className="flex items-start justify-between gap-4 border-b border-[var(--color-border)] px-5 py-4">
          <div>
            <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-[var(--color-text-subtle)]">
              Base de reglas
            </p>
            <h2 className="mt-1 font-[var(--font-display)] text-xl font-semibold text-[var(--color-text)]">
              {title}
            </h2>
            <p className="mt-2 max-w-xl text-sm leading-6 text-[var(--color-text-muted)]">
              {description}
            </p>
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

        <div className="grid gap-2 overflow-y-auto px-5 py-5">
          {isLoading ? (
            <p className="rounded-md border border-[var(--color-border)] bg-[var(--color-page-panel)] px-3 py-3 text-sm text-[var(--color-text-subtle)]">
              Cargando reglas...
            </p>
          ) : error ? (
            <p className="rounded-md border border-[var(--color-danger-border)] bg-[var(--color-danger-soft)] px-3 py-3 text-sm text-[var(--color-danger)]">
              {error}
            </p>
          ) : (
            rules.map((rule) => (
              <article
                key={rule.id}
                className="rounded-md border border-[var(--color-border)] bg-[var(--color-page-panel)] px-3 py-3"
              >
                <p className="font-mono text-[14px] font-semibold text-[var(--color-text-subtle)]">{rule.id}</p>
                <p className="text-[16px] mt-2 whitespace-pre-line rounded-md border border-[var(--color-border)] bg-[var(--color-page)] px-3 py-2 font-mono text-xs leading-5 text-[var(--color-text-muted)]">
                  {rule.expression}
                </p>
              </article>
            ))
          )}
        </div>
      </div>
    </div>
  );
}

export default FuzzyRulesModal;
