import { Divider, Group, Space, Text } from "@mantine/core";

function Brand() {
    return (
        <div>
            <Group position="center">
                <Text fw={700}>HAPPENING</Text>
            </Group>
            <Space h="md" />
            <Divider />
        </div>
    );
}

export default Brand;