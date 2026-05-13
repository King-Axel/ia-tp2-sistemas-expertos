import PageHeader from "../components/PageHeader.jsx";
import Panel from "../components/Panel.jsx";

function VentilationView() {
  return (
    <div className="space-y-8">
      <PageHeader
        title="Sistema de ventilación industrial"
        description="Se requiere automatizar el sistema de extracción de aire y ventilación. La velocidad del motor del extractor debe regularse de forma suave y continua para mantener un ambiente seguro y optimizar el consumo eléctrico. El sistema difuso debe tomar decisiones basadas en dos sensores principales: temperatura del ambiente en la planta y nivel de concentración de gases o humo."
      />

      <Panel title="Módulo difuso" eyebrow="Ventilación industrial">
        <p className="max-w-2xl text-sm leading-6 text-zinc-500">
          Esta sección queda preparada para incorporar la lógica del sistema.
        </p>
      </Panel>
    </div>
  );
}

export default VentilationView;
