import { IOrderDateProps } from "../../../../types/orderDetailsLandingProps";
import convertDate from "../../../../utils/convertDate";

export default function OrderDate(props: IOrderDateProps) {
  return (
    <div>
      <p className="bottom-8 text-xs text-neutral-500 md:absolute">
        Published on {convertDate(props?.orderDate)}
      </p>
    </div>
  );
}
