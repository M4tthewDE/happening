import { Box, Table } from "@mantine/core";
import axios from "axios";
import { useEffect, useState } from "react";

interface SubscriptionIfc {
    id: string;
    target_id: string;
    subscription_type: string;
    status: string;
}

function SubscriptionList() {
    const [subscriptions, setSubscriptions] = useState<SubscriptionIfc[]>([])
    const [rows, setRows] = useState<any>([])

    useEffect(() => {
        axios.get('https://happening.fdm.com.de/api/subscription').then(res => {
            const subscriptions = res.data.subscriptions
            setSubscriptions(subscriptions)
        })
    }, []);

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
        <Box sx={{ maxWidth: 700 }} mx="auto">
            <Table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Target</th>
                        <th>Type</th>
                        <th>Status</th>
                    </tr>
                </thead>
                <tbody>{rows}</tbody>
            </Table>
        </Box>
    );
}

export default SubscriptionList;