import { useState } from "react";
import OrderName from "./create-order-form/OrderName";
import OrderDescription from "./create-order-form/OrderDescription";
import OrderCategories from "./create-order-form/OrderCategories";
import OrderPrice from "./create-order-form/OrderPrice";
import OrderHeader from "./create-order-form/OrderHeader";
import { OrderCategory, OrderCurrency } from "../../../types/orderTypes";
import { gql, useMutation } from "@apollo/client";

const CREATE_ORDER = gql`
  mutation createOrder($input: CreateOrderInput!) {
    createOrder(input: $input) {
      id
      title
      price
    }
  }
`;

interface ICreateOrderPopupProps {
  setShowPopup: (value: boolean) => void;
}

export default function CreateOrderPopup(props: ICreateOrderPopupProps) {
  const [orderName, setOrderName] = useState<string>("");
  const [orderDescription, setOrderDescription] = useState<string>("");
  const [orderCategories, setOrderCategories] = useState<OrderCategory[]>([]);
  const [orderPrice, setOrderPrice] = useState<number>();
  const [orderCurrency, setOrderCurrency] = useState<OrderCurrency>(
    OrderCurrency.UAH
  );

  const [error, setError] = useState<string>("");

  const [createOrderMutation, { loading }] = useMutation(CREATE_ORDER, {
    onError: (error) => {
      setError(error.message);
    },
    onCompleted: (data) => {
      console.log(data);
      props.setShowPopup(false);
    },
  });

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    createOrderMutation({
      variables: {
        input: {
          title: orderName,
          description: orderDescription,
          category: orderCategories,
          price: orderPrice,
          currency: orderCurrency,
        },
      },
    });
  };

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

        <form className="flex flex-col gap-2 mt-4 mb-4" onSubmit={handleSubmit}>
          <OrderName setOrderName={setOrderName} />
          <OrderDescription setOrderDescription={setOrderDescription} />
          <OrderCategories
            orderCategories={orderCategories}
            setOrderCategories={setOrderCategories}
          />
          <OrderPrice
            setOrderPrice={setOrderPrice}
            setOrderCurrency={setOrderCurrency}
          />

          <div className="grid grid-cols-2 gap-4">
            <button
              onClick={() => props.setShowPopup(false)}
              className="mt-2 p-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-md"
            >
              Cancel
            </button>
            {loading ? (
              <button
                type="submit"
                className="mt-2 p-2 bg-indigo-300 hover:bg-indigo-400 cursor-wait rounded-md"
              >
                Selling...
              </button>
            ) : (
              <button
                type="submit"
                className="mt-2 p-2 bg-indigo-500 hover:bg-indigo-600 text-white rounded-md cursor-pointer"
              >
                Sell
              </button>
            )}
          </div>
        </form>
      </div>
    </div>
  );
}
