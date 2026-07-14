import { useState } from "react";

type Props = {
  onSave: (data: any) => void;
};

export default function GeofenceForm({ onSave }: Props) {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("delivery_zone");

  return (
    <div className="bg-white p-5 rounded shadow mb-5">
      <h2 className="text-xl font-bold mb-4">Create Geofence</h2>

      <input
        className="border p-2 w-full mb-3"
        placeholder="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />

      <textarea
        className="border p-2 w-full mb-3"
        placeholder="Description"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
      />

      <select
        className="border p-2 w-full"
        value={category}
        onChange={(e) => setCategory(e.target.value)}
      >
        <option value="delivery_zone">Delivery Zone</option>

        <option value="restricted_zone">Restricted Zone</option>

        <option value="customer_area">Customer Area</option>

        <option value="toll_zone">Toll Zone</option>
      </select>

      <button
        className="mt-4 bg-blue-600 text-white px-5 py-2 rounded"
        onClick={() =>
          onSave({
            name,
            description,
            category,
          })
        }
      >
        Save Geofence
      </button>
    </div>
  );
}
