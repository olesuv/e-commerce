import Cookies from "js-cookie";

type Cookie = {
  userEmail: string;
  exp: string;
};

function CheckCookies() {
  const cookies = Cookies.get("auth");
  if (cookies === undefined) {
    return false;
  }
  return cookies;
}

export default function Header() {
  return <>{CheckCookies() ? <h1>Logged in</h1> : <h1>Not logged in</h1>}</>;
}
