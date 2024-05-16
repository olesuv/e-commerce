import { useParams } from "react-router-dom";
import { gql, useLazyQuery, useQuery } from "@apollo/client";
import { useEffect } from "react";

import OrderStatus from "./OrderStatus";
import OrderRating from "./OrderRating";
import OrderDescription from "./OrderDescription";
import OrderCategory from "./OrderCategory";
import OrderAuthor from "./OrderAuthor";
import OrderDate from "./OrderDate";
import OrderPrice from "./OrderPrice";
import OrderButton from "./OrderButton";

const GET_ORDER = gql`
  query getOrder($id: String!) {
    order(id: $id) {
      id
      title
      description
      category
      price
      currency
      customerEmail
      authorEmail
      date
      status
    }
  }
`;

const GET_USER = gql`
  query getUserByEmail($email: String!) {
    user(email: $email) {
      email
      name
    }
  }
`;

export default function OrderDetails() {
  const { id } = useParams<{ id: string }>();

  const { loading: orderLoading, data: orderData } = useQuery(GET_ORDER, {
    variables: { id },
  });

  const [getUserByEmail, { loading: userLoading, data: userData }] =
    useLazyQuery(GET_USER);

  useEffect(() => {
    if (orderData?.order?.authorEmail) {
      getUserByEmail({ variables: { email: orderData.order.authorEmail } });
    }
  }, [orderData, getUserByEmail]);

  if (orderLoading || userLoading) {
    return <p>Loading...</p>;
  }

  const order = orderData?.order;
  const author = userData?.user;

  return (
    <div className="m-3 grid grid-rows-2 gap-4 md:grid-cols-5">
      <div className="md:col-span-3 md:col-start-2">
        <div className="md:grid md:grid-cols-2 md:gap-4">
          <div className="rounded-xl border border-neutral-300 bg-neutral-100 text-center text-neutral-500">
            <p className="my-64">Images are not available (for now)</p>
          </div>

          <div>
            <OrderStatus orderData={order} />
            <OrderCategory orderData={order} />
            <OrderRating orderData={order} />
            <OrderPrice orderData={order} />
            <OrderDescription orderData={order} />
            <OrderAuthor authorData={author} />
            <OrderDate orderData={order} />
            <OrderButton orderData={order} />
          </div>
        </div>
      </div>
    </div>
  );
}
