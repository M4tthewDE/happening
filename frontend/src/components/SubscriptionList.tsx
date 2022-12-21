import { Box, Table } from "@mantine/core";

interface SubscriptionListProps {
    rows: any;
}

function SubscriptionList({ rows }: SubscriptionListProps) {

    return (
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
    );
}

export default SubscriptionList;