import { Box, Table } from "@mantine/core";

interface SubscriptionListProps {
    rows: any;
}

function SubscriptionList({ rows }: SubscriptionListProps) {

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