function PageHeader({ title, description }) {
  return (
    <header className="border-b border-[var(--color-border)] pb-8">
      <p className="mb-3 font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.18em] text-[var(--color-text-subtle)]">
        Prototipo universitario
      </p>
      <h1 className="max-w-4xl font-[var(--font-display)] text-3xl font-semibold tracking-tight text-[var(--color-text)] md:text-[34px] md:leading-tight">
        {title}
      </h1>
      <p className="mt-4 max-w-3xl text-sm leading-6 text-[var(--color-text-muted)]">
        {description}
      </p>
    </header>
  );
}

export default PageHeader;
