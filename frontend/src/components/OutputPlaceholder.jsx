function OutputPlaceholder() {
  return (
    <section className="rounded-lg border border-zinc-800 bg-[#111111]">
      <div className="border-b border-zinc-800 px-5 py-4">
        <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-zinc-600">
          Variable de salida
        </p>
        <h2 className="mt-1 font-[var(--font-display)] text-base font-semibold text-zinc-100">
          Porcentaje de apertura de válvula
        </h2>
      </div>
      <div className="p-5">
        <div className="flex min-h-64 items-center justify-center rounded-md border border-dashed border-zinc-800 bg-[#0d0d0d] px-6 text-center">
          <p className="max-w-sm text-sm leading-6 text-zinc-500">
            Los conjuntos difusos de salida se agregarán más adelante.
          </p>
        </div>
      </div>
    </section>
  );
}

export default OutputPlaceholder;
