import { type Violation } from "../../types/violation";

type Props = {
  violations: Violation[];
};

export default function ViolationTable({ violations }: Props) {
  if (!violations || violations.length === 0) {
    return (
      <div className="bg-white rounded shadow p-6 text-center text-gray-500">
        No violations found.
      </div>
    );
  }

  return (
    <table className="w-full bg-white rounded shadow">
      <thead>
        <tr className="bg-gray-200">
          <th className="p-3 text-left">Vehicle</th>
          <th className="p-3 text-left">Geofence</th>
          <th className="p-3 text-left">Event</th>
          <th className="p-3 text-left">Latitude</th>
          <th className="p-3 text-left">Longitude</th>
          <th className="p-3 text-left">Time</th>
        </tr>
      </thead>

      <tbody>
        {violations.map((v) => (
          <tr key={v.ID}>
            <td className="p-3">{v.VehicleID}</td>

            <td>{v.GeofenceID}</td>

            <td>
              <span
                className={`px-2 py-1 rounded text-white ${
                  v.EventType === "entry" ? "bg-green-600" : "bg-red-600"
                }`}
              >
                {v.EventType.toUpperCase()}
              </span>
            </td>

            <td>{v.Latitude}</td>

            <td>{v.Longitude}</td>

            <td>{new Date(v.Timestamp).toLocaleString()}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
