import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "standalone",
  env: {
    NEXT_PUBLIC_API_URL:  process.env.NEXT_PUBLIC_API_URL  || "http://localhost:3000",
    NEXT_PUBLIC_AUTH_URL: process.env.NEXT_PUBLIC_AUTH_URL || "http://localhost:6000",
    NEXT_PUBLIC_AUTH_KID: process.env.NEXT_PUBLIC_AUTH_KID || "54bb2165-71e1-41a6-af3e-7da4a0e1e2c1",
  },
};

export default nextConfig;