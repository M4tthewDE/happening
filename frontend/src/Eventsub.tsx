import SubscriptionAddForm from "./components/SubscriptionAddForm";
import SubscriptionDeleteForm from "./components/SubscriptionDeleteForm";
import SubscriptionList from "./components/SubscriptionList";

export default function Eventsub() {

    return (
        <div>
            <SubscriptionDeleteForm></SubscriptionDeleteForm>
            <SubscriptionAddForm></SubscriptionAddForm>
            <SubscriptionList></SubscriptionList>
        </div>
    );
}