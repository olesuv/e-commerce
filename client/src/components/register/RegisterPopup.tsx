import Cookies from "js-cookie";
import { gql, useMutation } from "@apollo/client";
import { useState } from "react";

const REGISTER = gql`
  mutation register($input: CreateUserInput!) {
    createUser(input: $input) {
      email
    }
  }
`;

interface RegisterPopupProps {
  setAuthenticated: (authenticated: boolean) => void;
  setShowPopup: (showPopup: boolean) => void;
}

export default function RegisterPopup(props: RegisterPopupProps) {
  const [userEmail, setEmail] = useState("");
  const [userPassword, setPassword] = useState("");
  const [userConfirmPassword, setConfirmPassword] = useState("");
  const [userName, setName] = useState("");

  const [error, setError] = useState<string>("");

  const [registerMutation, { loading }] = useMutation(REGISTER, {
    onError: (error) => {
      setError(error.message);
    },
    onCompleted: (data) => {
      Cookies.set("auth", data.createUser.email);
      props.setAuthenticated(true);
      props.setShowPopup(false);
    },
  });

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (userPassword !== userConfirmPassword) {
      setError("passwords do not match");
      return;
    }

    registerMutation({
      variables: {
        input: {
          email: userEmail,
          password: userPassword,
          name: userName,
        },
      },
    });
  };

  return (
    <div className="fixed top-0 left-0 w-full h-full bg-gray-800 bg-opacity-60 flex justify-center items-center">
      <div className="bg-white p-8 rounded-lg">
        <p className="text-2xl font-serif font-bold text-center mb-4">
          Welcome to <span className="text-3xl inline-block">e-commerce</span>
        </p>
        {error && (
          <label className="flex flex-col bg-rose-500 w-full rounded-md my-2">
            <p className="text-xl text-white text-center font-semibold p-5">
              Oops, {error}
            </p>
          </label>
        )}
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <input
            type="text"
            placeholder="Name"
            value={userName}
            onChange={(e) => setName(e.target.value)}
            className="p-2 border border-gray-200 rounded-md outline-indigo-300"
          />
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
          <input
            type="password"
            placeholder="Confirm Password"
            value={userConfirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
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
              Sign Up
            </button>
          )}
        </form>
      </div>
    </div>
  );
}
