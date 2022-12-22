import { useForm } from "@mantine/form";
import { Button, Container, Group, Text, TextInput } from '@mantine/core';

function UserToIdForm() {
    const form = useForm(
        {
            initialValues: {
                name: '',
            },
        }
    );

    function onSubmit(event: any) {
        console.log(event.name)
    }

    return (
        <Container size="xs">
            <Text fw={700}>User to ID</Text>
            <form onSubmit={form.onSubmit(onSubmit)}>
                <TextInput withAsterisk label="Username"
                    {...form.getInputProps('name')}
                />
                <Group position="right" mt="md">
                    <Button type="submit">Submit</Button>
                </Group>
            </form>
        </Container>
    );
}

export default UserToIdForm;