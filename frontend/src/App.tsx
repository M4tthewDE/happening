import { AppShell, Header, Navbar } from "@mantine/core";
import Brand from "./Brand";
import Links from "./Links";

interface AppProps {
  children: any;
}

function App({ children }: AppProps) {
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
