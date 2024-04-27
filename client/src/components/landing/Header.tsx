import Cookies from "js-cookie";
import LoginButton from "./LoginButton";
import RegisterButton from "./RegisterButton";
import LogoutButton from "./LogoutButton";
import Logo from "./Logo";
import { useEffect, useState } from "react";

export default function Header() {
  const [authenticated, setAuthenticated] = useState(!!Cookies.get("auth"));

  useEffect(() => {
    const checkAuth = () => {
      const isAuth = !!Cookies.get("auth");
      setAuthenticated(isAuth);
    };
    checkAuth();
    return () => {};
  }, []);

  return (
    <div className="md:grid md:grid-cols-3 p-4 bg-gray-200 text-black">
      <div className="hidden md:block"></div>
      <div className="flex justify-between md:items-center">
        <Logo />
        <div>
          {authenticated ? (
            <LogoutButton setAuthenticated={setAuthenticated} />
          ) : (
            <>
              <RegisterButton setAuthenticated={setAuthenticated} />
              <LoginButton setAuthenticated={setAuthenticated} />
            </>
          )}
        </div>
      </div>
      <div className="hidden md:block"></div>
    </div>
  );
}
