import { useState } from "react";
import { IOrderDetailsProps } from "../../../types/orderDetailsProps";

export default function OrderDescription(props: IOrderDetailsProps) {
  const [showMore, setShowMore] = useState(false);

  const toggleShowMore = () => {
    setShowMore(!showMore);
  };

  return (
    <div className="mt-2">
      <p className="text-md font-medium">Description</p>
      <p className="text-sm text-neutral-500">
        {showMore || props.orderData?.description.length <= 200 ? (
          props.orderData?.description
        ) : (
          <>
            {props.orderData?.description.slice(0, 200)}{" "}
            <span
              onClick={toggleShowMore}
              className="cursor-pointer text-indigo-500"
            >
              Read more...
            </span>
          </>
        )}
        {showMore}
      </p>
    </div>
  );
}
