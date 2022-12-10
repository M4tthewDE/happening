
import { Box, Button, Group, Select, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';

function SubscriptionForm() {
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
            <form onSubmit={form.onSubmit((values) => console.log(values))}>
                <TextInput
                    withAsterisk
                    label="Target ID"
                    placeholder="1234"
                    {...form.getInputProps('target_id')}
                />

                <Select
                    withAsterisk
                    label="Type"
                    placeholder="Pick one"
                    data={[
                        { value: 'FOLLOW', label: 'Follow' },
                        { value: 'Sub', label: 'Subscription' },
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

export default SubscriptionForm;