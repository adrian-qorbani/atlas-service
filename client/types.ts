export type Role = "ADMIN" | "USER";

export interface User {
  id: string;
  name: string;
  email: string;
  roles: Role[];
  enabled: boolean;
  department: string;
  dateCreated: string;
  dateUpdated: string;
}

export interface NewUser {
  name: string;
  email: string;
  password: string;
  password_confirm: string;
  roles: Role[];
  department: string;
}

export interface UpdateUser {
  name?: string;
  email?: string;
  roles?: Role[];
  enabled?: boolean;
  department?: string;
}

export interface APIError {
  error: string;
  fields?: { field: string; err: string }[];
}