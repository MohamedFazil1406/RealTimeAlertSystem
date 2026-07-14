import ReactDOM from "react-dom/client";
import "leaflet/dist/leaflet.css";
import { BrowserRouter } from "react-router-dom";

import App from "./App";

import "./index.css";
import { AlertProvider } from "./context/AlertContext";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <AlertProvider>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </AlertProvider>,
);
