import React from "react";
import ReactDOM from "react-dom/client";

import { ApolloProvider } from "@apollo/client/react";

import App from "./App.jsx";
import client from "./apollo/client";

import "./index.css";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </React.StrictMode>
);
