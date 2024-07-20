"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";

const Dashboard = () => {
  const router = useRouter();
  const [error, setError] = useState("");

  const handleLogout = () => {
    router.push("/");
  };

  const handleDeleteAccount = async () => {
    try {
      const data = {};
      await axios.post("http://localhost:8000/api/deleteAccount", data, {
        headers: {
          "Content-Type": "application/json",
        },
      });
      router.push("/");
    } catch (error: any) {
      setError(error.response.data.message);
    }
  };

  const handleChangePassword = () => {
    router.push("/changePassword");
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-6 bg-[#1c1c1c]">
      <h2 className="text-3xl font-extrabold text-[#f4f4f5]">
        You Are Currently Logged In
      </h2>
      <button
        onClick={handleLogout}
        className="mt-4 w-full max-w-xs flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Log Out
      </button>
      <button
        onClick={handleChangePassword}
        className="mt-4 w-full max-w-xs flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        Change Password
      </button>
      <button
        onClick={handleDeleteAccount}
        className="mt-4 w-full max-w-xs flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
      >
        Delete Account
      </button>
      {error && (
        <p className="mt-4 text-center text-sm text-red-600">{error}</p>
      )}
    </main>
  );
};

export default Dashboard;
