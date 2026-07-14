import { useState } from "react";

type Props = {
  onFilter: (filters: any) => void;
};

export default function ViolationFilter({ onFilter }: Props) {
  const [vehicle, setVehicle] = useState("");

  const [geofence, setGeofence] = useState("");

  const [startDate, setStartDate] = useState("");

  const [endDate, setEndDate] = useState("");

  return (
    <div className="bg-white p-5 rounded shadow mb-5">
      <div className="grid grid-cols-4 gap-4">
        <input
          placeholder="Vehicle ID"
          className="border p-2"
          value={vehicle}
          onChange={(e) => setVehicle(e.target.value)}
        />

        <input
          placeholder="Geofence ID"
          className="border p-2"
          value={geofence}
          onChange={(e) => setGeofence(e.target.value)}
        />

        <input
          type="date"
          className="border p-2"
          value={startDate}
          onChange={(e) => setStartDate(e.target.value)}
        />

        <input
          type="date"
          className="border p-2"
          value={endDate}
          onChange={(e) => setEndDate(e.target.value)}
        />
      </div>

      <button
        className="mt-4 bg-blue-600 text-white px-5 py-2 rounded"
        onClick={() =>
          onFilter({
            vehicle_id: vehicle,

            geofence_id: geofence,

            start_date: startDate,

            end_date: endDate,
          })
        }
      >
        Apply Filters
      </button>
    </div>
  );
}
