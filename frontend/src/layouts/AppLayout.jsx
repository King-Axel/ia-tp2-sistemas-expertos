import Sidebar from "../components/Sidebar.jsx";

function AppLayout({ activeView, onViewChange, children }) {
  return (
    <div className="min-h-screen bg-[#090909] text-zinc-50">
      <Sidebar activeView={activeView} onViewChange={onViewChange} />
      <main className="min-h-screen px-4 py-5 transition-all duration-200 md:ml-64 md:px-8 md:py-8 lg:px-10">
        <div className="mx-auto max-w-6xl">{children}</div>
      </main>
    </div>
  );
}

export default AppLayout;
