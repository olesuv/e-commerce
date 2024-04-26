import Cookies from "js-cookie";

interface LogoutButtonProps {
  setAuthenticated: (value: boolean) => void;
}

export default function LogoutButton(props: LogoutButtonProps) {
  return (
    <button
      onClick={() => {
        Cookies.remove("auth");
        props.setAuthenticated(false);
      }}
      className="text-white font-semibold rounded-lg bg-indigo-500 hover:bg-indigo-700 px-3 py-1"
    >
      Sign Out
    </button>
  );
}
