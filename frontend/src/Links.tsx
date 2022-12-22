import { IconHome, IconTimelineEvent } from "@tabler/icons";
import NavbarLink from "./components/NavbarLink";

function Links() {

    return (
        <div>
            <NavbarLink route={"/"} text={"Home"} icon={<IconHome />} icon_color={"green"} />
            <NavbarLink route={"/eventsub"} text={"Eventsub"} icon={<IconTimelineEvent />} icon_color={"blue"} />
        </div>
    );
}

export default Links;