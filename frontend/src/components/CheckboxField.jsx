function CheckboxField({ name, label, checked, onChange }) {
  return (
    <label className="group flex cursor-pointer items-center gap-3 rounded-md border border-[var(--color-border)] bg-[var(--color-page-panel)] px-3 py-2.5 text-sm text-[var(--color-text-muted)] transition duration-200 hover:border-[var(--color-border-hover)] hover:bg-[var(--color-surface-soft)]">
      <input
        type="checkbox"
        name={name}
        checked={checked}
        onChange={onChange}
        className="h-4 w-4 rounded border-[var(--color-border-hover)] bg-[var(--color-page)] accent-[var(--color-accent)] outline-none transition focus:ring-2 focus:ring-[var(--color-accent-muted)]"
      />
      <span className="font-[var(--font-display)] text-[13px] font-medium text-[var(--color-text-muted)] transition group-hover:text-[var(--color-text-soft)]">
        {label}
      </span>
    </label>
  );
}

export default CheckboxField;
