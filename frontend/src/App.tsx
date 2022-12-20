import SubscriptionDeleteForm from './components/SubscriptionDeleteForm';
import SubscriptionForm from './components/SubscriptionForm';
import SubscriptionList from './components/SubscriptionList';

export default function App() {

  return (
    <div>
      <br />
      <SubscriptionDeleteForm></SubscriptionDeleteForm>
      <br />
      <SubscriptionForm></SubscriptionForm>
      <br />
      <SubscriptionList></SubscriptionList>
    </div>
  );
}
