function CheckboxField({ name, label, checked, onChange }) {
  return (
    <label className="group flex cursor-pointer items-center gap-3 rounded-md border border-zinc-800 bg-[#0d0d0d] px-3 py-2.5 text-sm text-zinc-300 transition duration-200 hover:border-zinc-700 hover:bg-[#151515]">
      <input
        type="checkbox"
        name={name}
        checked={checked}
        onChange={onChange}
        className="h-4 w-4 rounded border-zinc-700 bg-[#090909] accent-[#0070f3] outline-none transition focus:ring-2 focus:ring-[#0070f3]/30"
      />
      <span className="font-[var(--font-display)] text-[13px] font-medium text-zinc-400 transition group-hover:text-zinc-200">
        {label}
      </span>
    </label>
  );
}

export default CheckboxField;
