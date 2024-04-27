import RegisterPopup from "../register/RegisterPopup";
import { useState } from "react";

interface RegisterButtonProps {
  setAuthenticated: (authenticated: boolean) => void;
}

export default function RegisterButton(props: RegisterButtonProps) {
  const [showPopup, setShowPopup] = useState(false);

  return (
    <>
      <button
        onClick={() => setShowPopup(true)}
        className="text-white font-semibold rounded-lg bg-indigo-500 hover:bg-indigo-700 px-3 py-1"
      >
        Sign Up
      </button>
      {showPopup && (
        <RegisterPopup
          setAuthenticated={props.setAuthenticated}
          setShowPopup={setShowPopup}
        />
      )}
    </>
  );
}
