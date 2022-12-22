import { Group, Text, ThemeIcon, UnstyledButton } from "@mantine/core";
import { IconHome, IconTimelineEvent } from "@tabler/icons";
import { Link } from "react-router-dom";

function Links() {

    return (
        <div>
            <Link to={"/"} style={{ textDecoration: 'none' }}>
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
                        <ThemeIcon color="green" variant="light">
                            <IconHome></IconHome>
                        </ThemeIcon>
                        <Text>Home</Text>
                    </Group>
                </UnstyledButton >
            </Link>
            <Link to={"/eventsub"} style={{ textDecoration: 'none' }}>
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
                        <ThemeIcon color="blue" variant="light">
                            <IconTimelineEvent></IconTimelineEvent>
                        </ThemeIcon>
                        <Text>Eventsub</Text>
                    </Group >
                </UnstyledButton >
            </Link>
        </div>
    );
}

export default Links;