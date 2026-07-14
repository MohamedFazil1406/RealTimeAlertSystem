import { useState } from "react";

type Props = {
  onSubmit: (data: any) => void;
};

export default function VehicleForm({ onSubmit }: Props) {
  const [form, setForm] = useState({
    vehicle_number: "",

    driver_name: "",

    vehicle_type: "",

    phone: "",
  });

  function handleChange(e: any) {
    setForm({
      ...form,

      [e.target.name]: e.target.value,
    });
  }

  return (
    <div className="bg-white p-6 rounded shadow">
      <h2 className="text-xl font-bold mb-5">Register Vehicle</h2>

      <div className="grid grid-cols-2 gap-4">
        <input
          name="vehicle_number"
          placeholder="Vehicle Number"
          className="border p-2 rounded"
          onChange={handleChange}
        />

        <input
          name="driver_name"
          placeholder="Driver Name"
          className="border p-2 rounded"
          onChange={handleChange}
        />

        <input
          name="vehicle_type"
          placeholder="Vehicle Type"
          className="border p-2 rounded"
          onChange={handleChange}
        />

        <input
          name="phone"
          placeholder="Phone"
          className="border p-2 rounded"
          onChange={handleChange}
        />
      </div>

      <button
        className="mt-5 bg-blue-600 text-white px-6 py-2 rounded"
        onClick={() => onSubmit(form)}
      >
        Register
      </button>
    </div>
  );
}
