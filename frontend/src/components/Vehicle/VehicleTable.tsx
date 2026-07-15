import { type Vehicle } from "../../types/vehicle";

type Props = {
  vehicles: Vehicle[];
};

export default function VehicleTable({ vehicles }: Props) {
  return (
    <table className="w-full mt-8 bg-white rounded shadow">
      <thead>
        <tr className="bg-gray-200">
          <th className="p-3">ID</th>
          <th>Vehicle</th>
          <th>Driver</th>
          <th>Type</th>
          <th>Phone</th>
        </tr>
      </thead>

      <tbody>
        {vehicles.map((vehicle) => (
          <tr key={vehicle.id} className="border-b hover:bg-gray-50">
            <td className="p-3">{vehicle.id}</td>
            <td>{vehicle.vehicle_number}</td>

            <td>{vehicle.driver_name}</td>

            <td>{vehicle.vehicle_type}</td>

            <td>{vehicle.phone}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
