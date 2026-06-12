export const TOKEN_COOKIE = "atlas_token";

export function decodeToken(token: string) {
  try {
    const payload = token.split(".")[1];
    return JSON.parse(atob(payload));
  } catch {
    return null;
  }
}

export function isTokenExpired(token: string): boolean {
  const claims = decodeToken(token);
  if (!claims) return true;
  return Date.now() >= claims.exp * 1000;
}

export function isAdmin(token: string): boolean {
  const claims = decodeToken(token);
  if (!claims) return false;
  return claims.Roles?.includes("ADMIN") ?? false;
}