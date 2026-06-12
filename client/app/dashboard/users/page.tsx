import { cookies } from "next/headers";
import { getUsers } from "@/lib/api";
import Link from "next/link";
import type { User } from "@/types";

export default async function UsersPage() {
  const cookieStore = await cookies();
  const token       = cookieStore.get("atlas_token")?.value || "";

  let users: User[] = [];
  let error: string | null = null;

  try {
    users = await getUsers(token, 1, 20);
  } catch (err) {
    error = (err as { error: string }).error || "Failed to load users";
  }

  return (
    <div className="max-w-4xl">

      {/* Header */}
      <div className="flex items-center justify-between mb-8">
        <div>
          <p className="text-gray-500 text-xs font-mono tracking-widest uppercase mb-1">
            Management
          </p>
          <h1 className="text-xl font-mono font-semibold text-white">
            Users
          </h1>
        </div>
        <Link
          href="/dashboard/users/new"
          className="bg-blue-600 hover:bg-blue-500 text-white font-mono text-xs px-3 py-2 rounded transition-colors"
        >
          + New User
        </Link>
      </div>

      {error && (
        <div className="bg-gray-900 border border-red-800 rounded p-4 mb-6">
          <p className="text-red-400 text-sm font-mono">{error}</p>
        </div>
      )}

      {/* Table */}
      <div className="bg-gray-900 border border-gray-800 rounded overflow-hidden">
        <table className="w-full text-sm font-mono">
          <thead>
            <tr className="border-b border-gray-800">
              <Th>Name</Th>
              <Th>Email</Th>
              <Th>Roles</Th>
              <Th>Status</Th>
              <th className="px-4 py-3" />
            </tr>
          </thead>
          <tbody>
            {users.length === 0 && !error && (
              <tr>
                <td
                  colSpan={5}
                  className="px-4 py-8 text-center text-gray-600 text-xs font-mono"
                >
                  No users found
                </td>
              </tr>
            )}
            {users.map((user) => (
              <tr
                key={user.id}
                className="border-b border-gray-800 last:border-0 hover:bg-gray-950 transition-colors"
              >
                <td className="px-4 py-3 text-white">{user.name}</td>
                <td className="px-4 py-3 text-gray-400">{user.email}</td>
                <td className="px-4 py-3">
                  <div className="flex gap-1">
                    {user.roles?.map((role) => (
                      <span
                        key={role}
                        className={`text-xs px-2 py-0.5 rounded ${
                          role === "ADMIN"
                            ? "bg-blue-900/40 text-blue-400"
                            : "bg-gray-800 text-gray-400"
                        }`}
                      >
                        {role}
                      </span>
                    ))}
                  </div>
                </td>
                <td className="px-4 py-3">
                  <span className={`text-xs font-mono ${
                    user.enabled ? "text-green-500" : "text-red-400"
                  }`}>
                    {user.enabled ? "Active" : "Disabled"}
                  </span>
                </td>
                <td className="px-4 py-3 text-right">
                  <Link
                    href={`/dashboard/users/${user.id}`}
                    className="text-blue-500 hover:text-blue-400 text-xs transition-colors"
                  >
                    View →
                  </Link>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

    </div>
  );
}

function Th({ children }: { children: React.ReactNode }) {
  return (
    <th className="text-left text-gray-500 text-xs px-4 py-3 uppercase tracking-widest font-normal">
      {children}
    </th>
  );
}