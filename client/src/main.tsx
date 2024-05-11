import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import OrderDetails from "./components/order/OrderDetails.tsx";

import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import { loadErrorMessages, loadDevMessages } from "@apollo/client/dev";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import OrderDetailsError from "./components/order/OrderDetailsError.tsx";

const client = new ApolloClient({
  uri: "http://localhost:8080/query",
  cache: new InMemoryCache(),
});

if (process.env.NODE_ENV === "development") {
  loadDevMessages();
  loadErrorMessages();
}

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/order/:id",
    element: <OrderDetails />,
    errorElement: <OrderDetailsError />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <RouterProvider router={router} />
    </ApolloProvider>
  </React.StrictMode>,
);
