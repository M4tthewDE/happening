import { IconHome, IconTimelineEvent, IconUser } from "@tabler/icons";
import NavbarLink from "./components/NavbarLink";

function Links() {

    return (
        <div>
            <NavbarLink route={"/"} text={"Home"} icon={<IconHome />} icon_color={"green"} />
            <NavbarLink route={"/user"} text={"User"} icon={<IconUser />} icon_color={"blue"} />
            <NavbarLink route={"/eventsub"} text={"Eventsub"} icon={<IconTimelineEvent />} icon_color={"blue"} />
        </div>
    );
}

export default Links;