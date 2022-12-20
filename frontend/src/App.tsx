import SubscriptionAddForm from './components/SubscriptionAddForm';
import SubscriptionDeleteForm from './components/SubscriptionDeleteForm';
import SubscriptionList from './components/SubscriptionList';

export default function App() {

  return (
    <div>
      <br />
      <SubscriptionDeleteForm></SubscriptionDeleteForm>
      <br />
      <SubscriptionAddForm></SubscriptionAddForm>
      <br />
      <SubscriptionList></SubscriptionList>
    </div>
  );
}
