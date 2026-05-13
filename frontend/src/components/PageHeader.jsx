function PageHeader({ title, description }) {
  return (
    <header className="border-b border-zinc-800 pb-8">
      <p className="mb-3 font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.18em] text-zinc-500">
        Prototipo universitario
      </p>
      <h1 className="max-w-4xl font-[var(--font-display)] text-3xl font-semibold tracking-tight text-zinc-50 md:text-[34px] md:leading-tight">
        {title}
      </h1>
      <p className="mt-4 max-w-3xl text-sm leading-6 text-zinc-400">
        {description}
      </p>
    </header>
  );
}

export default PageHeader;
