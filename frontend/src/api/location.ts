import client from "./client";

export async function updateVehicleLocation(
  vehicleId: string,
  latitude: number,
  longitude: number,
) {
  const { data } = await client.post("/vehicles/location", {
    vehicle_id: vehicleId,
    latitude,
    longitude,
    timestamp: new Date().toISOString(),
  });

  return data;
}
