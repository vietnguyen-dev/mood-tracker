import { SignIn } from "@clerk/clerk-react";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/login")({
  component: Login,
});

function Login() {
  return (
    <div className="flex place-items-center">
      <SignIn fallbackRedirectUrl={"/admin"} />
    </div>
  );
}
