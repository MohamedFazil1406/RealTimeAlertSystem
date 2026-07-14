import { MapContainer, TileLayer, Polygon, useMapEvents } from "react-leaflet";

import { useState } from "react";

type Props = {
  polygons: any[];

  onPolygonComplete: (points: any[]) => void;
};

function DrawPolygon({ onPolygonComplete }: any) {
  const [points, setPoints] = useState<any[]>([]);

  useMapEvents({
    click(e) {
      const updated = [
        ...points,

        {
          latitude: e.latlng.lat,
          longitude: e.latlng.lng,
        },
      ];

      setPoints(updated);

      if (updated.length >= 4) {
        onPolygonComplete(updated);
      }
    },
  });

  return points.length > 2 ? (
    <Polygon positions={points.map((p) => [p.latitude, p.longitude])} />
  ) : null;
}

export default function GeofenceMap({ polygons, onPolygonComplete }: Props) {
  return (
    <MapContainer
      center={[12.9716, 77.5946]}
      zoom={13}
      style={{ height: "600px" }}
    >
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />

      <DrawPolygon onPolygonComplete={onPolygonComplete} />

      {polygons.map((polygon, index) => (
        <Polygon key={index} positions={polygon} />
      ))}
    </MapContainer>
  );
}
