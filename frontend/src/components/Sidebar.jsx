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
    <aside className="border-b border-[var(--color-border)] bg-[var(--color-page)] px-4 py-4 md:fixed md:inset-y-0 md:left-0 md:w-64 md:border-b-0 md:border-r">
      <div className="flex items-center gap-3 border-b border-[var(--color-border)] pb-5">
        <div className="flex h-9 w-9 items-center justify-center rounded-md border border-[var(--color-border)] bg-[var(--color-surface)]">
          <GraduationCap className="h-4 w-4 text-[var(--color-text-soft)]" strokeWidth={1.8} />
        </div>
        <div>
          <p className="font-[var(--font-display)] text-sm font-semibold text-[var(--color-text)]">
            Sistemas expertos
          </p>
          <p className="font-[var(--font-display)] text-[11px] font-medium uppercase tracking-[0.14em] text-[var(--color-text-subtle)]">
            Inteligencia artificial
          </p>
        </div>
      </div>

      <div className="mt-6">
        <p className="mb-2 px-2 font-[var(--font-display)] text-[11px] font-semibold uppercase tracking-[0.16em] text-[var(--color-text-faint)]">
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
                    ? "border-[var(--color-border-hover)] bg-[var(--color-surface-raised)] text-[var(--color-text)]"
                    : "border-transparent text-[var(--color-text-subtle)] hover:border-[var(--color-border)] hover:bg-[var(--color-surface)] hover:text-[var(--color-text-soft)]"
                }`}
              >
                <Icon
                  className={`h-4 w-4 shrink-0 transition duration-200 ${
                    isActive ? "text-[var(--color-accent)]" : "text-[var(--color-text-faint)] group-hover:text-[var(--color-text-muted)]"
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
