import { useForm } from "@mantine/form";
import {
  Button,
  Container,
  Group,
  Space,
  Switch,
  Text,
  TextInput,
} from "@mantine/core";
import { useState } from "react";

interface UserFormProps {
  parentSubmit: (event: any) => void;
}

function UserForm({ parentSubmit }: UserFormProps) {
  const [checked, setChecked] = useState(false);

  const form = useForm({
    initialValues: {
      input: "",
    },
  });

  function onSubmit(event: any) {
    form.reset();

    const type = checked ? "id" : "name";
    parentSubmit({ value: event.input, type: type });
  }

  function getLabel() {
    const label = checked ? "ID" : "Name";
    return label;
  }

  return (
    <Container size="xs">
      <Text fw={700}>User Information</Text>
      <Space h="md" />
      <form onSubmit={form.onSubmit(onSubmit)}>
        <TextInput
          withAsterisk
          label={getLabel()}
          {...form.getInputProps("input")}
        />
        <Group position="apart" mt="md">
          <Switch
            label="ID"
            checked={checked}
            onChange={(event) => setChecked(event.currentTarget.checked)}
          />
          <Button type="submit">Submit</Button>
        </Group>
      </form>
    </Container>
  );
}

export default UserForm;
