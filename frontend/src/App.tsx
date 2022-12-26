import { AppShell, Header, Navbar } from "@mantine/core";
import Brand from "./Brand";
import Links from "./Links";
import { v4 as uuidv4 } from "uuid";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import React from "react";

interface AppProps {
  children: any;
}

interface PermissionsIfc {
  permissions: string[];
}

function App({ children }: AppProps) {
  const navigate = useNavigate();

  React.useEffect(() => {
    if (localStorage.getItem("user_token") === null) {
      window.location.href = `https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=${process.env.REACT_APP_TWITCH_CLIENT_ID}&redirect_uri=${process.env.REACT_APP_PATH}/auth&scope=&force_verify=true`;
    } else {
      const token = localStorage.getItem("user_token");

      axios
        .get(`${process.env.REACT_APP_API_DOMAIN}/permissions?token=${token}`)
        .then((res) => {
          console.log(res.data);
          const permissions: PermissionsIfc = {
            permissions: res.data,
          };
          if (!permissions.permissions.includes("ALL")) {
            navigate("/disallowed");
          }
        })
        .catch((error) => {
          if (error.response) {
            if (error.response.status === 403) {
              navigate("/disallowed");
            }
          }
        });
    }
  });

  return (
    <AppShell
      padding="md"
      navbar={
        <Navbar width={{ base: 300 }} p="xs">
          <Links />
        </Navbar>
      }
      header={
        <Header height={60} p="xs">
          <Brand />
        </Header>
      }
    >
      {children}
    </AppShell>
  );
}

export default App;
