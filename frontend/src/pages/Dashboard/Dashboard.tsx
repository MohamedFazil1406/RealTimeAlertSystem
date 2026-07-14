import { useEffect, useState } from "react";
import { getDashboardStats } from "../../api/dashboard";
import StatCard from "../../components/Dashboard/StatCard";

export default function Dashboard() {
  const [stats, setStats] = useState({
    vehicles: 0,
    geofences: 0,
    alerts: 0,
    violations: 0,
  });

  useEffect(() => {
    getDashboardStats().then(setStats);
  }, []);

  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">Dashboard</h1>

      <div className="grid grid-cols-4 gap-6">
        <StatCard title="Vehicles" value={stats.vehicles} />

        <StatCard title="Geofences" value={stats.geofences} />

        <StatCard title="Alerts" value={stats.alerts} />

        <StatCard title="Violations" value={stats.violations} />
      </div>
    </div>
  );
}
