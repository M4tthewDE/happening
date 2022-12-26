import { AppShell, Header, Navbar } from "@mantine/core";
import Brand from "./Brand";
import Links from "./Links";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

interface AppProps {
  children: any;
}

function App({ children }: AppProps) {
  const [authenticated, setAuthenticated] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    if (localStorage.getItem("user_token") === null) {
      window.location.href = `https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=${process.env.REACT_APP_TWITCH_CLIENT_ID}&redirect_uri=${process.env.REACT_APP_PATH}/auth&scope=&force_verify=true`;
    } else {
      const token = localStorage.getItem("user_token");

      axios
        .get(`${process.env.REACT_APP_API_DOMAIN}/permissions?token=${token}`)
        .then(() => {
          setAuthenticated(true);
        })
        .catch((error) => {
          if (error.response) {
            navigate("/disallowed");
          }
        });
    }
  }, [navigate]);

  return (
    <>
      {authenticated && (
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
      )}
    </>
  );
}

export default App;
