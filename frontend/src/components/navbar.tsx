import { Link } from "@tanstack/react-router";
import { SignedIn, SignedOut, UserButton, useUser } from "@clerk/clerk-react";
import ThemeSwitch from "./theme";

const Navbar = () => {
  const user = useUser();
  return (
    <nav className="flex justify-between items-center gap-3 p-6 md:w-[75%] md:ml-[12.5%]">
      <SignedOut>
        <div className="flex gap-5">
          <Link to="/" className="[&.active]:font-bold">
            Home
          </Link>{" "}
          <Link to="/login" className="[&.active]:font-bold">
            Login
          </Link>
          <Link to="/sign-up" className="[&.active]:font-bold">
            Sign Up
          </Link>
        </div>
        <ThemeSwitch />
      </SignedOut>
      <SignedIn>
        <Link to="/admin" className="[&.active]:font-bold">
          Hi, {user.user?.username}!
        </Link>
        <div className="flex gap-6">
          <ThemeSwitch />
          <UserButton />
        </div>
      </SignedIn>
    </nav>
  );
};

export default Navbar;
