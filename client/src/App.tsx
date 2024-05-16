import "./App.css";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";
import Header from "./components/landing/Header";
import Organizer from "./components/organizer/Organizer";
import LatestOrders from "./components/landing/LatestOrders";

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
      <Organizer authenticated={authenticated} />
      <LatestOrders />
    </>
  );
}
