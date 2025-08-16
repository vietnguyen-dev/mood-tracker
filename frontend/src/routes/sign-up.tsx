import { createFileRoute } from "@tanstack/react-router";
import { SignUp } from "@clerk/clerk-react";
export const Route = createFileRoute("/sign-up")({
  component: Signup,
});

function Signup() {
  return (
    <div className="flex place-items-center">
      <SignUp />
    </div>
  );
}
