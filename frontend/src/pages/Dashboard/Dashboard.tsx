import { useEffect, useState } from "react";

import { getDashboardStats } from "../../api/dashboard";
import { getGeofences } from "../../api/geofence";

import { type DashboardStats } from "../../types/dashboard";

import Card from "../../components/Dashboard/Card";

type Geofence = {
  ID: string;
  Name: string;
  Category: string;
  Status: string;
  Polygon: string;
};

export default function Dashboard() {
  const [stats, setStats] = useState<DashboardStats>({
    total_vehicles: 0,
    total_geofences: 0,
    total_violations: 0,
    total_alerts: 0,
  });

  const [geofences, setGeofences] = useState<Geofence[]>([]);

  useEffect(() => {
    loadData();
  }, []);

  async function loadData() {
    try {
      const [statsData, geofenceData] = await Promise.all([
        getDashboardStats(),
        getGeofences(),
      ]);

      setStats(statsData);
      setGeofences(geofenceData.geofences || []);
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div className="space-y-8">
      {/* Dashboard Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <Card title="Vehicles" value={stats.total_vehicles} />
        <Card title="Geofences" value={stats.total_geofences} />
        <Card title="Alerts" value={stats.total_alerts} />
        <Card title="Violations" value={stats.total_violations} />
      </div>

      {/* Geofence Table */}
      <div className="bg-white rounded-lg shadow p-5">
        <h2 className="text-xl font-semibold mb-4">Geofences</h2>

        <table className="w-full border-collapse">
          <thead>
            <tr className="bg-gray-100">
              <th className="border p-2 text-left">ID</th>
              <th className="border p-2 text-left">Name</th>
              <th className="border p-2 text-left">Category</th>

              <th className="border p-2 text-left">Coordinates</th>
            </tr>
          </thead>

          <tbody>
            {geofences.length === 0 ? (
              <tr>
                <td
                  colSpan={5}
                  className="border p-4 text-center text-gray-500"
                >
                  No geofences found
                </td>
              </tr>
            ) : (
              geofences.map((geofence) => {
                const polygon =
                  geofence.Polygon && geofence.Polygon !== "[]"
                    ? JSON.parse(geofence.Polygon)
                    : [];

                return (
                  <tr key={geofence.ID}>
                    <td className="border p-2">{geofence.ID}</td>
                    <td className="border p-2">{geofence.Name}</td>
                    <td className="border p-2">{geofence.Category}</td>

                    <td className="border p-2">
                      {polygon.length === 0 ? (
                        <span className="text-red-500">No Coordinates</span>
                      ) : (
                        polygon.map(
                          (
                            point: {
                              latitude: number;
                              longitude: number;
                            },
                            index: number,
                          ) => (
                            <div key={index}>
                              Lat: {point.latitude.toFixed(6)}
                              <br />
                              Lng: {point.longitude.toFixed(6)}
                            </div>
                          ),
                        )
                      )}
                    </td>
                  </tr>
                );
              })
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}
