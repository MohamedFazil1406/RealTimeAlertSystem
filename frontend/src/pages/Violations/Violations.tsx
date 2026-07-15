import { useEffect, useState } from "react";

import { getViolations } from "../../api/violation";
import { type Violation } from "../../types/violation";

import ViolationFilter from "../../components/Violation/ViolationFilter";
import ViolationTable from "../../components/Violation/ViolationTable";

export type ViolationFilters = {
  vehicle_id: string;
  geofence_id: string;
  start_date: string;
  end_date: string;
};

export default function Violations() {
  const [violations, setViolations] = useState<Violation[]>([]);

  const [filters, setFilters] = useState<ViolationFilters>(() => {
    const saved = localStorage.getItem("violationFilters");

    return saved
      ? JSON.parse(saved)
      : {
          vehicle_id: "",
          geofence_id: "",
          start_date: "",
          end_date: "",
        };
  });

  async function load(currentFilters: ViolationFilters = filters) {
    const data = await getViolations(currentFilters);
    console.log("API Response:", data);
    console.log("Violations:", data.violations);
    setViolations(data.violations);
  }

  useEffect(() => {
    load(filters);
  }, []);

  async function handleFilter(newFilters: ViolationFilters) {
    setFilters(newFilters);

    localStorage.setItem("violationFilters", JSON.stringify(newFilters));

    await load(newFilters);
  }

  return (
    <div>
      <ViolationFilter filters={filters} onFilter={handleFilter} />

      <ViolationTable violations={violations} />
    </div>
  );
}
