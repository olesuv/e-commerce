interface ICreateOrderPopupProps {
  setShowPopup: (value: boolean) => void;
}

export default function CreateOrderPopup(props: ICreateOrderPopupProps) {
  return (
    <div className="fixed top-0 left-0 w-full h-full bg-gray-800 bg-opacity-60 flex justify-center items-center">
      <div className="bg-white p-8 rounded-lg">
        <button onClick={() => props.setShowPopup(false)}>Hello</button>
      </div>
    </div>
  );
}
