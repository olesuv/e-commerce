interface OrderCategoriesProps {
  setOrderCategories: (value: number[]) => void;
}

export default function OrderCategories(props: OrderCategoriesProps) {
  return (
    <div>
      <p className="inline-block">
        <span className="after:content-['*'] after:ml-0.5 after:text-red-500 font-medium">
          Categories
        </span>
        <span className="pl-1 text-xs text-gray-400">
          (You can select <span className="underline">multiple</span>)
        </span>
      </p>

      <ul className="flex flex-wrap *:rounded-full *:border *:px-2 *:py-0.5 *:mr-1 *:mt-1">
        <li className="border-sky-100 bg-sky-50 text-sky-500 hover:border-sky-200 hover:bg-sky-100 cursor-pointer">
          Electronics
        </li>
        <li className="border-pink-100 bg-pink-50 text-pink-500 hover:border-pink-200 hover:bg-pink-100 cursor-pointer">
          Fashion
        </li>
        <li className="border-orange-100 bg-orange-50 text-orange-500 hover:border-orange-200 hover:bg-orange-100 cursor-pointer">
          Home
        </li>
        <li className="border-indigo-100 bg-indigo-50 text-indigo-500 hover:border-indigo-200 hover:bg-indigo-100 cursor-pointer">
          Sports
        </li>
        <li className="border-green-100 bg-green-50 text-green-500 hover:border-green-200 hover:bg-green-100 cursor-pointer">
          Books
        </li>
        <li className="border-emerald-100 bg-emerald-50 text-emerald-500 hover:border-emerald-200 hover:bg-emerald-100 cursor-pointer">
          Automotive
        </li>
        <li className="border-neutral-100 bg-neutral-50 text-neutral-500 hover:border-neutral-200 hover:bg-neutral-100 cursor-pointer">
          Other
        </li>
      </ul>
    </div>
  );
}
