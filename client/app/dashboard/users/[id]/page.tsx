import { cookies } from "next/headers";
import { getUserByID } from "@/lib/api";
import { isAdmin } from "@/lib/auth";
import Link from "next/link";
import DeleteUserButton from "./DeleteUserButton";

export default async function UserDetailPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const cookieStore = await cookies();
  const token       = cookieStore.get("atlas_token")?.value || "";
  const { id }      = await params;
  const admin       = isAdmin(token);

  let user = null;
  let error: string | null = null;

  try {
    user = await getUserByID(token, id);
  } catch (err) {
    error = (err as { error: string }).error || "Failed to load user";
  }

  if (error) {
    return (
      <div className="max-w-2xl">
        <div className="bg-gray-900 border border-red-800 rounded p-4">
          <p className="text-red-400 text-sm font-mono">{error}</p>
        </div>
      </div>
    );
  }

  if (!user) return null;

  return (
    <div className="max-w-2xl">

      {/* Header */}
      <div className="flex items-center justify-between mb-8">
        <div>
          <Link
            href="/dashboard/users"
            className="text-gray-500 hover:text-white text-xs font-mono transition-colors"
          >
            ← Users
          </Link>
          <h1 className="text-xl font-mono font-semibold text-white mt-1">
            {user.name}
          </h1>
        </div>
        {admin && (
          <DeleteUserButton token={token} userID={user.id} />
        )}
      </div>

      {/* Details */}
      <div className="bg-gray-900 border border-gray-800 rounded p-4 space-y-3 font-mono text-sm mb-4">
        <Row label="ID"         value={user.id} />
        <Row label="Name"       value={user.name} />
        <Row label="Email"      value={user.email} />
        <Row label="Department" value={user.department || "—"} />
        <Row label="Roles"      value={user.roles?.join(", ") || "—"} />
        <Row label="Status"     value={user.enabled ? "Active" : "Disabled"} />
        <Row label="Created" value={new Date(user.dateCreated).toLocaleString()} />
        <Row label="Updated" value={new Date(user.dateUpdated).toLocaleString()} />
      </div>

    </div>
  );
}

function Row({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex gap-4">
      <span className="text-gray-500 w-28 shrink-0">{label}</span>
      <span className="text-white break-all">{value}</span>
    </div>
  );
}