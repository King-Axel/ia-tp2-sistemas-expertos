function OutputPlaceholder() {
  return (
    <section className="rounded-lg border border-[var(--color-border)] bg-[var(--color-surface)]">
      <div className="border-b border-[var(--color-border)] px-5 py-4">
        <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-[var(--color-text-faint)]">
          Variable de salida
        </p>
        <h2 className="mt-1 font-[var(--font-display)] text-base font-semibold text-[var(--color-text)]">
          Porcentaje de apertura de válvula
        </h2>
      </div>
      <div className="p-5">
        <div className="flex min-h-64 items-center justify-center rounded-md border border-dashed border-[var(--color-border)] bg-[var(--color-page-panel)] px-6 text-center">
          <p className="max-w-sm text-sm leading-6 text-[var(--color-text-subtle)]">
            Los conjuntos difusos de salida se agregarán más adelante.
          </p>
        </div>
      </div>
    </section>
  );
}

export default OutputPlaceholder;
