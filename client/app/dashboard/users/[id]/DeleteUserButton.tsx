"use client";

import { useRouter } from "next/navigation";
import { deleteUser } from "@/lib/api";
import { useState } from "react";

export default function DeleteUserButton({
  token,
  userID,
}: {
  token: string;
  userID: string;
}) {
  const router  = useRouter();
  const [loading, setLoading] = useState(false);
  const [confirm, setConfirm] = useState(false);

  async function handleDelete() {
    if (!confirm) {
      setConfirm(true);
      return;
    }

    setLoading(true);
    try {
      await deleteUser(token, userID);
      router.push("/dashboard/users");
    } catch (err) {
      console.error(err);
      setLoading(false);
      setConfirm(false);
    }
  }

  return (
    <button
      onClick={handleDelete}
      disabled={loading}
      className={`text-xs font-mono px-3 py-2 rounded transition-colors disabled:opacity-50 ${
        confirm
          ? "bg-red-600 hover:bg-red-500 text-white"
          : "border border-gray-700 text-gray-500 hover:text-red-400 hover:border-red-800"
      }`}
    >
      {loading ? "Deleting..." : confirm ? "Confirm delete" : "Delete"}
    </button>
  );
}