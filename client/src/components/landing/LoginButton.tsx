import { useState } from "react";
import LoginPopup from "../login/LoginPopup";

interface ILoginButtonProps {
  setAuthenticated: (value: boolean) => void;
}

export default function LoginButton(props: ILoginButtonProps) {
  const [showPopup, setShowPopup] = useState(false);

  return (
    <>
      <button
        onClick={() => setShowPopup(true)}
        className="text-white font-semibold rounded-lg bg-indigo-500 hover:bg-indigo-700 px-3 py-1 ml-3"
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
