// import { NextResponse } from "next/server";

// // const AUTH_URL = process.env.NEXT_PUBLIC_AUTH_URL || "http://localhost:6000";
// // temp fix
// const AUTH_URL = "http://localhost:6000";

// export async function GET(
//   request: Request,
//   { params }: { params: Promise<{ kid: string }> }
// ) {
//   const { kid } = await params;
//   const authHeader = request.headers.get("Authorization");

//   try {
//     const res = await fetch(`${AUTH_URL}/auth/token/${kid}`, {
//       headers: {
//         Authorization: authHeader || "",
//       },
//     });

//     const data = await res.json();

//     if (!res.ok) {
//       return NextResponse.json(data, { status: res.status });
//     }

//     return NextResponse.json(data);
//   } catch {
//     return NextResponse.json({ error: "Auth service unreachable" }, { status: 503 });
//   }
// }

import { NextResponse } from "next/server";

export async function GET(
  request: Request,
  { params }: { params: Promise<{ kid: string }> }
) {
  const { kid } = await params;
  const authHeader = request.headers.get("Authorization");

  try {
    console.log("Proxying to:", `http://localhost:6000/auth/token/${kid}`);
    console.log("Auth header present:", !!authHeader);
    
    const res = await fetch(`http://localhost:6000/auth/token/${kid}`, {
      headers: {
        Authorization: authHeader || "",
      },
    });

    console.log("Auth response status:", res.status);
    const data = await res.json();
    console.log("Auth response data:", data);

    if (!res.ok) {
      return NextResponse.json(data, { status: res.status });
    }

    return NextResponse.json(data);
  } catch (err) {
    console.error("Fetch error:", err);
    return NextResponse.json({ error: "Auth service unreachable", detail: String(err) }, { status: 503 });
  }
}