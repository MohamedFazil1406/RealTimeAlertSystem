import { toast } from "react-hot-toast";
import type { ViolationFilters } from "../../pages/Violations/Violations";

type Props = {
  filters: ViolationFilters;
  onFilter: (filters: ViolationFilters) => Promise<void>;
};

export default function ViolationFilter({ filters, onFilter }: Props) {
  const handleApplyFilters = async () => {
    try {
      await onFilter(filters);
      toast.success("Filters applied successfully!");
    } catch (error) {
      console.error(error);
      toast.error("Failed to apply filters.");
    }
  };

  return (
    <div className="bg-white p-5 rounded shadow mb-5">
      <div className="grid grid-cols-4 gap-4">
        <input
          placeholder="Vehicle ID"
          className="border p-2 rounded"
          value={filters.vehicle_id}
          onChange={(e) =>
            onFilter({
              ...filters,
              vehicle_id: e.target.value,
            })
          }
        />

        <input
          placeholder="Geofence ID"
          className="border p-2 rounded"
          value={filters.geofence_id}
          onChange={(e) =>
            onFilter({
              ...filters,
              geofence_id: e.target.value,
            })
          }
        />

        <input
          type="date"
          className="border p-2 rounded"
          value={filters.start_date}
          onChange={(e) =>
            onFilter({
              ...filters,
              start_date: e.target.value,
            })
          }
        />

        <input
          type="date"
          className="border p-2 rounded"
          value={filters.end_date}
          onChange={(e) =>
            onFilter({
              ...filters,
              end_date: e.target.value,
            })
          }
        />
      </div>

      <button
        className="mt-4 bg-blue-600 text-white px-5 py-2 rounded hover:bg-blue-700 transition"
        onClick={handleApplyFilters}
      >
        Apply Filters
      </button>
    </div>
  );
}
