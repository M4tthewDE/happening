import { Navbar } from "@mantine/core";
import Brand from "./Brand";

function App() {

  return (
    <Navbar height={600} p="xs" width={{ base: 300 }}>
      <Navbar.Section>
        <Brand />
      </Navbar.Section>
      <Navbar.Section grow mt="md">{/* Links sections */}</Navbar.Section>
    </Navbar>
  );
}

export default App;