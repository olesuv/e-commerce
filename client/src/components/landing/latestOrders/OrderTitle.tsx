import { IOrderTitleProps } from "../../../../types/orderDetailsLandingProps";

export default function OrderTitle(props: IOrderTitleProps) {
  return (
    <div className="my-2 flex flex-wrap items-center justify-between text-lg font-semibold">
      <div className="mr-2 inline">{props.orderTitle}</div>
      <span className="text-xs font-normal text-neutral-500">
        <span className="text-green-500">✓ {props.orderStatus}</span>
      </span>
    </div>
  );
}
