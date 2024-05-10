import { useParams } from "react-router-dom";
import { gql, useQuery } from "@apollo/client";

const GET_ORDER = gql`
  query getOrder($id: ID!) {
    order(id: $id) {
      id
      title
      description
      category
      price
      currency
    }
  }
`;

export default function OrderDetails() {
  const { id } = useParams<{ id: string }>(); // Get the orderId from route params

  const { loading, error, data } = useQuery(GET_ORDER, {
    variables: { id },
  });

  if (loading) {
    return <p>Loading...</p>; // Render loading state while fetching data
  }

  if (error || !data?.order) {
    return <p>Order not found</p>; // Render error message if order is not found
  }

  const order = data.order;

  return (
    <div>
      <h1>{order.title}</h1>
      <p>{order.description}</p>
      <p>Category: {order.category}</p>
      <p>
        Price: {order.price} {order.currency}
      </p>
    </div>
  );
}