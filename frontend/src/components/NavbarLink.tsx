import { UnstyledButton, Group, ThemeIcon, Text } from "@mantine/core";
import { Link } from "react-router-dom";

interface NavbarLinkProps {
  route: string;
  text: string;
  icon: any;
  icon_color: string;
}

function NavbarLink(props: NavbarLinkProps) {
  return (
    <Link to={props.route} style={{ textDecoration: "none" }}>
      <UnstyledButton
        sx={(theme) => ({
          display: "block",
          width: "100%",
          padding: theme.spacing.xs,
          borderRadius: theme.radius.sm,
          color:
            theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.black,

          "&:hover": {
            backgroundColor:
              theme.colorScheme === "dark"
                ? theme.colors.dark[6]
                : theme.colors.gray[0],
          },
        })}
      >
        <Group>
          <ThemeIcon color={props.icon_color} variant="light">
            {props.icon}
          </ThemeIcon>
          <Text>{props.text}</Text>
        </Group>
      </UnstyledButton>
    </Link>
  );
}

export default NavbarLink;
