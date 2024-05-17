import { useState } from "react";
import CreateOrderPopup from "./CreateOrderPopup";
import { gql, useLazyQuery } from "@apollo/client";

import OrderSearchDetails from "./searchBar/OrderSearchDetails";

const ORDER_SEARCH = gql`
  query orderSearch($search: String!) {
    searchOrder(userInput: $search) {
      id
      title
      price
      currency
      date
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
          searchOrders();
        }}
        className="rounded-xl border border-gray-200 bg-gray-100 p-2 outline-indigo-300 md:mr-2 md:basis-2/3"
      />
      {loading && <p>Loading...</p>}
      {error && <p>Error: {error.message}</p>}
      {data && (
        <div className="mt-2 divide-y divide-indigo-100 rounded-lg border border-indigo-200">
          {data.searchOrder.map((order: any) => (
            <div key={order.id}>
              <OrderSearchDetails order={order} />
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
