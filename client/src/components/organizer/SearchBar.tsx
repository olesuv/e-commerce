import { useState } from "react";

export default function SearchBar() {
  const [search, setSearch] = useState("");

  return (
    <div className="flex flex-col md:flex-row">
      <input
        type="text"
        placeholder="Search for products..."
        onChange={(e) => {
          setSearch(e.target.value);
          // TODO: implement search function
          // handleSearch(e.target.value);
        }}
        className="md:basis-2/3 md:mr-2 p-2 border border-gray-200 outline-indigo-300 rounded-xl bg-gray-100"
      />
      {/*FIX: medium screen responsive*/}
      <div className="md:basis-1/3 grid grid-cols-2 gap-4 mt-2 md:mt-0">
        <button className="p-2 bg-indigo-500 hover:bg-indigo-700 text-white rounded-xl">
          My Orders
        </button>
        <button className="p-2 bg-indigo-500 hover:bg-indigo-700 text-white rounded-xl">
          Sell products
        </button>
      </div>
    </div>
  );
}
