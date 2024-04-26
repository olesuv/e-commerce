import Cookies from "js-cookie";
import LoginButton from "./LoginButton";
import RegisterButton from "./RegisterButton";
import LogoutButton from "./LogoutButton";
import Logo from "./Logo";

export default function Header() {
  const authenticated = Cookies.get("auth");

  return (
    <div className="md:grid md:grid-cols-3 p-4 bg-gray-200 text-black">
      <div className="hidden md:block"></div>
      <div className="flex justify-between md:items-center">
        <Logo />
        <div>
          {authenticated ? (
            <LogoutButton />
          ) : (
            <>
              <RegisterButton />
              <LoginButton />
            </>
          )}
        </div>
      </div>
      <div className="hidden md:block"></div>
    </div>
  );
}
