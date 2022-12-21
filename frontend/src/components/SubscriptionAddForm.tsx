import { Box, Button, Group, Select, TextInput, Text } from '@mantine/core';
import { useForm } from '@mantine/form';


interface SubscriptionAddFormProps {
    onSubmit: (event: any) => void;
}

function SubscriptionAddForm({ onSubmit }: SubscriptionAddFormProps) {
    const form = useForm(
        {
            initialValues: {
                target_id: '',
                subscription_type: 'FOLLOW',
            },
        }
    );

    return (
        <Box sx={{ maxWidth: 300 }} mx="auto">
            <Text fw={700}>Add Subscription</Text>
            <form onSubmit={form.onSubmit(onSubmit)}>
                <TextInput
                    withAsterisk
                    label="Target ID"
                    {...form.getInputProps('target_id')}
                />

                <Select
                    withAsterisk
                    label="Type"
                    placeholder="Pick one"
                    data={[
                        { value: 'FOLLOW', label: 'Follow' },
                        { value: 'SUB', label: 'Subscription' },
                    ]}
                    {...form.getInputProps('subscription_type')}
                />

                <Group position="right" mt="md">
                    <Button type="submit">Submit</Button>
                </Group>
            </form>
        </Box>
    );
}

export default SubscriptionAddForm;