import "./App.css";
import Cookies from "js-cookie";
import Header from "./components/landing/Header";
import { useEffect, useState } from "react";

export default function App() {
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
    <>
      <Header
        authenticated={authenticated}
        setAuthenticated={setAuthenticated}
      />
      <h1>wassup ladies</h1>
      <h6>(i know u got no bitches)</h6>
    </>
  );
}
