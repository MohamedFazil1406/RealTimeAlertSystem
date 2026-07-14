import { Routes, Route } from "react-router-dom";

import DashboardLayout from "./layouts/DashboardLayout";
import { Toaster } from "react-hot-toast";

import Dashboard from "./pages/Dashboard/Dashboard";
import Vehicles from "./pages/Vehicles/Vehicles";
import Geofences from "./pages/Geofences/Geofences";
import Alerts from "./pages/Alerts/Alerts";
import Violations from "./pages/Violations/Violations";

export default function App() {
  return (
    <>
      <Toaster position="top-right" />
      <Routes>
        <Route element={<DashboardLayout />}>
          <Route path="/" element={<Dashboard />} />

          <Route path="/vehicles" element={<Vehicles />} />

          <Route path="/geofences" element={<Geofences />} />

          <Route path="/alerts" element={<Alerts />} />

          <Route path="/violations" element={<Violations />} />
        </Route>
      </Routes>
    </>
  );
}
