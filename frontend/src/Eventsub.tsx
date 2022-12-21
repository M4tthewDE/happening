import axios from "axios";
import SubscriptionAddForm from "./components/SubscriptionAddForm";
import SubscriptionDeleteForm from "./components/SubscriptionDeleteForm";
import SubscriptionList from "./components/SubscriptionList";
import { useEffect, useState } from "react";
import { Button, Container, Space, Stack } from "@mantine/core";
import { IconRefresh } from "@tabler/icons";


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
            <Space h="xl" />
            <Container size="xs">
                <SubscriptionDeleteForm parentSubmit={onDeleteSub}></SubscriptionDeleteForm>
            </Container>
            <Space h="xl" />
            <Container size="xs">
                <SubscriptionAddForm parentSubmit={onAddSub}></SubscriptionAddForm>
            </Container>
            <Space h="xl" />
            <Container>
                <Stack align="center">
                    <SubscriptionList rows={rows}></SubscriptionList>
                    <Button onClick={reloadSubs}>
                        <IconRefresh></IconRefresh>
                    </Button>
                </Stack>
            </Container>
        </div >
    );
}

export default Eventsub;