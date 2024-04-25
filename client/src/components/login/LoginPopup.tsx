import { useState } from "react";
import { gql, useMutation } from "@apollo/client";

const LOGIN = gql`
  mutation login($input: Any!) {
    loginUser(input: input)
  }
`;

export default function LoginPopup() {
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

  if (loading) console.log("Loading...");
  if (error) console.error(error);
  if (data) {
    console.log(data.loginUser);
    // Cookies.set("auth", data.login.token);
  }

  return (
    <div className="fixed top-0 left-0 w-full h-full bg-gray-800 bg-opacity-60 flex justify-center items-center">
      <div className="bg-white p-8 rounded-lg">
        <h1 className="text-2xl font-bold text-center">Login</h1>
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
            className="p-2 border border-gray-300 rounded-md"
          />
          <input
            type="password"
            placeholder="Password"
            value={userPassword}
            onChange={(e) => setPassword(e.target.value)}
            className="p-2 border border-gray-300 rounded-md"
          />
          <button
            type="submit"
            className="p-2 bg-blue-500 text-white rounded-md"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  );
}
