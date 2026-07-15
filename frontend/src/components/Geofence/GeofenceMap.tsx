import {
  MapContainer,
  TileLayer,
  Polygon,
  Marker,
  Popup,
  useMapEvents,
} from "react-leaflet";
import { useState } from "react";

import { type Vehicle } from "../../types/vehicle";

type DrawPoint = {
  latitude: number;
  longitude: number;
};

type Props = {
  polygons: [number, number][][];
  vehicles: Vehicle[];
  onPolygonComplete: (points: DrawPoint[]) => void;
  onVehicleMove: (
    vehicleId: string,
    latitude: number,
    longitude: number,
  ) => void;
};

function DrawPolygon({
  onPolygonComplete,
}: {
  onPolygonComplete: (points: DrawPoint[]) => void;
}) {
  const [points, setPoints] = useState<DrawPoint[]>([]);

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

  if (points.length < 3) return null;

  return (
    <Polygon
      positions={points.map((p) => [p.latitude, p.longitude])}
      pathOptions={{
        color: "blue",
        weight: 2,
        fillOpacity: 0.2,
      }}
    />
  );
}

export default function GeofenceMap({
  polygons,
  vehicles,
  onPolygonComplete,
  onVehicleMove,
}: Props) {
  return (
    <MapContainer
      center={[11.24, 78.15]}
      zoom={14}
      style={{
        height: "600px",
        width: "100%",
      }}
    >
      <TileLayer
        attribution="© OpenStreetMap contributors"
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />

      <DrawPolygon onPolygonComplete={onPolygonComplete} />

      {/* Existing Geofences */}
      {polygons.map((polygon, index) => (
        <Polygon
          key={index}
          positions={polygon}
          pathOptions={{
            color: "red",
            weight: 3,
            fillOpacity: 0.3,
          }}
        />
      ))}

      {/* Vehicle Markers */}
      {vehicles.map((vehicle) => {
        if (vehicle.latitude == null || vehicle.longitude == null) {
          return null;
        }

        return (
          <Marker
            key={vehicle.id}
            position={[vehicle.latitude, vehicle.longitude]}
            draggable
            eventHandlers={{
              dragend: (event) => {
                const marker = event.target;
                const position = marker.getLatLng();
                console.log("Dragged to:", position);
                onVehicleMove(vehicle.id, position.lat, position.lng);
              },
            }}
          >
            <Popup>
              <div>
                <strong>{vehicle.vehicle_number}</strong>
                <br />
                Driver: {vehicle.driver_name}
                <br />
                Type: {vehicle.vehicle_type}
                <br />
                Status: {vehicle.status}
              </div>
            </Popup>
          </Marker>
        );
      })}
    </MapContainer>
  );
}
