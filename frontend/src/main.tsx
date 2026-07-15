import ReactDOM from "react-dom/client";
import "leaflet/dist/leaflet.css";
import { BrowserRouter } from "react-router-dom";
import { VehicleProvider } from "./context/VehicleContext";

import App from "./App";

import "./index.css";
import { AlertProvider } from "./context/AlertContext";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <VehicleProvider>
    <AlertProvider>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </AlertProvider>
  </VehicleProvider>,
);
