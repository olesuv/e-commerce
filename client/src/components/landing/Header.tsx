import Cookies from "js-cookie";
import LoginButton from "./LoginButton";

function CheckCookies() {
  const cookies = Cookies.get("auth");
  if (cookies === undefined) {
    return false;
  }
  return cookies;
}

export default function Header() {
  return <>{CheckCookies() ? <h1>Logged in</h1> : <LoginButton />}</>;
}
