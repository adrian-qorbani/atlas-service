import { cookies } from "next/headers";
import { decodeToken, isAdmin } from "@/lib/auth";
import { getLiveness } from "@/lib/api";

async function getHealthStatus() {
  try {
    await getLiveness();
    return true;
  } catch {
    return false;
  }
}

export default async function DashboardPage() {
  const cookieStore = await cookies();
  const token       = cookieStore.get("atlas_token")?.value || "";
  const claims      = decodeToken(token);
  const admin       = isAdmin(token);
  const alive       = await getHealthStatus();

  return (
    <div className="max-w-3xl">

      {/* Header */}
      <div className="mb-8">
        <p className="text-gray-500 text-xs font-mono tracking-widest uppercase mb-1">
          Overview
        </p>
        <h1 className="text-xl font-mono font-semibold text-white">
          Atlas Service
        </h1>
      </div>

      {/* Status cards */}
      <div className="grid grid-cols-3 gap-4 mb-8">
        <StatusCard label="Sales API"     status={alive} detail="port 3000" />
        <StatusCard label="Auth Service"  status={true}  detail="port 6000" />
        <StatusCard label="Database"      status={true}  detail="postgres"  />
      </div>

      {/* Session */}
      <div className="bg-gray-900 border border-gray-800 rounded p-4">
        <p className="text-xs text-gray-500 font-mono mb-3 uppercase tracking-widest">
          Current Session
        </p>
        <div className="space-y-2 font-mono text-sm">
          <Row label="Subject" value={claims?.sub  || "—"} />
          <Row label="Roles"   value={claims?.Roles?.join(", ") || "—"} />
          <Row label="Issuer"  value={claims?.iss  || "—"} />
          <Row label="Access"  value={admin ? "Admin" : "User"} />
        </div>
      </div>

    </div>
  );
}

function StatusCard({
  label,
  status,
  detail,
}: {
  label: string;
  status: boolean;
  detail: string;
}) {
  return (
    <div className="bg-gray-900 border border-gray-800 rounded p-4">
      <div className="flex items-center justify-between mb-2">
        <span className="text-xs font-mono text-gray-500 uppercase tracking-widest">
          {label}
        </span>
        <span className={`w-2 h-2 rounded-full ${status ? "bg-green-500" : "bg-red-500"}`} />
      </div>
      <p className="text-xs font-mono text-gray-600">{detail}</p>
    </div>
  );
}

function Row({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex gap-4">
      <span className="text-gray-500 w-20 shrink-0">{label}</span>
      <span className="text-white truncate">{value}</span>
    </div>
  );
}