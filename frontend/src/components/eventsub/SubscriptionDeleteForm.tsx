import { Button, Group, Space, Text, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";

interface SubscriptionDeleteFormProps {
  parentSubmit: (event: any) => void;
}

function SubscriptionDeleteForm({ parentSubmit }: SubscriptionDeleteFormProps) {
  const form = useForm({
    initialValues: {
      id: "",
    },
  });
  function onSubmit(event: any) {
    form.reset();
    parentSubmit(event);
  }

  return (
    <div>
      <Text fw={700}>Delete Subscription</Text>
      <Space h="md" />
      <form onSubmit={form.onSubmit(onSubmit)}>
        <TextInput withAsterisk label="ID" {...form.getInputProps("id")} />
        <Group position="right" mt="md">
          <Button type="submit">Submit</Button>
        </Group>
      </form>
    </div>
  );
}

export default SubscriptionDeleteForm;
