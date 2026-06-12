"use client";

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { isAdmin, decodeToken } from "@/lib/auth";

interface SidebarProps {
  token: string;
}

export default function Sidebar({ token }: SidebarProps) {
  const pathname = usePathname();
  const router   = useRouter();
  const admin    = isAdmin(token);
  const claims   = decodeToken(token);

  async function handleLogout() {
    await fetch("/api/auth/logout", { method: "POST" });
    router.push("/login");
  }

  const navItems = [
    { href: "/dashboard",       label: "Overview"           },
    { href: "/dashboard/users", label: "Users", adminOnly: true },
  ];

  return (
    <aside className="w-52 shrink-0 bg-gray-900 border-r border-gray-800 flex flex-col">

      {/* Logo */}
      <div className="px-4 py-5 border-b border-gray-800">
        <span className="font-mono text-sm font-semibold text-white">
          atlas<span className="text-blue-500">.</span>service
        </span>
      </div>

      {/* Nav */}
      <nav className="flex-1 px-2 py-4 space-y-0.5">
        {navItems.map((item) => {
          if (item.adminOnly && !admin) return null;
          const active =
            pathname === item.href ||
            (item.href !== "/dashboard" && pathname.startsWith(item.href));

          return (
            <Link
              key={item.href}
              href={item.href}
              className={`block px-3 py-2 rounded text-xs font-mono transition-colors ${
                active
                  ? "bg-gray-950 text-white"
                  : "text-gray-500 hover:text-white hover:bg-gray-950"
              }`}
            >
              {item.label}
            </Link>
          );
        })}
      </nav>

      {/* User + logout */}
      <div className="px-4 py-4 border-t border-gray-800">
        <p className="text-xs font-mono text-gray-500 truncate mb-2">
          {claims?.Roles?.join(", ") || "—"}
        </p>
        <button
          onClick={handleLogout}
          className="text-xs font-mono text-gray-500 hover:text-red-400 transition-colors"
        >
          Sign out
        </button>
      </div>

    </aside>
  );
}