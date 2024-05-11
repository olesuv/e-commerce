import { useParams } from "react-router-dom";
import { gql, useQuery } from "@apollo/client";

const GET_ORDER = gql`
  query getOrder($id: String!) {
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
  const { id } = useParams<{ id: string }>();

  const { loading, data } = useQuery(GET_ORDER, {
    variables: { id },
  });

  if (loading) {
    return <p>Loading...</p>;
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
