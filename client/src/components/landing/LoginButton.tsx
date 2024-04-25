import { useState } from "react";
import LoginPopup from "../login/LoginPopup";

export default function LoginButton() {
  const [showPopup, setShowPopup] = useState(false);

  return (
    <>
      <button
        onClick={() => setShowPopup(true)}
        className="text-white font-semibold rounded-lg bg-indigo-500 hover:bg-indigo-700 px-3 py-1 ml-3"
      >
        Sign In
      </button>
      {showPopup && <LoginPopup />}
    </>
  );
}
