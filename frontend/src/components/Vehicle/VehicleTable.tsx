import { type Vehicle } from "../../types/vehicle";

type Props = {
  vehicles: Vehicle[];
};

export default function VehicleTable({ vehicles }: Props) {
  return (
    <table className="w-full mt-8 bg-white rounded shadow">
      <thead>
        <tr className="bg-gray-200">
          <th className="p-3">Vehicle</th>

          <th>Driver</th>

          <th>Type</th>

          <th>Phone</th>

          <th>Status</th>
        </tr>
      </thead>

      <tbody>
        {vehicles.map((vehicle: any) => (
          <tr key={vehicle.ID}>
            <td>{vehicle.VehicleNumber}</td>
            <td>{vehicle.DriverName}</td>
            <td>{vehicle.VehicleType}</td>
            <td>{vehicle.Phone}</td>
            <td>{vehicle.Status}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
