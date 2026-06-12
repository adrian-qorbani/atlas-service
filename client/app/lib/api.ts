import type { User, NewUser, UpdateUser, APIError } from "../types";

const API_URL  = process.env.NEXT_PUBLIC_API_URL  || "http://localhost:3000";
const AUTH_URL = process.env.NEXT_PUBLIC_AUTH_URL || "http://localhost:6000";

async function apiFetch<T>(
  baseUrl: string,
  path: string,
  options: RequestInit = {},
  token?: string
): Promise<T> {
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...(options.headers as Record<string, string>),
  };

  if (token) headers["Authorization"] = `Bearer ${token}`;

  const res = await fetch(`${baseUrl}${path}`, { ...options, headers });

  if (!res.ok) {
    const err: APIError = await res.json().catch(() => ({ error: `HTTP ${res.status}` }));
    throw err;
  }

  if (res.status === 204) return undefined as T;
  return res.json();
}

export async function getToken(email: string, password: string, kid: string): Promise<string> {
  const credentials = btoa(`${email}:${password}`);
  const res = await fetch(`${AUTH_URL}/auth/token/${kid}`, {
    headers: { Authorization: `Basic ${credentials}` },
  });

  if (!res.ok) {
    const err: APIError = await res.json().catch(() => ({ error: "Login failed" }));
    throw err;
  }

  const data = await res.json();
  return data.token;
}

export async function getUsers(token: string, page = 1, rows = 20): Promise<User[]> {
  return apiFetch(API_URL, `/users?page=${page}&rows=${rows}`, {}, token);
}

export async function getUserByID(token: string, userID: string): Promise<User> {
  return apiFetch(API_URL, `/users/${userID}`, {}, token);
}

export async function createUser(token: string, user: NewUser): Promise<User> {
  return apiFetch(API_URL, "/users", { method: "POST", body: JSON.stringify(user) }, token);
}

export async function updateUser(token: string, userID: string, update: UpdateUser): Promise<User> {
  return apiFetch(API_URL, `/users/${userID}`, { method: "PUT", body: JSON.stringify(update) }, token);
}

export async function deleteUser(token: string, userID: string): Promise<void> {
  return apiFetch(API_URL, `/users/${userID}`, { method: "DELETE" }, token);
}

export async function getLiveness(): Promise<unknown> {
  return apiFetch(API_URL, "/liveness");
}