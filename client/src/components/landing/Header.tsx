import Cookies from "js-cookie";
import LoginButton from "./LoginButton";
import RegisterButton from "./RegisterButton";
import Logo from "./Logo";

function checkCookies() {
  const cookies = Cookies.get("auth");
  if (cookies === undefined) {
    return false;
  }
  return cookies;
}

export default function Header() {
  return (
    <div className="md:grid md:grid-cols-3 p-4 bg-gray-200 text-black">
      {checkCookies() ? (
        <>wassup</>
      ) : (
        <>
          <div className="hidden md:block"></div>

          <div className="flex justify-between md:items-center">
            <Logo />
            <div>
              <RegisterButton />
              <LoginButton />
            </div>
          </div>

          <div className="hidden md:block"></div>
        </>
      )}
    </div>
  );
}
