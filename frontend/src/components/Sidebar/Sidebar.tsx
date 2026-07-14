import { Link, useLocation } from "react-router-dom";

const menu = [
  {
    name: "Dashboard",
    path: "/",
  },
  {
    name: "Vehicles",
    path: "/vehicles",
  },
  {
    name: "Geofences",
    path: "/geofences",
  },
  {
    name: "Alerts",
    path: "/alerts",
  },
  {
    name: "Violations",
    path: "/violations",
  },
];

export default function Sidebar() {
  const location = useLocation();

  return (
    <aside className="w-64 h-screen bg-slate-900 text-white">
      <div className="text-2xl font-bold p-6">GeoTracker</div>

      <nav>
        {menu.map((item) => (
          <Link
            key={item.path}
            to={item.path}
            className={`block px-6 py-3 hover:bg-slate-700 ${
              location.pathname === item.path ? "bg-slate-700" : ""
            }`}
          >
            {item.name}
          </Link>
        ))}
      </nav>
    </aside>
  );
}
