"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";

const Home = () => {
  const [isSignIn, setIsSignIn] = useState(true);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const router = useRouter();

  const toggleForm = () => {
    setIsSignIn(!isSignIn);
    setError("");
    setEmail("");
    setPassword("");
    setConfirmPassword("");
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!isSignIn && password !== confirmPassword) {
      setError("Passwords do not match");
      return;
    }

    try {
      if (isSignIn) {
        await axios.post("http://localhost:8080/signin", {
          email,
          password,
        });
        router.push("/dashboard");
      } else {
        await axios.post("http://localhost:8080/signup", {
          email,
          password,
        });
        setEmail("");
        setPassword("");
        setConfirmPassword("");
        setError("");
        setIsSignIn(true);
        router.push("/");
      }
    } catch (error) {
      setError("An error occurred. Please try again.");
    }
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-6 bg-[#1c1c1c]">
      <div className="w-full max-w-md p-8 space-y-4">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-[#f4f4f5]">
          {isSignIn ? "Sign In" : "Sign Up"}
        </h2>
        <form className="mt-8" onSubmit={handleSubmit} method="POST">
          <input type="hidden" name="remember" defaultValue="true" />
          <div className="rounded-md shadow-sm -space-y-px">
            <div>
              <label htmlFor="email-address" className="sr-only">
                Email address
              </label>
              <input
                id="email-address"
                name="email"
                type="email"
                autoComplete="email"
                required
                className="appearance-none rounded-none relative block bg-[#f4f4f5] w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="Email address"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </div>
            <div>
              <label htmlFor="password" className="sr-only">
                Password
              </label>
              <input
                id="password"
                name="password"
                type="password"
                autoComplete="current-password"
                required
                className={`appearance-none relative block bg-[#f4f4f5] w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm ${
                  isSignIn ? "rounded-b-md" : ""
                }`}
                placeholder="Password"
                value={password}
                onChange={(e) => {
                  setPassword(e.target.value);
                  setError("");
                }}
              />
            </div>
            {!isSignIn && (
              <div>
                <label htmlFor="confirm-password" className="sr-only">
                  Confirm Password
                </label>
                <input
                  id="confirm-password"
                  name="confirmPassword"
                  type="password"
                  autoComplete="current-password"
                  required
                  className="appearance-none rounded-none relative block bg-[#f4f4f5] w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                  placeholder="Confirm Password"
                  value={confirmPassword}
                  onChange={(e) => {
                    setConfirmPassword(e.target.value);
                    setError("");
                  }}
                />
              </div>
            )}
          </div>

          {error && (
            <p className="mt-4 text-center text-sm text-red-600">{error}</p>
          )}

          {isSignIn && (
            <div className="mt-4 flex items-center justify-center">
              <div className="text-sm">
                <a
                  href="#"
                  className="font-medium text-indigo-600 hover:text-indigo-500"
                >
                  Forgot your password?
                </a>
              </div>
            </div>
          )}

          <div className="mt-4">
            <button
              type="submit"
              className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-[#f4f4f5] bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              {isSignIn ? "Sign In" : "Sign Up"}
            </button>
          </div>
        </form>
        <div className="mt-2 text-center">
          <p className="text-sm text-gray-600">
            {isSignIn ? "Don't have an account?" : "Already have an account?"}{" "}
            <a
              href="#"
              onClick={toggleForm}
              className="font-medium text-indigo-600 hover:text-indigo-500"
            >
              {isSignIn ? "Sign Up" : "Sign In"}
            </a>
          </p>
        </div>
      </div>
    </main>
  );
};

export default Home;
