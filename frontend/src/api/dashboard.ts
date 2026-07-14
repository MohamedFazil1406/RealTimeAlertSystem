import { getVehicles } from "./vehicle";
import { getGeofences } from "./geofence";
import { getViolations } from "./violation";
import client from "./client";

export async function getDashboardStats() {
  const [vehicles, geofences, violations, alerts] = await Promise.all([
    getVehicles(),
    getGeofences(),
    getViolations(),
    client.get("/alerts"),
  ]);

  return {
    vehicles: vehicles.vehicles.length,
    geofences: geofences.geofences.length,
    violations: violations.total_count,
    alerts: alerts.data.alerts.length,
  };
}
