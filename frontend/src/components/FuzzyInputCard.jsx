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
    <section className="rounded-lg border border-[var(--color-border)] bg-[var(--color-surface)]">
      <div className="flex flex-col gap-4 border-b border-[var(--color-border)] px-5 py-4 sm:flex-row sm:items-start sm:justify-between">
        <div>
          <p className="font-[var(--font-display)] text-[12px] font-semibold uppercase tracking-[0.16em] text-[var(--color-text-faint)]">
            Variable de entrada
          </p>
          <h2 className="mt-1 font-[var(--font-display)] text-lg font-semibold text-[var(--color-text)]">
            {variable.label}
          </h2>
        </div>
        <div className="rounded-md border border-[var(--color-border)] bg-[var(--color-page-panel)] px-3 py-2 text-right">
          <label
            htmlFor={`${variable.id}-value`}
            className="block font-[var(--font-display)] text-[12px] font-semibold uppercase tracking-[0.14em] text-[var(--color-text-faint)]"
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
              className="w-20 border-none bg-transparent p-0 text-right font-mono text-lg text-[var(--color-text)] outline-none transition placeholder:text-[var(--color-border-hover)] focus:text-[var(--color-text)]"
              aria-label={`Valor actual de ${variable.label}`}
            />
            {(variable.inputUnit || variable.unit) && (
              <span className="font-[var(--font-display)] text-sm font-semibold text-[var(--color-text-subtle)]">
                {variable.inputUnit || variable.unit}
              </span>
            )}
          </div>
          {variable.formatValue && (
            <p className="mt-1 font-mono text-[11px] text-[var(--color-text-subtle)]">
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
            className="h-2 w-full cursor-pointer appearance-none rounded-full bg-[var(--color-border)] accent-[var(--color-accent)]"
            aria-label={variable.label}
          />
          <div className="mt-2 flex justify-between font-mono text-lg text-[var(--color-text-faint)]">
            <span>{formatValue(variable.universe[0], variable.unit, variable.formatValue)}</span>
            <span>{formatValue(variable.universe[1], variable.unit, variable.formatValue)}</span>
          </div>
        </div>

        <div className="rounded-md border border-[var(--color-border)] bg-[var(--color-page-panel)] p-2">
          <MembershipChart variable={variable} value={value} />
        </div>
      </div>
    </section>
  );
}

export default FuzzyInputCard;
