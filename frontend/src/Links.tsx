import { Group, Text, UnstyledButton } from "@mantine/core";
import { IconNotification } from "@tabler/icons";

function Links() {
    return (
        <UnstyledButton
            sx={(theme) => ({
                display: 'block',
                width: '100%',
                padding: theme.spacing.xs,
                borderRadius: theme.radius.sm,
                color: theme.colorScheme === 'dark' ? theme.colors.dark[0] : theme.black,

                '&:hover': {
                    backgroundColor:
                        theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
                },
            })}
        >
            <Group>
                <IconNotification></IconNotification>
                <Text>Eventsub</Text>
            </Group>
        </UnstyledButton>
    );
}

export default Links;