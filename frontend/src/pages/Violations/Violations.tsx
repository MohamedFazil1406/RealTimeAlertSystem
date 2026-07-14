import { useEffect, useState } from "react";

import { getViolations } from "../../api/violation";

import { type Violation } from "../../types/violation";

import ViolationFilter from "../../components/Violation/ViolationFilter";

import ViolationTable from "../../components/Violation/ViolationTable";

export default function Violations() {
  const [violations, setViolations] = useState<Violation[]>([]);

  async function load(filters = {}) {
    const data = await getViolations(filters);
    console.log(data);

    setViolations(data.violations);
  }

  useEffect(() => {
    load();
  }, []);

  return (
    <div>
      <ViolationFilter onFilter={load} />

      <ViolationTable violations={violations} />
    </div>
  );
}
