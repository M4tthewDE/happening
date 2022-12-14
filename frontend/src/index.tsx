import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

import { MantineProvider } from "@mantine/core";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ErrorPage from "./Error";
import Eventsub from "./Eventsub";
import User from "./User";
import { NotificationsProvider } from "@mantine/notifications";
import Auth from "./Auth";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

const router = createBrowserRouter([
  {
    path: "/",
    element: <App>HOME</App>,
    errorElement: <ErrorPage />,
  },
  {
    path: "/disallowed",
    element: <>NOT ALLOWED (YET)</>,
    errorElement: <ErrorPage />,
  },
  {
    path: "/auth",
    element: <Auth />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/user",
    element: (
      <App>
        <User />
      </App>
    ),
    errorElement: <ErrorPage />,
  },
  {
    path: "/eventsub",
    element: (
      <App>
        <Eventsub />
      </App>
    ),
    errorElement: <ErrorPage />,
  },
]);

root.render(
  <React.StrictMode>
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <NotificationsProvider>
        <RouterProvider router={router} />
      </NotificationsProvider>
    </MantineProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
