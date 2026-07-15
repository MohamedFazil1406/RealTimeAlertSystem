import client from "./client";

export async function updateVehicleLocation(
  vehicleId: string,
  latitude: number,
  longitude: number,
) {
  const payload = {
    vehicle_id: vehicleId,
    latitude,
    longitude,
    timestamp: new Date().toISOString(),
  };

  console.log("Sending location update:", payload);

  const { data } = await client.post("/vehicles/location", payload);

  console.log("Location update response:", data);

  return data;
}
