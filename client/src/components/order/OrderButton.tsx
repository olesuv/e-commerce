import { useNavigate } from "react-router-dom";
import { decodeEmailFromToken } from "../../../utils/getEmailFromJWT";
import { gql, useMutation } from "@apollo/client";
import { IOrderDetailsProps } from "../../../types/orderDetailsProps";

const BUY_ORDER = gql`
  mutation buy($orderId: String!, $customerEmail: String!) {
    buyOrder(orderId: $orderId, customerEmail: $customerEmail) {
      status
      customerEmail
      authorEmail
    }
  }
`;

export default function OrderButton(props: IOrderDetailsProps) {
  const navigate = useNavigate();

  const [buyOrder, { loading: buyOrderLoading }] = useMutation(BUY_ORDER);

  function handleBuyOrder() {
    console.log("Buying order...");
    buyOrder({
      variables: {
        orderId: props?.orderData?.id,
        customerEmail: decodeEmailFromToken(),
      },
    }).then(() => navigate(0));
  }

  return (
    <div className="flex flex-wrap justify-center pt-2 md:justify-start">
      {buyOrderLoading ? (
        <button
          type="submit"
          className="w-2/3 cursor-not-allowed rounded-full bg-indigo-300 p-2 text-center text-white shadow-2xl"
        >
          Loading...
        </button>
      ) : (
        <button
          className="w-2/3 cursor-pointer rounded-full bg-indigo-500 p-2 text-center text-white shadow-2xl hover:bg-indigo-700"
          onClick={handleBuyOrder}
        >
          Buy
        </button>
      )}
    </div>
  );
}
