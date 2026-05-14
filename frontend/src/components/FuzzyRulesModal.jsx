import { X } from "lucide-react";

function FuzzyRulesModal({
  isOpen,
  onClose,
  rules,
  title = "Reglas de riego difuso",
  description = "Consulta las reglas cargadas para el sistema difuso de riego.",
}) {
  if (!isOpen) {
    return null;
  }

  function handlePanelClick(event) {
    event.stopPropagation();
  }

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 px-4 py-6"
      onClick={onClose}
    >
      <div
        className="flex max-h-[86vh] w-full max-w-3xl flex-col rounded-lg border border-zinc-800 bg-[#111111]"
        onClick={handlePanelClick}
      >
        <div className="flex items-start justify-between gap-4 border-b border-zinc-800 px-5 py-4">
          <div>
            <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-zinc-500">
              Base de reglas
            </p>
            <h2 className="mt-1 font-[var(--font-display)] text-xl font-semibold text-zinc-50">
              {title}
            </h2>
            <p className="mt-2 max-w-xl text-sm leading-6 text-zinc-400">
              {description}
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

        <div className="grid gap-2 overflow-y-auto px-5 py-5">
          {rules.map((rule) => (
            <article
              key={rule.id}
              className="rounded-md border border-zinc-800 bg-[#0d0d0d] px-3 py-3"
            >
              <p className="font-mono text-[11px] font-semibold text-zinc-500">{rule.id}</p>
              <p className="mt-2 whitespace-pre-line rounded-md border border-zinc-800 bg-[#090909] px-3 py-2 font-mono text-xs leading-5 text-zinc-400">
                {rule.expression}
              </p>
            </article>
          ))}
        </div>
      </div>
    </div>
  );
}

export default FuzzyRulesModal;
