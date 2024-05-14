import { IOrderDetailsProps } from "../../../types/orderDetailsProps";
import convertDate from "../../../utils/convertDate";

export default function OrderDate(props: IOrderDetailsProps) {
  return (
    <div>
      <p className="text-xs text-neutral-500">
        Published on {convertDate(props.orderData?.date)}
      </p>
    </div>
  );
}
