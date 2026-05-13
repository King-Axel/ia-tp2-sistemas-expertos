import { Activity, Droplets, Fan, GraduationCap } from "lucide-react";

const navigationItems = [
  {
    id: "diagnosis",
    label: "Sistema de diagnóstico",
    icon: Activity,
  },
  {
    id: "irrigation",
    label: "Sistema de riego automático",
    icon: Droplets,
  },
  {
    id: "ventilation",
    label: "Sistema de ventilación industrial",
    icon: Fan,
  },
];

function Sidebar({ activeView, onViewChange }) {
  return (
    <aside className="border-b border-zinc-800 bg-[#090909] px-4 py-4 md:fixed md:inset-y-0 md:left-0 md:w-64 md:border-b-0 md:border-r">
      <div className="flex items-center gap-3 border-b border-zinc-800 pb-5">
        <div className="flex h-9 w-9 items-center justify-center rounded-md border border-zinc-800 bg-[#111111]">
          <GraduationCap className="h-4 w-4 text-zinc-200" strokeWidth={1.8} />
        </div>
        <div>
          <p className="font-[var(--font-display)] text-sm font-semibold text-zinc-100">
            Sistemas expertos
          </p>
          <p className="font-[var(--font-display)] text-[11px] font-medium uppercase tracking-[0.14em] text-zinc-500">
            Inteligencia artificial
          </p>
        </div>
      </div>

      <div className="mt-6">
        <p className="mb-2 px-2 font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-zinc-600">
          Módulos
        </p>
        <nav className="grid gap-1">
          {navigationItems.map((item) => {
            const Icon = item.icon;
            const isActive = activeView === item.id;

            return (
              <button
                key={item.id}
                type="button"
                onClick={() => onViewChange(item.id)}
                className={`group flex w-full items-center gap-3 rounded-md border px-2.5 py-2.5 text-left text-sm transition duration-200 ${
                  isActive
                    ? "border-zinc-700 bg-[#171717] text-white"
                    : "border-transparent text-zinc-500 hover:border-zinc-800 hover:bg-[#111111] hover:text-zinc-200"
                }`}
              >
                <Icon
                  className={`h-4 w-4 shrink-0 transition duration-200 ${
                    isActive ? "text-[#0070f3]" : "text-zinc-600 group-hover:text-zinc-300"
                  }`}
                  strokeWidth={1.8}
                />
                <span className="font-medium">{item.label}</span>
              </button>
            );
          })}
        </nav>
      </div>
    </aside>
  );
}

export default Sidebar;
