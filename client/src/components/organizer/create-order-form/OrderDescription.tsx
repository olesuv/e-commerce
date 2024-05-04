interface OrderDescriptionProps {
  setOrderDescription: (value: string) => void;
}

export default function OrderDescription(props: OrderDescriptionProps) {
  return (
    <div>
      <p className="font-medium after:content-['*'] after:ml-0.5 after:text-red-500">
        Description
      </p>
      <textarea
        onChange={(e) => props.setOrderDescription(e.target.value)}
        placeholder="Example: Brand new Samsung Galaxy S21 Ultra, 256GB, 12GB RAM, Phantom Black"
        className="p-2 border border-gray-200 outline-indigo-300 rounded-md h-36 w-full text-left"
      />
    </div>
  );
}
