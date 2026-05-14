import { CheckCircle2, X } from "lucide-react";

const resultLabels = {
  compatible_con_dengue: "Compatible con dengue",
  compatible_con_covid: "Compatible con COVID-19",
  no_puede_concluir: "Sin conclusión",
};

function InferenceResultModal({ isOpen, onClose, result, activatedRules }) {
  if (!isOpen) {
    return null;
  }

  const resultLabel = resultLabels[result] ?? "Sin conclusión";

  function handlePanelClick(event) {
    event.stopPropagation();
  }

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 px-4 py-6"
      onClick={onClose}
    >
      <div
        className="scrollbar-dark flex-col max-h-[calc(100vh-3rem)] overflow-y-auto w-full max-w-2xl rounded-lg border border-zinc-800 bg-[#111111]"
        onClick={handlePanelClick}
      >
        <div className="flex items-start justify-between gap-4 border-b border-zinc-800 px-5 py-4">
          <div>
            <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-zinc-500">
              Resultado de inferencia
            </p>
            <h2 className="mt-1 font-[var(--font-display)] text-xl font-semibold text-zinc-50">
              {resultLabel}
            </h2>
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

        <div className="space-y-5 px-5 py-5">
          <div className="rounded-md border border-zinc-800 bg-[#0d0d0d] p-4">
            <p className="text-sm leading-6 text-zinc-400">
              Resultado calculado por el motor SEBR con la base de hechos enviada.
            </p>
          </div>

          <div>
            <p className="mb-3 font-[var(--font-display)] text-xs font-semibold uppercase tracking-[0.14em] text-zinc-500">
              Reglas activadas
            </p>
            <div className="grid gap-2">
              {activatedRules.length > 0 ? (
                activatedRules.map((rule) => (
                <article
                  key={rule.id}
                  className="rounded-md border border-zinc-800 bg-[#0d0d0d] px-3 py-3"
                >
                  <div className="flex items-center gap-2">
                    <CheckCircle2
                      className="h-4 w-4 shrink-0 text-[#0070f3]"
                      strokeWidth={1.8}
                    />
                    <p className="font-[var(--font-display)] text-[13px] font-semibold text-zinc-200">
                      {rule.id} · {rule.name}
                    </p>
                  </div>
                  <p className="mt-2 font-mono text-xs leading-5 text-zinc-500">
                    {rule.expression}
                  </p>
                </article>
                ))
              ) : (
                <p className="rounded-md border border-zinc-800 bg-[#0d0d0d] px-3 py-3 text-sm text-zinc-500">
                  No se activaron reglas.
                </p>
              )}
            </div>
          </div>
        </div>

        <div className="flex justify-end border-t border-zinc-800 px-5 py-4">
          <button
            type="button"
            onClick={onClose}
            className="rounded-md border border-zinc-700 bg-[#171717] px-4 py-2 text-sm font-medium text-zinc-100 transition duration-200 hover:border-zinc-600 hover:bg-[#1f1f1f] focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35"
          >
            Cerrar
          </button>
        </div>
      </div>
    </div>
  );
}

export default InferenceResultModal;
