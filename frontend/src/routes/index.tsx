import { createFileRoute, redirect } from "@tanstack/react-router";
import { useUser } from "@clerk/clerk-react";
import { useEffect } from "react";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  const { isSignedIn } = useUser();

  useEffect(() => {
    if (isSignedIn) {
      redirect({ to: "/admin" });
    } else {
      redirect({ to: "/login" });
    }
  }, [isSignedIn]);

  return (
    <div className="p-2">
      <h3>Welcome Home!</h3>
      <button className="btn btn-primary">Click me</button>
    </div>
  );
}
