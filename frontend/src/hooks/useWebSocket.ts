import { useContext, useEffect } from "react";

import toast from "react-hot-toast";

import { AlertContext } from "../context/AlertContext";

export default function useWebSocket() {
  const { addAlert } = useContext(AlertContext);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws/alerts");

    socket.onmessage = (event) => {
      const alert = JSON.parse(event.data);

      addAlert(alert);

      toast.success(`${alert.event_type.toUpperCase()} detected`);
    };

    return () => {
      socket.close();
    };
  }, []);
}
