import { IOrderDetailsProps } from "../../../types/orderDetailsProps";

export default function OrderStatus(props: IOrderDetailsProps) {
  return (
    <div className="my-2 flex flex-wrap items-center text-2xl font-semibold">
      <div className="mr-2">{props.orderData?.title}</div>
      <span className="text-sm font-normal text-neutral-500">
        {props.orderData?.status === "Available" ? (
          <span className="text-green-500">✓ {props.orderData?.status}</span>
        ) : (
          <span className="text-neutral-500">
            ✗ {props.orderData?.status} (by{" "}
            <span className="text-indigo-500">
              {props.orderData?.customerEmail}
            </span>
            )
          </span>
        )}
      </span>
    </div>
  );
}
