import { useEffect, useState } from "react";

import GeofenceForm from "../../components/Geofence/GeofenceForm";

import GeofenceMap from "../../components/Geofence/GeofenceMap";

import { createGeofence, getGeofences } from "../../api/geofence";

export default function Geofences() {
  const [points, setPoints] = useState<any[]>([]);

  const [polygons, setPolygons] = useState<any[]>([]);

  async function load() {
    const data = await getGeofences();

    console.log("this is data", data);

    setPolygons(
      data.geofences
        .filter((g: any) => g.Polygon && g.Polygon !== "[]")
        .map((g: any) => {
          const points = JSON.parse(g.Polygon);

          return points.map((p: any) => [p.latitude, p.longitude]);
        }),
    );
  }

  async function handleSave(form: any) {
    await createGeofence({
      ...form,

      coordinates: points,
    });

    load();
  }

  useEffect(() => {
    load();
  }, []);

  return (
    <div>
      <GeofenceForm onSave={handleSave} />

      <GeofenceMap polygons={polygons} onPolygonComplete={setPoints} />
    </div>
  );
}
