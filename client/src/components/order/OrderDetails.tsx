import { useParams } from "react-router-dom";
import { gql, useLazyQuery, useQuery } from "@apollo/client";
import { mapCurrencySymbol } from "../../../utils/mapCurrency";
import { useEffect, useState } from "react";

const GET_ORDER = gql`
  query getOrder($id: String!) {
    order(id: $id) {
      id
      title
      description
      category
      price
      currency
      authorEmail
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
  const [showMore, setShowMore] = useState(false);

  const toggleShowMore = () => {
    setShowMore(!showMore);
  };

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
    <div className="m-3 grid grid-rows-2 gap-4">
      <div className="rounded-xl border border-neutral-300 bg-neutral-100 text-center text-neutral-500">
        <p className="my-64">Images are not available (for now)</p>
      </div>

      <div>
        <div className="my-2 text-2xl font-semibold">{order?.title}</div>
        <ul className="flex flex-wrap *:my-1 *:mr-1 *:rounded-full *:border *:px-2 *:py-0.5 *:text-sm">
          {order?.category.map((category: string) => {
            switch (category) {
              case "Electronics":
                return (
                  <li
                    key={category}
                    className="cursor-pointer border-sky-100 bg-sky-50 text-sky-500 hover:border-sky-200 hover:bg-sky-100"
                  >
                    Electronics
                  </li>
                );
              case "Fashion":
                return (
                  <li
                    key={category}
                    className="cursor-pointer border-pink-100 bg-pink-50 text-pink-500 hover:border-pink-200 hover:bg-pink-100"
                  >
                    Fashion
                  </li>
                );
              case "Home":
                return (
                  <li
                    key={category}
                    className="cursor-pointer border-orange-100 bg-orange-50 text-orange-500 hover:border-orange-200 hover:bg-orange-100"
                  >
                    Home
                  </li>
                );
              case "Sports":
                return (
                  <li
                    key={category}
                    className="cursor-pointer border-indigo-100 bg-indigo-50 text-indigo-500 hover:border-indigo-200 hover:bg-indigo-100"
                  >
                    Sports
                  </li>
                );
              case "Books":
                return (
                  <li
                    key={category}
                    className="cursor-pointer border-green-100 bg-green-50 text-green-500 hover:border-green-200 hover:bg-green-100"
                  >
                    Books
                  </li>
                );
              default:
                return (
                  <li
                    key={category}
                    className="cursor-pointer border-neutral-100 bg-neutral-50 text-neutral-500 hover:border-neutral-200 hover:bg-neutral-100"
                  >
                    Other
                  </li>
                );
            }
          })}
        </ul>
        <div className="my-2 text-sm text-neutral-500">
          {order?.rating ? (
            <span>
              Rating: <span className="font-semibold">{order?.rating}</span>
            </span>
          ) : (
            <span>Here would be rating (not available for now)</span>
          )}
        </div>
        <div className="my-2 text-2xl font-semibold">
          {mapCurrencySymbol(order?.currency)} {order?.price}
        </div>
        <div className="my-2">
          <p className="text-md font-medium">Description</p>
          <p className="text-sm text-neutral-500">
            {showMore || order?.description.length <= 200 ? (
              order?.description
            ) : (
              <>
                {order?.description.slice(0, 200)}{" "}
                <span
                  onClick={toggleShowMore}
                  className="cursor-pointer text-indigo-500"
                >
                  Read more...
                </span>
              </>
            )}
            {showMore}
          </p>
          <div className="my-2">
            <p className="text-ms font-medium">Author</p>
            <p className="text-sm text-neutral-500">{author?.name}</p>
          </div>
        </div>
      </div>
    </div>
  );
}
