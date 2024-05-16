import OrderTitle from "./latestOrders/OrderTitle";
import OrderCategory from "./latestOrders/OrderCategory";
import OrderPrice from "../order/OrderPrice";
import OrderDate from "./latestOrders/OrderDate";

import { gql, useQuery } from "@apollo/client";
import { useNavigate } from "react-router-dom";

const GET_LAST_ORDERS = gql`
  query GetLastPosts {
    latestOrders {
      id
      title
      status
      price
      currency
      description
      category
      date
    }
  }
`;

export default function LatestOrders() {
  const navigate = useNavigate();
  const { loading, error, data } = useQuery(GET_LAST_ORDERS);

  if (loading) return <p>Loading...</p>;

  if (error) return console.error(error);

  return (
    <div className="m-3 grid grid-rows-1 gap-4 md:grid-cols-5">
      <div className="md:col-span-3 md:col-start-2">
        <div className="md:relative md:grid md:grid-cols-5 md:gap-4">
          {data.latestOrders.map((order: any) => (
            <div
              key={order.id}
              className="my-4 cursor-pointer rounded-xl border border-indigo-200 bg-white p-4 shadow-md hover:border-indigo-500 hover:duration-300"
              onClick={() => navigate(`order/${order?.id}`)}
            >
              <OrderTitle
                orderTitle={order?.title}
                orderStatus={order?.status}
              />
              <OrderCategory orderCategory={order?.category} />
              <OrderPrice orderData={order} />
              <OrderDate orderDate={order?.date} />
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
