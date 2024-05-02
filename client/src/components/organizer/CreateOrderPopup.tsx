import { useState } from "react";
import OrderName from "./create-order-form/OrderName";
import OrderDescription from "./create-order-form/OrderDescription";
import OrderCategories from "./create-order-form/OrderCategories";
import OrderPrice from "./create-order-form/OrderPrice";
import OrderHeader from "./create-order-form/OrderHeader";

interface ICreateOrderPopupProps {
  setShowPopup: (value: boolean) => void;
}

export default function CreateOrderPopup(props: ICreateOrderPopupProps) {
  const [orderName, setOrderName] = useState<string>("");
  const [orderDescription, setOrderDescription] = useState<string>("");
  const [orderCategories, setOrderCategories] = useState<number[]>([]);
  const [orderPrice, setOrderPrice] = useState<number>(0);

  const [error, setError] = useState<string>("");

  return (
    <div className="fixed top-0 left-0 w-full h-full bg-gray-800 bg-opacity-60 flex justify-center items-center">
      <div className="bg-white p-8 rounded-lg">
        <OrderHeader />

        {error && (
          <label className="flex flex-col bg-rose-500 w-full rounded-md my-2">
            <p className="text-xl text-white text-center font-semibold p-5">
              Oops, {error}
            </p>
          </label>
        )}

        <form className="flex flex-col gap-2 mt-4 mb-4">
          <OrderName setOrderName={setOrderName} />
          <OrderDescription setOrderDescription={setOrderDescription} />
          <OrderCategories setOrderCategories={setOrderCategories} />
          <OrderPrice setOrderPrice={setOrderPrice} />

          <div className="grid grid-cols-2 gap-4">
            <button
              onClick={() => props.setShowPopup(false)}
              className="mt-2 p-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-md"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="mt-2 p-2 bg-indigo-500 hover:bg-indigo-600 text-white rounded-md"
            >
              Sell
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
