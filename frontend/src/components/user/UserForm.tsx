import { useForm } from "@mantine/form";
import { Button, Container, Group, Space, Text, TextInput } from '@mantine/core';

interface UserFormProps {
    parentSubmit: (event: any) => void;
}

function UserForm({ parentSubmit }: UserFormProps) {
    const form = useForm(
        {
            initialValues: {
                name: '',
            },
        }
    );

    function onSubmit(event: any) {
        form.reset()
        parentSubmit(event);
    }

    return (
        <Container size="xs">
            <Text fw={700}>User Information</Text>
            <Space h="md" />
            <form onSubmit={form.onSubmit(onSubmit)}>
                <TextInput withAsterisk label="Name"
                    {...form.getInputProps('name')}
                />
                <Group position="right" mt="md">
                    <Button type="submit">Submit</Button>
                </Group>
            </form>
        </Container>
    );
}

export default UserForm;