import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { isTokenExpired } from "@/lib/auth";
import Sidebar from "@/app/components/layout/Sidebar";

export default async function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const cookieStore = await cookies();
  const token = cookieStore.get("atlas_token")?.value;

  if (!token || isTokenExpired(token)) {
    redirect("/login");
  }

  return (
    <div className="flex h-screen overflow-hidden bg-gray-950">
      <Sidebar token={token} />
      <main className="flex-1 overflow-auto p-8">
        {children}
      </main>
    </div>
  );
}