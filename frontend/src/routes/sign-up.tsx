import { createFileRoute, useNavigate, Link } from "@tanstack/react-router";
import { useState } from "react";
import type { FormEvent } from "react";
import logo from "../assets/logo.png";

export const Route = createFileRoute("/sign-up")({
  component: Signup,
});

function Signup() {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const navigate = useNavigate();

  const submitDisable = email.length > 0 && password.length > 0;

  const isEmailValid = (email: string) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };

  let emailcheck = isEmailValid(email);

  const changeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  };

  const changePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  return (
    <div className="flex place-items-center">
      <div className="flex w-full flex-col">
        <div className="card rounded-box place-items-center p-4">
          <h1 className="font-bold text-3xl my-3">Moodtrackerfor.me</h1>
          <img src={logo} className="w-24 mb-12" />
          <h2 className="font-bold text-2xl">Sign Up</h2>
          <form>
            <fieldset className="fieldset mb-3">
              <legend className="fieldset-legend">Email</legend>
              <input
                type="text"
                className="input"
                placeholder="email@domain.com"
                value={email}
                onChange={changeEmail}
              />
              {email.length > 0 && !emailcheck && (
                <p className="label text-red-400">Must be proper email</p>
              )}
            </fieldset>
            <fieldset className="fieldset mb-3">
              <legend className="fieldset-legend">Password</legend>
              <input
                type="text"
                className="input"
                placeholder="enter password here"
                value={password}
                onChange={changePassword}
              />
              <p className="label">Optional</p>
            </fieldset>
            <button
              className="btn btn-primary w-full my-3"
              disabled={!submitDisable}
            >
              Sign In
            </button>
          </form>
          <Link to="/login" className="underline">
            Login Here
          </Link>
        </div>
      </div>
    </div>
  );
}
