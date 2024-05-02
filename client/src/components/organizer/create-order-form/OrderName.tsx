interface OrderNameProps {
  setOrderName: (value: string) => void;
}

export default function OrderName(props: OrderNameProps) {
  return (
    <div>
      <div className="after:content-['*'] after:ml-0.5 after:text-red-500 font-medium">
        Order name
      </div>
      <input
        onChange={(e) => {
          props.setOrderName(e.target.value);
        }}
        type="text"
        placeholder="Example: Samsung Galaxy S21 Ultra"
        className="p-2 border border-gray-200 outline-indigo-300 rounded-md w-full"
      />
    </div>
  );
}
