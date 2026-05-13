import { useState } from "react";
import AppLayout from "./layouts/AppLayout.jsx";
import DiagnosisView from "./views/DiagnosisView.jsx";
import IrrigationView from "./views/IrrigationView.jsx";
import VentilationView from "./views/VentilationView.jsx";

const views = {
  diagnosis: DiagnosisView,
  irrigation: IrrigationView,
  ventilation: VentilationView,
};

function App() {
  const [activeView, setActiveView] = useState("diagnosis");
  const ActiveView = views[activeView];

  return (
    <AppLayout activeView={activeView} onViewChange={setActiveView}>
      <ActiveView />
    </AppLayout>
  );
}

export default App;
