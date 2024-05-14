import { IOrderDetailsProps } from "../../../types/orderDetailsProps";
import { mapCurrencySymbol } from "../../../utils/mapCurrency";

export default function OrderPrice(props: IOrderDetailsProps) {
  return (
    <div className="my-2 text-2xl font-semibold">
      {mapCurrencySymbol(props?.orderData?.currency)} {props?.orderData?.price}
    </div>
  );
}
