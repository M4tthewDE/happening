import axios from "axios";
import SubscriptionAddForm from "./components/SubscriptionAddForm";
import SubscriptionDeleteForm from "./components/SubscriptionDeleteForm";
import SubscriptionList from "./components/SubscriptionList";

export default function Eventsub() {

    function onAddSub(event: any) {
        console.log(event);
        axios.post('https://happening.fdm.com.de/api/subscription', event)
            .then(res => {
                console.log(res);
            })
    }

    function onDeleteSub(event: any) {
        axios.delete('https://happening.fdm.com.de/api/subscription?id=' + event.id)
            .then(res => {
                console.log(res);
            })
    }

    return (
        <div>
            <SubscriptionDeleteForm onSubmit={onDeleteSub}></SubscriptionDeleteForm>
            <SubscriptionAddForm onSubmit={onAddSub}></SubscriptionAddForm>
            <SubscriptionList></SubscriptionList>
        </div>
    );
}