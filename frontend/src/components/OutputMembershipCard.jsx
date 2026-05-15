import MembershipChart from "./MembershipChart.jsx";

function OutputMembershipCard({ variable }) {
  return (
    <section className="rounded-lg border border-[var(--color-border)] bg-[var(--color-surface)]">
      <div className="border-b border-[var(--color-border)] px-5 py-4">
        <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-[var(--color-text-faint)]">
          Variable de salida
        </p>
        <h2 className="mt-1 font-[var(--font-display)] text-base font-semibold text-[var(--color-text)]">
          {variable.label}
        </h2>
      </div>

      <div className="p-5">
        <div className="rounded-md border border-[var(--color-border)] bg-[var(--color-page-panel)] p-2">
          <MembershipChart variable={variable} showSelectedValue={false} />
        </div>
      </div>
    </section>
  );
}

export default OutputMembershipCard;
