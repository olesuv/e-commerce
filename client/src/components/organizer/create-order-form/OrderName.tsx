export default function OrderName() {
  return (
    <div>
      <div className="after:content-['*'] after:ml-0.5 after:text-red-500">
        Order name
      </div>
      <input
        type="text"
        placeholder="Example: Samsung Galaxy S21 Ultra"
        className="p-2 border border-gray-200 outline-indigo-300 rounded-md w-full"
      />
    </div>
  );
}
