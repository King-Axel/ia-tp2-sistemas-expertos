function Panel({ children, title, eyebrow, className = "" }) {
  return (
    <section
      className={`rounded-lg border border-[var(--color-border)] bg-[var(--color-surface)] ${className}`}
    >
      {(title || eyebrow) && (
        <div className="border-b border-[var(--color-border)] px-5 py-4 md:px-6">
          {eyebrow && (
            <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-[var(--color-text-faint)]">
              {eyebrow}
            </p>
          )}
          {title && (
            <h2 className="mt-1 font-[var(--font-display)] text-base font-semibold text-[var(--color-text)]">
              {title}
            </h2>
          )}
        </div>
      )}
      <div className="p-5 md:p-6">{children}</div>
    </section>
  );
}

export default Panel;
