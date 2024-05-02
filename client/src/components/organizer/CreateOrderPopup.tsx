import { useState } from "react";

interface ICreateOrderPopupProps {
  setShowPopup: (value: boolean) => void;
}

export default function CreateOrderPopup(props: ICreateOrderPopupProps) {
  const [error, setError] = useState<string>("");

  return (
    <div className="fixed top-0 left-0 w-full h-full bg-gray-800 bg-opacity-60 flex justify-center items-center">
      <div className="bg-white p-8 rounded-lg">
        <p className="text-3xl font-semibold mt-2 text-center">
          Sell your favourite{" "}
          <span className="font-serif inline-block">goods</span>
        </p>

        {error && (
          <label className="flex flex-col bg-rose-500 w-full rounded-md my-2">
            <p className="text-xl text-white text-center font-semibold p-5">
              Oops, {error}
            </p>
          </label>
        )}

        <form className="flex flex-col gap-2 mt-4 mb-4">
          <div>
            <div className="after:content-['*'] after:ml-0.5 after:text-red-500">
              Order name
            </div>
            <input
              type="text"
              placeholder="Example: Samsung Galaxy S21 Ultra"
              className="p-2 border border-gray-200 outline-indigo-300 rounded-md w-full"
            />
          </div>

          <div>
            <p>Description</p>
            <textarea
              placeholder="Example: Brand new Samsung Galaxy S21 Ultra, 256GB, 12GB RAM, Phantom Black"
              className="p-2 border border-gray-200 outline-indigo-300 rounded-md h-36 w-full text-left"
            />
          </div>

          <div>
            <p className="inline-block">
              <span className="after:content-['*'] after:ml-0.5 after:text-red-500">
                Categories
              </span>
              <span className="pl-1 text-xs text-gray-400">
                (You can select <span className="underline">multiple</span>)
              </span>
            </p>

            <ul className="flex flex-wrap *:rounded-full *:border *:px-2 *:py-0.5 *:mr-1 *:mt-1">
              <li className="border-sky-100 bg-sky-50 text-sky-500 hover:border-sky-200 hover:bg-sky-100 cursor-pointer">
                Electronics
              </li>
              <li className="border-pink-100 bg-pink-50 text-pink-500 hover:border-pink-200 hover:bg-pink-100 cursor-pointer">
                Fashion
              </li>
              <li className="border-orange-100 bg-orange-50 text-orange-500 hover:border-orange-200 hover:bg-orange-100 cursor-pointer">
                Home
              </li>
              <li className="border-indigo-100 bg-indigo-50 text-indigo-500 hover:border-indigo-200 hover:bg-indigo-100 cursor-pointer">
                Sports
              </li>
              <li className="border-green-100 bg-green-50 text-green-500 hover:border-green-200 hover:bg-green-100 cursor-pointer">
                Books
              </li>
              <li className="border-emerald-100 bg-emerald-50 text-emerald-500 hover:border-emerald-200 hover:bg-emerald-100 cursor-pointer">
                Automotive
              </li>
              <li className="border-neutral-100 bg-neutral-50 text-neutral-500 hover:border-neutral-200 hover:bg-neutral-100 cursor-pointer">
                Other
              </li>
            </ul>
          </div>

          <div>
            <label
              htmlFor="price"
              className="block text-sm font-medium leading-6 text-gray-900 after:content-['*'] after:ml-0.5 after:text-red-500"
            >
              Price
            </label>
            <div className="relative mt-2 rounded-md shadow-sm">
              <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                <span className="text-gray-500 sm:text-sm">$</span>
              </div>
              <input
                type="text"
                name="price"
                id="price"
                className="block w-full rounded-md border-0 py-1.5 pl-7 pr-20 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm sm:leading-6 outline-indigo-300"
                placeholder="0.00"
              />
              <div className="absolute inset-y-0 right-0 flex items-center">
                <label htmlFor="currency" className="sr-only">
                  Currency
                </label>
                <select
                  id="currency"
                  name="currency"
                  className="h-full rounded-md border-0 bg-transparent py-0 pl-2 pr-7 text-gray-500 sm:text-sm sm:leading-6 outline-indigo-300"
                >
                  <option>UAH</option>
                  <option>USD</option>
                  <option>EUR</option>
                </select>
              </div>
            </div>
          </div>

          <div className="flex">
            <button
              type="submit"
              onClick={() => props.setShowPopup(false)}
              className="w-full mt-2 p-2 bg-indigo-500 hover:bg-indigo-700 text-white rounded-md"
            >
              Sell
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
