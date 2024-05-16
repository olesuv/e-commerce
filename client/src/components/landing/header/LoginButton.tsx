import { useState } from "react";
import LoginPopup from "../../login/LoginPopup";

interface ILoginButtonProps {
  setAuthenticated: (value: boolean) => void;
}

export default function LoginButton(props: ILoginButtonProps) {
  const [showPopup, setShowPopup] = useState(false);

  return (
    <>
      <button
        onClick={() => setShowPopup(true)}
        className="ml-3 rounded-lg bg-indigo-500 px-3 py-1 font-semibold text-white hover:bg-indigo-700"
      >
        Sign In
      </button>
      {showPopup && (
        <LoginPopup
          setAuthenticated={props.setAuthenticated}
          setShowPopup={setShowPopup}
        />
      )}
    </>
  );
}
