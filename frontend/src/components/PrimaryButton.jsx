function PrimaryButton({ children, className = "", ...props }) {
  return (
    <button
      className={`rounded-md border border-[var(--color-border-hover)] bg-[var(--color-surface-raised)] px-4 py-2.5 text-sm font-medium text-[var(--color-text)] transition duration-200 hover:border-[var(--color-border-strong)] hover:bg-[var(--color-surface-hover)] focus:outline-none focus:ring-2 focus:ring-[var(--color-accent-soft)] disabled:opacity-50 ${className}`}
      {...props}
    >
      {children}
    </button>
  );
}

export default PrimaryButton;
