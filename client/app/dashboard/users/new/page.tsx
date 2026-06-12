"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { createUser } from "@/lib/api";
import type { APIError, Role } from "@/types";

const COOKIE_NAME = "atlas_token";

function getTokenFromCookie(): string {
  const match = document.cookie.match(new RegExp(`(^| )${COOKIE_NAME}=([^;]+)`));
  return match ? match[2] : "";
}

export default function NewUserPage() {
  const router = useRouter();

  const [name, setName]           = useState("");
  const [email, setEmail]         = useState("");
  const [password, setPassword]   = useState("");
  const [confirm, setConfirm]     = useState("");
  const [department, setDepartment] = useState("");
  const [roles, setRoles]         = useState<Role[]>(["USER"]);
  const [error, setError]         = useState<string | null>(null);
  const [fieldErrors, setFieldErrors] = useState<Record<string, string>>({});
  const [loading, setLoading]     = useState(false);

  function toggleRole(role: Role) {
    setRoles((prev) =>
      prev.includes(role) ? prev.filter((r) => r !== role) : [...prev, role]
    );
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError(null);
    setFieldErrors({});

    if (password !== confirm) {
      setFieldErrors({ password_confirm: "Passwords do not match" });
      return;
    }

    setLoading(true);
    try {
      const token = getTokenFromCookie();
      await createUser(token, {
        name,
        email,
        password,
        password_confirm: confirm,
        roles,
        department,
      });
      router.push("/dashboard/users");
    } catch (err) {
      const apiErr = err as APIError;
      if (apiErr.fields) {
        const map: Record<string, string> = {};
        apiErr.fields.forEach((f) => { map[f.field] = f.err; });
        setFieldErrors(map);
      } else {
        setError(apiErr.error || "Failed to create user");
      }
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="max-w-lg">

      {/* Header */}
      <div className="mb-8">
        <button
          onClick={() => router.back()}
          className="text-gray-500 hover:text-white text-xs font-mono transition-colors"
        >
          ← Users
        </button>
        <h1 className="text-xl font-mono font-semibold text-white mt-1">
          New User
        </h1>
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">

        <Field
          label="Name"
          error={fieldErrors.name}
        >
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Jane Smith"
            required
            className={inputClass}
          />
        </Field>

        <Field label="Email" error={fieldErrors.email}>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="jane@example.com"
            required
            className={inputClass}
          />
        </Field>

        <Field label="Department" error={fieldErrors.department}>
          <input
            type="text"
            value={department}
            onChange={(e) => setDepartment(e.target.value)}
            placeholder="Engineering"
            className={inputClass}
          />
        </Field>

        <Field label="Password" error={fieldErrors.password}>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Min 8 characters"
            required
            className={inputClass}
          />
        </Field>

        <Field label="Confirm Password" error={fieldErrors.password_confirm}>
          <input
            type="password"
            value={confirm}
            onChange={(e) => setConfirm(e.target.value)}
            placeholder="••••••••"
            required
            className={inputClass}
          />
        </Field>

        {/* Roles */}
        <div>
          <label className="block text-xs text-gray-500 mb-2 font-mono">
            Roles
          </label>
          <div className="flex gap-2">
            {(["USER", "ADMIN"] as Role[]).map((role) => (
              <button
                key={role}
                type="button"
                onClick={() => toggleRole(role)}
                className={`text-xs font-mono px-3 py-1.5 rounded border transition-colors ${
                  roles.includes(role)
                    ? "border-blue-500 text-blue-400 bg-blue-900/20"
                    : "border-gray-700 text-gray-500 hover:border-gray-600"
                }`}
              >
                {role}
              </button>
            ))}
          </div>
        </div>

        {error && (
          <p className="text-red-400 text-xs font-mono">{error}</p>
        )}

        <div className="flex gap-3 pt-2">
          <button
            type="submit"
            disabled={loading || roles.length === 0}
            className="bg-blue-600 hover:bg-blue-500 disabled:opacity-50 disabled:cursor-not-allowed text-white font-mono text-sm py-2 px-4 rounded transition-colors"
          >
            {loading ? "Creating..." : "Create User"}
          </button>
          <button
            type="button"
            onClick={() => router.back()}
            className="border border-gray-700 text-gray-500 hover:text-white font-mono text-sm py-2 px-4 rounded transition-colors"
          >
            Cancel
          </button>
        </div>

      </form>
    </div>
  );
}

const inputClass =
  "w-full bg-gray-900 border border-gray-800 rounded px-3 py-2 text-sm font-mono text-white placeholder-gray-600 focus:outline-none focus:border-blue-500 transition-colors";

function Field({
  label,
  error,
  children,
}: {
  label: string;
  error?: string;
  children: React.ReactNode;
}) {
  return (
    <div>
      <label className="block text-xs text-gray-500 mb-1 font-mono">
        {label}
      </label>
      {children}
      {error && (
        <p className="text-red-400 text-xs font-mono mt-1">{error}</p>
      )}
    </div>
  );
}