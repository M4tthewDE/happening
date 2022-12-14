import axios from "axios";
import { useEffect, useState } from "react";
import { Button, Container, Space, Stack } from "@mantine/core";
import { IconRefresh } from "@tabler/icons";
import SubscriptionAddForm from "./components/eventsub/SubscriptionAddForm";
import SubscriptionDeleteForm from "./components/eventsub/SubscriptionDeleteForm";
import SubscriptionList from "./components/eventsub/SubscriptionList";

interface SubscriptionIfc {
  id: string;
  target_id: string;
  subscription_type: string;
  status: string;
}

function Eventsub() {
  const [subscriptions, setSubscriptions] = useState<SubscriptionIfc[]>([]);
  const [rows, setRows] = useState<any>([]);

  function reloadSubs() {
    axios
      .get(process.env.REACT_APP_API_DOMAIN + "/subscription")
      .then((res) => {
        const subscriptions = res.data.subscriptions;
        setSubscriptions(subscriptions);
      });
  }

  function onAddSub(event: any) {
    console.log(event);
    axios
      .post(process.env.REACT_APP_API_DOMAIN + "/subscription", event)
      .then((_) => {
        reloadSubs();
      });
  }

  function onDeleteSub(event: any) {
    axios
      .delete(process.env.REACT_APP_API_DOMAIN + "/subscription?id=" + event.id)
      .then((_) => {
        reloadSubs();
      });
  }

  useEffect(() => {
    reloadSubs();
  }, []);

  useEffect(() => {
    const rows = subscriptions.map((sub) => (
      <tr key={sub.id}>
        <td>{sub.id}</td>
        <td>{sub.target_id}</td>
        <td>{sub.subscription_type}</td>
        <td>{sub.status}</td>
      </tr>
    ));

    setRows(rows);
  }, [subscriptions]);

  return (
    <div>
      <Space h="xl" />
      <Container size="xs">
        <SubscriptionDeleteForm
          parentSubmit={onDeleteSub}
        ></SubscriptionDeleteForm>
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
    </div>
  );
}

export default Eventsub;
