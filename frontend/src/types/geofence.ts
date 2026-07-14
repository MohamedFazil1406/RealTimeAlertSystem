export interface Coordinate {
  latitude: number;
  longitude: number;
}

export interface Geofence {
  id: string;
  name: string;
  description: string;
  category: string;
  polygon: string;
}
