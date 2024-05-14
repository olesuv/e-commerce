import { IOrderDetailsProps } from "../../../types/orderDetailsProps";

export default function OrderRating(props: IOrderDetailsProps) {
  return (
    <div className="my-2 text-sm text-neutral-500">
      {props.orderData?.rating ? (
        <span>
          Rating:{" "}
          <span className="font-semibold">{props.orderData?.rating}</span>
        </span>
      ) : (
        <span>Here would be rating (not available for now)</span>
      )}
    </div>
  );
}
