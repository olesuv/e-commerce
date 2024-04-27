import Cookies from "js-cookie";
import { useState } from "react";
import { gql, useMutation } from "@apollo/client";

interface ILoginPopupProps {
  setAuthenticated: (value: boolean) => void;
  setShowPopup: (value: boolean) => void;
}

const LOGIN = gql`
  mutation login($input: LoginUserInput!) {
    loginUser(input: $input)
  }
`;

export default function LoginPopup(props: ILoginPopupProps) {
  const [userEmail, setEmail] = useState("");
  const [userPassword, setPassword] = useState("");

  const [loginMutation, { loading, error, data }] = useMutation(LOGIN, {
    variables: {
      input: {
        email: userEmail,
        password: userPassword,
      },
    },
  });

  if (data) {
    Cookies.set("auth", data.loginUser);
    props.setAuthenticated(true);
    props.setShowPopup(false);
  }

  return (
    <div className="fixed top-0 left-0 w-full h-full bg-gray-800 bg-opacity-60 flex justify-center items-center">
      <div className="bg-white p-8 rounded-lg">
        <p className="text-2xl font-serif font-bold text-center mb-4">
          Welcome back!
        </p>
        {error && (
          <label className="flex flex-col bg-rose-500 w-full rounded-md my-2">
            <p className="text-xl text-white text-center font-semibold p-5">
              Oops, {error.message}
            </p>
          </label>
        )}
        <form
          onSubmit={(e) => {
            e.preventDefault();
            loginMutation();
          }}
          className="flex flex-col gap-4"
        >
          <input
            type="email"
            placeholder="Email"
            value={userEmail}
            onChange={(e) => setEmail(e.target.value)}
            className="p-2 border border-gray-200 rounded-md outline-indigo-300"
          />
          <input
            type="password"
            placeholder="Password"
            value={userPassword}
            onChange={(e) => setPassword(e.target.value)}
            className="p-2 border border-gray-200 rounded-md outline-indigo-300"
          />
          {loading ? (
            <button
              type="submit"
              className="p-2 bg-indigo-400 text-white rounded-md cursor-not-allowed"
              disabled
            >
              Loading...
            </button>
          ) : (
            <button
              type="submit"
              className="p-2 bg-indigo-500 hover:bg-indigo-700 text-white rounded-md"
            >
              Login
            </button>
          )}
        </form>
      </div>
    </div>
  );
}
