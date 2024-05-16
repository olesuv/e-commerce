import { useState } from "react";
import OrderName from "./create-order-form/OrderName";
import OrderDescription from "./create-order-form/OrderDescription";
import OrderCategories from "./create-order-form/OrderCategories";
import OrderPrice from "./create-order-form/OrderPrice";
import OrderHeader from "./create-order-form/OrderHeader";
import { OrderCategory, OrderCurrency } from "../../../types/orderTypes";
import { gql, useMutation } from "@apollo/client";
import { useNavigate } from "react-router";

const CREATE_ORDER = gql`
  mutation createOrder($input: CreateOrderInput!) {
    createOrder(input: $input) {
      id
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
    OrderCurrency.UAH,
  );

  const [error, setError] = useState<string>("");
  const navigate = useNavigate();

  const [createOrderMutation, { loading }] = useMutation(CREATE_ORDER, {
    onError: (error) => {
      setError(error.message);
    },
    onCompleted: (data) => {
      const orderId = data?.createOrder?.id;
      props.setShowPopup(false);
      navigate(`order/${orderId}`);
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
    <div className="fixed left-0 top-0 flex h-full w-full items-center justify-center bg-gray-800 bg-opacity-60">
      <div className="rounded-lg bg-white p-8">
        <OrderHeader />

        {error && (
          <label className="my-2 flex w-full flex-col rounded-md bg-rose-500">
            <p className="p-5 text-center text-xl font-semibold text-white">
              Oops, {error}
            </p>
          </label>
        )}

        <form className="mb-4 mt-4 flex flex-col gap-2" onSubmit={handleSubmit}>
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
              className="mt-2 rounded-md bg-gray-200 p-2 text-gray-700 hover:bg-gray-300"
            >
              Cancel
            </button>
            {loading ? (
              <button
                type="submit"
                className="mt-2 cursor-wait rounded-md bg-indigo-300 p-2 hover:bg-indigo-400"
              >
                Selling...
              </button>
            ) : (
              <button
                type="submit"
                className="mt-2 cursor-pointer rounded-md bg-indigo-500 p-2 text-white hover:bg-indigo-600"
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
