import { useNavigate } from "react-router-dom";
import { decodeEmailFromToken } from "../../../utils/JWTHelpers";
import { gql, useMutation } from "@apollo/client";
import { IOrderDetailsProps } from "../../../types/orderDetailsProps";
import { verifyEmailFromCookie } from "../../../utils/cookieHelpers";

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

  const [buyOrder, { loading: buyOrderLoading, error: buyOrderError }] =
    useMutation(BUY_ORDER);

  function handleBuyOrder() {
    buyOrder({
      variables: {
        orderId: props?.orderData?.id,
        customerEmail: decodeEmailFromToken(),
      },
    }).then(() => navigate(0));
  }

  return verifyEmailFromCookie() ? (
    <div className="flex flex-wrap justify-center pt-2 md:justify-start">
      {buyOrderLoading ? (
        <button
          type="submit"
          className="w-2/3 cursor-not-allowed rounded-full bg-indigo-300 p-2 text-center text-white shadow-2xl"
        >
          Loading...
        </button>
      ) : buyOrderError?.message === "order already buyed" ? (
        <button
          type="submit"
          className="w-2/3 cursor-not-allowed rounded-full bg-indigo-300 p-2 text-center text-white shadow-2xl"
        >
          Already bought by someone
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
  ) : (
    <></>
  );
}
