import { useState } from "react";
import { LoaderFunctionArgs, useLoaderData } from "react-router-dom";
import { gql, useMutation } from "@apollo/client";

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

const [error, setError] = useState<string>("");

const [getOrderMutation, { loading }] = useMutation(GET_ORDER, {
  onError: (error) => {
    setError(error.message);
  },
  onCompleted: (data) => {
    return data;
  },
});

export async function loader(args: LoaderFunctionArgs<{ orderId: string }>) {
  return await getOrderMutation({
    variables: {
      id: args.params.orderId,
    },
  });
}

export default function OrderDetails() {
  const orderData = useLoaderData();

  return <h1>{orderData as React.ReactNode}</h1>;
}
