import { createContext, useState, type ReactNode } from "react";
import { type Vehicle } from "../types/vehicle";

type VehicleContextType = {
  vehicles: Vehicle[];
  setVehicles: React.Dispatch<React.SetStateAction<Vehicle[]>>;
};

export const VehicleContext = createContext<VehicleContextType>(null!);

export function VehicleProvider({ children }: { children: ReactNode }) {
  const [vehicles, setVehicles] = useState<Vehicle[]>([]);

  return (
    <VehicleContext.Provider value={{ vehicles, setVehicles }}>
      {children}
    </VehicleContext.Provider>
  );
}
