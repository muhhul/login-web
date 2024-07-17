"use client";

import { useRouter } from "next/navigation";

const Dashboard = () => {
  const router = useRouter();

  const handleLogout = () => {
    router.push("/");
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-6 bg-[#1c1c1c]">
      <h2 className="text-3xl font-extrabold text-white">
        You Are Currently Logged In
      </h2>
      <button
        onClick={handleLogout}
        className="mt-4 w-full max-w-xs flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Log Out
      </button>
    </main>
  );
};

export default Dashboard;
