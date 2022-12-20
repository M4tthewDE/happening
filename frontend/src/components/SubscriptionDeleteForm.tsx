import { Box, Button, Group, Text, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import axios from 'axios';

function handleSubmit(event: any) {
    axios.delete('https://happening.fdm.com.de/api/subscription?id=' + event.id)
        .then(res => {
            console.log(res);
        })
}

function SubscriptionDeleteForm() {
    const form = useForm(
        {
            initialValues: {
                id: '',
            },
        }
    );

    return (
        <Box sx={{ maxWidth: 300 }} mx="auto">
            <Text fw={700}>Delete Subscription</Text>
            <form onSubmit={form.onSubmit(handleSubmit)}>
                <TextInput
                    withAsterisk
                    label="ID"
                    placeholder="55e30e6f-d3f3-4222-91e8-63eb07bc9e9d"
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