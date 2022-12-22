import { useForm } from "@mantine/form";
import { Button, Container, Group, Text, TextInput } from '@mantine/core';
import axios from "axios";

function UserForm() {
    const form = useForm(
        {
            initialValues: {
                name: '',
            },
        }
    );

    function onSubmit(event: any) {
        axios.get('https://happening.fdm.com.de/api/user?name=' + event.name).then(res => {
            console.log(res.data)
        })
    }

    // TODO:
    // - handle 404
    // - add option to input id
    // - display results

    return (
        <Container size="xs">
            <Text fw={700}>User Information</Text>
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