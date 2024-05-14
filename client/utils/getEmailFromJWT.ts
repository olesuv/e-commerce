import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

export function decodeEmailFromToken(): string | null {
  const token = Cookies.get("auth");

  if (!token) return null;
  try {
    const decodedToken = jwtDecode(token);
    return decodedToken?.email;
  } catch (error) {
    console.error("Error decoding token:", error);
  }

  return null;
}
