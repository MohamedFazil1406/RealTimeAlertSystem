import { useContext, useEffect, useState } from "react";

import GeofenceForm from "../../components/Geofence/GeofenceForm";
import GeofenceMap from "../../components/Geofence/GeofenceMap";

import { createGeofence, getGeofences } from "../../api/geofence";
import { getVehicles } from "../../api/vehicle";

import { VehicleContext } from "../../context/VehicleContext";
import { updateVehicleLocation } from "../../api/location";

import toast from "react-hot-toast";

type DrawPoint = {
  latitude: number;
  longitude: number;
};

export default function Geofences() {
  const [points, setPoints] = useState<DrawPoint[]>([]);
  const [polygons, setPolygons] = useState<[number, number][][]>([]);

  const { vehicles, setVehicles } = useContext(VehicleContext);

  const load = async () => {
    try {
      const [geoResponse, vehicleResponse] = await Promise.all([
        getGeofences(),
        getVehicles(),
      ]);

      setVehicles(vehicleResponse.vehicles ?? []);

      const parsedPolygons: [number, number][][] = (geoResponse.geofences ?? [])
        .filter(
          (geofence: any) => geofence.Polygon && geofence.Polygon !== "[]",
        )
        .map((geofence: any) => {
          const coordinates = JSON.parse(geofence.Polygon);

          return coordinates.map(
            (point: { latitude: number; longitude: number }) =>
              [point.latitude, point.longitude] as [number, number],
          );
        });

      setPolygons(parsedPolygons);
    } catch (error) {
      console.error("Error loading geofences:", error);
    }
  };

  async function handleSave(form: any) {
    try {
      if (points.length < 4) {
        toast.error("Please draw a geofence with at least 4 points.");
        return;
      }

      await createGeofence({
        ...form,
        coordinates: points,
      });

      toast.success("Geofence created successfully!");

      setPoints([]);

      await load();
    } catch (error) {
      toast.error("Failed to create geofence.");
      console.error(error);
    }
  }

  async function handleVehicleMove(
    vehicleId: string,
    latitude: number,
    longitude: number,
  ) {
    try {
      await updateVehicleLocation(vehicleId, latitude, longitude);

      // Refresh vehicles and geofences
      await load();
    } catch (error) {
      console.error("Failed to update vehicle location:", error);
    }
  }

  useEffect(() => {
    load();

    // Temporary polling.
    // Remove this after WebSocket location updates are working.
    const interval = setInterval(load, 3000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div className="space-y-6">
      <GeofenceForm onSave={handleSave} />

      <GeofenceMap
        polygons={polygons}
        vehicles={vehicles}
        onPolygonComplete={setPoints}
        onVehicleMove={handleVehicleMove}
      />
    </div>
  );
}
