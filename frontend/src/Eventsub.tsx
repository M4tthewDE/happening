import axios from "axios";
import SubscriptionAddForm from "./components/SubscriptionAddForm";
import SubscriptionDeleteForm from "./components/SubscriptionDeleteForm";
import SubscriptionList from "./components/SubscriptionList";
import { useEffect, useState } from "react";


interface SubscriptionIfc {
    id: string;
    target_id: string;
    subscription_type: string;
    status: string;
}

function Eventsub() {
    const [subscriptions, setSubscriptions] = useState<SubscriptionIfc[]>([])
    const [rows, setRows] = useState<any>([])

    function reloadSubs() {
        axios.get('https://happening.fdm.com.de/api/subscription').then(res => {
            const subscriptions = res.data.subscriptions
            setSubscriptions(subscriptions)
        })
    }

    function onAddSub(event: any) {
        console.log(event);
        axios.post('https://happening.fdm.com.de/api/subscription', event)
            .then(_ => {
                reloadSubs()
            })
    }

    function onDeleteSub(event: any) {
        axios.delete('https://happening.fdm.com.de/api/subscription?id=' + event.id)
            .then(_ => {
                reloadSubs()
            })
    }

    useEffect(() => {
        const rows = subscriptions.map((sub) => (
            <tr key={sub.id}>
                <td>{sub.id}</td>
                <td>{sub.target_id}</td>
                <td>{sub.subscription_type}</td>
                <td>{sub.status}</td>
            </tr>
        ))

        setRows(rows)
    }, [subscriptions]);

    return (
        <div>
            <SubscriptionDeleteForm parentSubmit={onDeleteSub}></SubscriptionDeleteForm>
            <SubscriptionAddForm parentSubmit={onAddSub}></SubscriptionAddForm>
            <SubscriptionList rows={rows}></SubscriptionList>
        </div>
    );
}

export default Eventsub;