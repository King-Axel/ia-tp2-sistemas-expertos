import MembershipChart from "./MembershipChart.jsx";

function OutputMembershipCard({ variable }) {
  return (
    <section className="rounded-lg border border-zinc-800 bg-[#111111]">
      <div className="border-b border-zinc-800 px-5 py-4">
        <p className="font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-zinc-600">
          Variable de salida
        </p>
        <h2 className="mt-1 font-[var(--font-display)] text-base font-semibold text-zinc-100">
          {variable.label}
        </h2>
      </div>

      <div className="p-5">
        <div className="rounded-md border border-zinc-800 bg-[#0d0d0d] p-2">
          <MembershipChart variable={variable} showSelectedValue={false} />
        </div>
      </div>
    </section>
  );
}

export default OutputMembershipCard;
