import RegisterPopup from "../../register/RegisterPopup";
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
        className="rounded-lg bg-indigo-500 px-3 py-1 font-semibold text-white hover:bg-indigo-700"
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
