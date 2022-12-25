import { AppShell, Header, Navbar } from "@mantine/core";
import Brand from "./Brand";
import Links from "./Links";
import { v4 as uuidv4 } from "uuid";

interface AppProps {
  children: any;
}

function App({ children }: AppProps) {
  if (localStorage.getItem("state") === null) {
    const state = uuidv4();
    localStorage.setItem("state", state);
  }

  if (localStorage.getItem("user_token") === null) {
    console.log(process.env);
    window.location.href = `https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=${process.env.REACT_APP_TWITCH_CLIENT_ID}&redirect_uri=${process.env.REACT_APP_PATH}/auth&scope=&state=${state}&force_verify=true`;
  }

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
