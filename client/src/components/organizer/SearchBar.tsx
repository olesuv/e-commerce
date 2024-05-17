import { useState } from "react";
import CreateOrderPopup from "./CreateOrderPopup";
import { gql, useLazyQuery, useQuery } from "@apollo/client";

const ORDER_SEARCH = gql`
  query orderSearch($search: String!) {
    searchOrder(userInput: $search) {
      id
      title
      price
      currency
    }
  }
`;

export default function SearchBar() {
  const [search, setSearch] = useState("");
  const [showPopup, setShowPopup] = useState(false);

  const [searchOrders, { data, loading, error }] = useLazyQuery(ORDER_SEARCH, {
    variables: { search },
  });

  return (
    <div className="flex flex-col md:flex-row">
      <input
        type="text"
        placeholder="Search for products..."
        onChange={(e) => {
          setSearch(e.target.value);

          console.log("searching");
          searchOrders();
        }}
        className="rounded-xl border border-gray-200 bg-gray-100 p-2 outline-indigo-300 md:mr-2 md:basis-2/3"
      />
      {loading && <p>Loading...</p>}
      {error && <p>Error: {error.message}</p>}
      {data && (
        <div className="">
          {data.searchOrder.map((order: any) => (
            <div
              key={order.id}
              className="flex items-center justify-between rounded-xl bg-gray-100 p-2"
            >
              <p>{order.title}</p>
              <p>
                {order.price} {order.currency}
              </p>
            </div>
          ))}
        </div>
      )}

      {/*FIX: medium screen responsive*/}
      <div className="mt-2 grid grid-cols-2 gap-4 md:mt-0 md:basis-1/3">
        <button className="rounded-xl bg-indigo-500 p-2 text-white hover:bg-indigo-700">
          My Orders
        </button>
        <button
          onClick={() => setShowPopup(true)}
          className="rounded-xl bg-indigo-500 p-2 text-white hover:bg-indigo-700"
        >
          Sell products
        </button>
        {showPopup && <CreateOrderPopup setShowPopup={setShowPopup} />}
      </div>
    </div>
  );
}
