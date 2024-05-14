import { IOrderDetailsProps } from "../../../types/orderDetailsProps";

export default function OrderCategory(props: IOrderDetailsProps) {
  return (
    <ul className="flex flex-wrap *:my-1 *:mr-1 *:rounded-full *:border *:px-2 *:py-0.5 *:text-sm">
      {props.orderData?.category.map((category: string) => {
        switch (category) {
          case "Electronics":
            return (
              <li
                key={category}
                className="cursor-pointer border-sky-100 bg-sky-50 text-sky-500 hover:border-sky-200 hover:bg-sky-100"
              >
                Electronics
              </li>
            );
          case "Fashion":
            return (
              <li
                key={category}
                className="cursor-pointer border-pink-100 bg-pink-50 text-pink-500 hover:border-pink-200 hover:bg-pink-100"
              >
                Fashion
              </li>
            );
          case "Home":
            return (
              <li
                key={category}
                className="cursor-pointer border-orange-100 bg-orange-50 text-orange-500 hover:border-orange-200 hover:bg-orange-100"
              >
                Home
              </li>
            );
          case "Sports":
            return (
              <li
                key={category}
                className="cursor-pointer border-indigo-100 bg-indigo-50 text-indigo-500 hover:border-indigo-200 hover:bg-indigo-100"
              >
                Sports
              </li>
            );
          case "Books":
            return (
              <li
                key={category}
                className="cursor-pointer border-green-100 bg-green-50 text-green-500 hover:border-green-200 hover:bg-green-100"
              >
                Books
              </li>
            );
          default:
            return (
              <li
                key={category}
                className="cursor-pointer border-neutral-100 bg-neutral-50 text-neutral-500 hover:border-neutral-200 hover:bg-neutral-100"
              >
                Other
              </li>
            );
        }
      })}
    </ul>
  );
}
