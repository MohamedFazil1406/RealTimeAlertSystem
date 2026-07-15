import { useContext, useEffect } from "react";
import toast from "react-hot-toast";

import { AlertContext } from "../context/AlertContext";
import { VehicleContext } from "../context/VehicleContext";

export default function useWebSocket() {
  const { addAlert } = useContext(AlertContext);

  const { setVehicles } = useContext(VehicleContext);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws/alerts");

    socket.onopen = () => {
      console.log("WebSocket Connected");
    };

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);

      // Live vehicle location updates
      if (data.type === "location") {
        setVehicles((prev) =>
          prev.map((vehicle) =>
            vehicle.id === data.vehicle_id
              ? {
                  ...vehicle,
                  latitude: data.latitude,
                  longitude: data.longitude,
                }
              : vehicle,
          ),
        );

        return;
      }

      // Entry / Exit alerts
      addAlert(data);

      toast.success(`${data.event_type.toUpperCase()} detected`);
    };

    socket.onclose = () => {
      console.log("WebSocket Disconnected");
    };

    socket.onerror = (error) => {
      console.error("WebSocket Error:", error);
    };

    return () => {
      socket.close();
    };
  }, [addAlert, setVehicles]);
}
