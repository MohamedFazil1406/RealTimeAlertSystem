import { createContext, useState } from "react";
import { type Alert } from "../types/alert";

type ContextType = {
  alerts: Alert[];

  addAlert: (alert: Alert) => void;
};

export const AlertContext = createContext<ContextType>(null!);

export function AlertProvider({ children }: any) {
  const [alerts, setAlerts] = useState<Alert[]>([]);

  function addAlert(alert: Alert) {
    setAlerts((prev) => [alert, ...prev]);
  }

  return (
    <AlertContext.Provider
      value={{
        alerts,
        addAlert,
      }}
    >
      {children}
    </AlertContext.Provider>
  );
}
