import Cookies from "js-cookie";

export default function LogoutButton() {
  return (
    <button
      onClick={() => Cookies.remove("auth")}
      className="text-white font-semibold rounded-lg bg-indigo-500 hover:bg-indigo-700 px-3 py-1"
    >
      Sign Out
    </button>
  );
}
