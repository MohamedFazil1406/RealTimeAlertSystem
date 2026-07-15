import { useEffect, useState } from "react";

import { getDashboardStats } from "../../api/dashboard";

import { type DashboardStats } from "../../types/dashboard";

import Card from "../../components/Dashboard/Card";

export default function Dashboard() {
  const [stats, setStats] = useState<DashboardStats>({
    total_vehicles: 0,
    total_geofences: 0,
    total_violations: 0,
    total_alerts: 0,
  });

  useEffect(() => {
    loadStats();
  }, []);

  async function loadStats() {
    try {
      const data = await getDashboardStats();
      setStats(data);
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div className="grid grid-cols-4 gap-4">
      <Card title="Vehicles" value={stats.total_vehicles} />

      <Card title="Geofences" value={stats.total_geofences} />

      <Card title="Alerts" value={stats.total_alerts} />

      <Card title="Violations" value={stats.total_violations} />
    </div>
  );
}
