import { createRootRoute, Outlet, useNavigate } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { useAuth, ClerkLoaded, ClerkLoading } from "@clerk/clerk-react";
import { useEffect } from "react";
import Navbar from "../components/navbar";

// Create a component that handles the authentication and layout
function RootLayout() {
  const { isLoaded, userId } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    // Check if Clerk has loaded and there's no logged-in user
    if (isLoaded && !userId) {
      navigate({ to: "/login" });
    }
  }, [isLoaded, userId, navigate]);

  return (
    <>
      <ClerkLoading>
        {/* You can show a custom loading spinner or message here */}
        <div>Loading...</div>
      </ClerkLoading>
      <ClerkLoaded>
        <Navbar />
        <div className="md:w-[75%] md:ml-[12.5%]">
          <Outlet />
        </div>
      </ClerkLoaded>
      <TanStackRouterDevtools />
    </>
  );
}

export const Route = createRootRoute({
  component: RootLayout,
});
