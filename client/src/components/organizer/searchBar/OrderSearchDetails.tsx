import convertDate from "../../../../utils/convertDate";
import IOrderInfoProps from "../../../../types/orderDetailsSearch";

import { mapCurrencySymbol } from "../../../../utils/mapCurrency";
import { useNavigate } from "react-router-dom";

export default function OrderSearchDetails(props: IOrderInfoProps) {
  const navigate = useNavigate();

  return (
    <div
      className="flex items-center justify-between p-3 font-semibold"
      onClick={() => navigate(`order/${props.order?.id}`)}
    >
      <div className="grid grid-cols-1">
        <div>{props.order?.title}</div>
        <div className="text-xs font-normal text-gray-500">
          Published on {convertDate(props.order?.date)}
        </div>
      </div>
      <div>
        {props.order?.price} {mapCurrencySymbol(props.order?.currency)}
      </div>
    </div>
  );
}
