import { useContext } from "react";

import { AlertContext } from "../../context/AlertContext";

export default function AlertFeed() {
  const { alerts } = useContext(AlertContext);

  return (
    <div className="bg-white rounded shadow p-4">
      <h2 className="text-xl font-bold mb-4">Live Alerts</h2>

      {alerts.length === 0 && <p>No alerts</p>}

      {alerts.map((alert, index) => (
        <div key={index} className="border-b py-3">
          <div>
            <strong>{alert.event_type.toUpperCase()}</strong>
          </div>

          <div>
            Vehicle
            {alert.vehicle_id}
          </div>

          <div>{alert.timestamp}</div>
        </div>
      ))}
    </div>
  );
}
