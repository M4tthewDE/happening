import { Button, Group, Select, TextInput, Text, Space } from "@mantine/core";
import { useForm } from "@mantine/form";

interface SubscriptionAddFormProps {
  parentSubmit: (event: any) => void;
}

function SubscriptionAddForm({ parentSubmit }: SubscriptionAddFormProps) {
  const form = useForm({
    initialValues: {
      target_id: "",
      subscription_type: "channel.follow",
    },
  });

  function onSubmit(event: any) {
    form.reset();
    parentSubmit(event);
  }

  return (
    <div>
      <Text fw={700}>Add Subscription</Text>
      <Space h="md" />
      <form onSubmit={form.onSubmit(onSubmit)}>
        <TextInput
          withAsterisk
          label="Target ID"
          {...form.getInputProps("target_id")}
        />

        <Select
          withAsterisk
          label="Type"
          placeholder="Pick one"
          data={[{ value: "channel.follow", label: "Follow" }]}
          {...form.getInputProps("subscription_type")}
        />

        <Group position="right" mt="md">
          <Button type="submit">Submit</Button>
        </Group>
      </form>
    </div>
  );
}

export default SubscriptionAddForm;
