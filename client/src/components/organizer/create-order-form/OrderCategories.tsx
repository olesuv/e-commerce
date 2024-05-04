import { OrderCategory } from "../../../../types/orderTypes";

interface OrderCategoriesProps {
  orderCategories: OrderCategory[];
  setOrderCategories: (value: OrderCategory[]) => void;
}

function addCategory(category: OrderCategory, categories: OrderCategory[]) {
  if (categories.includes(category)) {
    return categories.filter((cat) => cat !== category);
  } else {
    return [...categories, category];
  }
}

export default function OrderCategories(props: OrderCategoriesProps) {
  return (
    <div onChange={() => console.log(props.orderCategories)}>
      <p className="inline-block">
        <span className="font-medium">Category</span>
        <span className="pl-1 text-xs text-gray-400">
          {props.orderCategories.length === 0 ? (
            <span>
              You can select <span className="underline">multiple</span>{" "}
              (default: <span className="italic">Other</span>)
            </span>
          ) : (
            <span>
              Selected:{" "}
              <span className="font-semibold">
                {props.orderCategories.length}
              </span>{" "}
              category / categories
            </span>
          )}
        </span>
      </p>

      <ul className="flex flex-wrap *:rounded-full *:border *:px-2 *:py-0.5 *:mr-1 *:mt-1">
        <li
          className="border-sky-100 bg-sky-50 text-sky-500 hover:border-sky-200 hover:bg-sky-100 cursor-pointer"
          onClick={() =>
            props.setOrderCategories(
              addCategory(OrderCategory.Electronics, props.orderCategories)
            )
          }
        >
          Electronics
        </li>
        <li
          className="border-pink-100 bg-pink-50 text-pink-500 hover:border-pink-200 hover:bg-pink-100 cursor-pointer"
          onClick={() =>
            props.setOrderCategories(
              addCategory(OrderCategory.Fashion, props.orderCategories)
            )
          }
        >
          Fashion
        </li>
        <li
          className="border-orange-100 bg-orange-50 text-orange-500 hover:border-orange-200 hover:bg-orange-100 cursor-pointer"
          onClick={() =>
            props.setOrderCategories(
              addCategory(OrderCategory.Home, props.orderCategories)
            )
          }
        >
          Home
        </li>
        <li
          className="border-indigo-100 bg-indigo-50 text-indigo-500 hover:border-indigo-200 hover:bg-indigo-100 cursor-pointer"
          onClick={() =>
            props.setOrderCategories(
              addCategory(OrderCategory.Sports, props.orderCategories)
            )
          }
        >
          Sports
        </li>
        <li
          className="border-green-100 bg-green-50 text-green-500 hover:border-green-200 hover:bg-green-100 cursor-pointer"
          onClick={() =>
            props.setOrderCategories(
              addCategory(OrderCategory.Books, props.orderCategories)
            )
          }
        >
          Books
        </li>
        <li
          className="border-emerald-100 bg-emerald-50 text-emerald-500 hover:border-emerald-200 hover:bg-emerald-100 cursor-pointer"
          onClick={() =>
            props.setOrderCategories(
              addCategory(OrderCategory.Automotive, props.orderCategories)
            )
          }
        >
          Automotive
        </li>
        <li
          className="border-neutral-100 bg-neutral-50 text-neutral-500 hover:border-neutral-200 hover:bg-neutral-100 cursor-pointer"
          onClick={() =>
            props.setOrderCategories(
              addCategory(OrderCategory.Other, props.orderCategories)
            )
          }
        >
          Other
        </li>
      </ul>
    </div>
  );
}
