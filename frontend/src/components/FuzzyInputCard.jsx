import MembershipChart from "./MembershipChart.jsx";
import { clampValue, formatValue } from "../utils/membershipFunctions.js";

function FuzzyInputCard({ variable, value, onChange }) {
  function handleInputChange(event) {
    const numericValue = Number(event.target.value);

    if (event.target.value === "") {
      onChange(variable.universe[0]);
      return;
    }

    onChange(clampValue(numericValue, variable.universe));
  }

  return (
    <section className="rounded-lg border border-zinc-800 bg-[#111111]">
      <div className="flex flex-col gap-4 border-b border-zinc-800 px-5 py-4 sm:flex-row sm:items-start sm:justify-between">
        <div>
          <p className="font-[var(--font-display)] text-[12px] font-semibold uppercase tracking-[0.16em] text-zinc-600">
            Variable de entrada
          </p>
          <h2 className="mt-1 font-[var(--font-display)] text-lg font-semibold text-zinc-100">
            {variable.label}
          </h2>
        </div>
        <div className="rounded-md border border-zinc-800 bg-[#0d0d0d] px-3 py-2 text-right">
          <label
            htmlFor={`${variable.id}-value`}
            className="block font-[var(--font-display)] text-[12px] font-semibold uppercase tracking-[0.14em] text-zinc-600"
          >
            Valor actual
          </label>
          <div className="mt-1 flex items-baseline justify-end gap-2">
            <input
              id={`${variable.id}-value`}
              type="number"
              min={variable.universe[0]}
              max={variable.universe[1]}
              step={variable.step}
              value={value}
              onChange={handleInputChange}
              className="w-20 border-none bg-transparent p-0 text-right font-mono text-lg text-zinc-100 outline-none transition placeholder:text-zinc-700 focus:text-white"
              aria-label={`Valor actual de ${variable.label}`}
            />
            {(variable.inputUnit || variable.unit) && (
              <span className="font-[var(--font-display)] text-sm font-semibold text-zinc-500">
                {variable.inputUnit || variable.unit}
              </span>
            )}
          </div>
          {variable.formatValue && (
            <p className="mt-1 font-mono text-[11px] text-zinc-500">
              {formatValue(value, variable.unit, variable.formatValue)}
            </p>
          )}
        </div>
      </div>

      <div className="space-y-5 p-5">
        <div>
          <input
            type="range"
            min={variable.universe[0]}
            max={variable.universe[1]}
            step={variable.step}
            value={value}
            onChange={(event) => onChange(Number(event.target.value))}
            className="h-2 w-full cursor-pointer appearance-none rounded-full bg-zinc-800 accent-[#0070f3]"
            aria-label={variable.label}
          />
          <div className="mt-2 flex justify-between font-mono text-lg text-zinc-600">
            <span>{formatValue(variable.universe[0], variable.unit, variable.formatValue)}</span>
            <span>{formatValue(variable.universe[1], variable.unit, variable.formatValue)}</span>
          </div>
        </div>

        <div className="rounded-md border border-zinc-800 bg-[#0d0d0d] p-2">
          <MembershipChart variable={variable} value={value} />
        </div>
      </div>
    </section>
  );
}

export default FuzzyInputCard;
