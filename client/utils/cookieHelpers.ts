import Cookies from "js-cookie";
import { decodeEmailFromToken } from "./JWTHelpers";
import { gql, useQuery } from "@apollo/client";

export function getCookie(): Boolean {
  if (Cookies.get("auth")) {
    return true;
  } else {
    return false;
  }
}

export function getEmailFromCookie(): string | null {
  if (getCookie()) {
    const cookieEmail = decodeEmailFromToken();
    if (cookieEmail !== null) {
      return cookieEmail;
    }
  }

  return null;
}

const CHECK_USER_EMAIL = gql`
  query checkUserEmail($email: String!) {
    user(email: $email) {
      email
    }
  }
`;

export function verifyEmailFromCookie(): Boolean {
  const cookieEmail = getEmailFromCookie();

  if (cookieEmail !== null) {
    const { data } = useQuery(CHECK_USER_EMAIL, {
      variables: {
        email: cookieEmail,
      },
    });
    if (data?.user?.email === cookieEmail) {
      return true;
    }
  }

  return false;
}
