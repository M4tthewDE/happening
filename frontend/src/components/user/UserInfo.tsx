import { Container, Group, Paper, Text } from "@mantine/core";
import { UserIfc } from "../../User";

interface UserInfoProps {
    user: UserIfc;
}

function UserInfo({ user }: UserInfoProps) {
    return (
        <Container size="xs">
            <Paper shadow="xl" p="xl">
                <Group>
                    <Text fw={700}>Name:</Text>
                    <Text>{user.login}</Text>
                </Group>
                <Group>
                    <Text fw={700}>ID:</Text>
                    <Text>{user.id}</Text>
                </Group>
                <Group>
                    <Text fw={700}>Created at:</Text>
                    <Text>{user.created_at}</Text>
                </Group>
            </Paper>
        </Container>
    );
}

export default UserInfo;