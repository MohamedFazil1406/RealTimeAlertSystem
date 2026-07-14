export interface Alert {
  event_type: string;

  vehicle_id: string;

  geofence_id: string;

  latitude: number;

  longitude: number;

  timestamp: string;
}
