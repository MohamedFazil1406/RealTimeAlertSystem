export default function Dashboard() {
  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">Dashboard</h1>

      <div className="grid grid-cols-4 gap-6">
        <div className="bg-white rounded-lg shadow p-6">
          <h2 className="text-gray-500">Vehicles</h2>

          <p className="text-3xl font-bold">0</p>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <h2 className="text-gray-500">Geofences</h2>

          <p className="text-3xl font-bold">0</p>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <h2 className="text-gray-500">Alerts</h2>

          <p className="text-3xl font-bold">0</p>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <h2 className="text-gray-500">Violations</h2>

          <p className="text-3xl font-bold">0</p>
        </div>
      </div>
    </div>
  );
}
