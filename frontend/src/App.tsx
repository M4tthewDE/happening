import { AppShell, Header, Navbar } from "@mantine/core";
import Brand from "./Brand";
import Links from "./Links";
import Eventsub from "./Eventsub";

function App() {

  return (
    <AppShell
      padding="md"
      navbar={<Navbar width={{ base: 300 }} height={500} p="xs">
        <Links /></Navbar>
      }
      header={<Header height={60} p="xs">
        <Brand />
      </Header>}
    >
      <Eventsub />
    </AppShell>
  );
}

export default App;