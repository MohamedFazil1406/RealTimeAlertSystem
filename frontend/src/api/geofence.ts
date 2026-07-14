import client from "./client";

export const getGeofences = async () => {
  const response = await client.get("/geofences");
  return response.data;
};

export const createGeofence = async (data: any) => {
  const response = await client.post("/geofences", data);
  return response.data;
};
