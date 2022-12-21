import { Box, Button, Group, Text, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';


interface SubscriptionDeleteFormProps {
    parentSubmit: (event: any) => void;
}

function SubscriptionDeleteForm({ parentSubmit }: SubscriptionDeleteFormProps) {
    const form = useForm(
        {
            initialValues: {
                id: '',
            },
        }
    );
    function onSubmit(event: any) {
        form.reset()
        parentSubmit(event);
    }

    return (
        <Box sx={{ maxWidth: 300 }} mx="auto">
            <Text fw={700}>Delete Subscription</Text>
            <form onSubmit={form.onSubmit(onSubmit)}>
                <TextInput
                    withAsterisk
                    label="ID"
                    {...form.getInputProps('id')}
                />
                <Group position="right" mt="md">
                    <Button type="submit">Submit</Button>
                </Group>
            </form>
        </Box>
    );
}

export default SubscriptionDeleteForm;