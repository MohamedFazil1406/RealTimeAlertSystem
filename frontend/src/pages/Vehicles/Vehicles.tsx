import { useEffect, useState } from "react";

import { createVehicle, getVehicles } from "../../api/vehicle";

import VehicleForm from "../../components/Vehicle/VehicleForm";

import VehicleTable from "../../components/Vehicle/VehicleTable";

import { type Vehicle } from "../../types/vehicle";

export default function Vehicles() {
  const [vehicles, setVehicles] = useState<Vehicle[]>([]);

  async function loadVehicles() {
    const data = await getVehicles();

    console.log("this is data", data);
    console.log("First Vehicle:", JSON.stringify(data.vehicles[0], null, 2));

    setVehicles(data.vehicles);
  }

  async function handleCreate(form: any) {
    await createVehicle(form);

    loadVehicles();
  }

  useEffect(() => {
    loadVehicles();
  }, []);

  return (
    <div>
      <VehicleForm onSubmit={handleCreate} />

      <VehicleTable vehicles={vehicles} />
    </div>
  );
}
