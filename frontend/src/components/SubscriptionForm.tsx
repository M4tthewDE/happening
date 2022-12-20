import { Box, Button, Group, Select, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import axios from 'axios';

function handleSubmit(event: any) {
    console.log(event);
    axios.post('https://happening.fdm.com.de/api/subscription', event)
        .then(res => {
            console.log(res);
        })
}

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
            <form onSubmit={form.onSubmit(handleSubmit)}>
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

export default SubscriptionForm;