"use client";

import React, { useState } from "react";
import Image from "next/image";
import Link from "next/link";
import validatePassword from "@/app/lib/passwordValidator";

export default function Page() {
  //   const [isBlurred, setIsBlurred] = useState(false);

  //   const handleFormClick = (): void => {
  //     setIsBlurred(true);
  //   };

  const [password, setPassword] = useState<string>("");
  const [reenteredPassword, setReenteredPassword] = useState<string>("");
  const [passwordError, setPasswordError] = useState<string>("");
  const [reenteredPasswordError, setReenteredPasswordError] =
    useState<string>("");

  const handlePasswordChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ): void => {
    const newPassword = e.target.value;
    setPassword(newPassword);
    setPasswordError("");
    if (passwordError) setPasswordError("");
  };

  const handleReenteredPasswordChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ): void => {
    const newReenteredPassword = e.target.value;
    setReenteredPassword(newReenteredPassword);
    setReenteredPasswordError("");
    if (reenteredPasswordError) setReenteredPasswordError("");
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    const isValidPassword = validatePassword(password);
    const doPasswordsMatch = password === reenteredPassword;

    if (!isValidPassword) {
      setPasswordError(
        "Password must have at least one letter, one number, one special character, and at least 8 characters long."
      );
      return;
    }

    if (!doPasswordsMatch) {
      setReenteredPasswordError("Passwords do not match.");
      return;
    }

    console.log("Form submitted");
  };

  return (
    <div className="flex h-screen">
      <div className="h-screen bg-black py-8 px-32 w-1/2">
        <div className="flex items-center">
          <Image
            src="/backgrounds/NYU-logo.png"
            width="100"
            height="39"
            alt="NYU Logo"
            className="z-10"
          />
          <Image
            src="/backgrounds/darkspace.png"
            width="150"
            height="39"
            alt="Darkspace Logo"
            className="z-10"
          />
        </div>
        <div className="flex flex-col my-16">
          <h1 className="text-white font-bold text-3xl pb-8">Sign up</h1>
          <form
            action="login.php"
            method="post"
            className="flex flex-col"
            // onClick={handleFormClick}
            onSubmit={handleSubmit}
          >
            <label htmlFor="email" className="text-white font-light py-2">
              Email<span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              id="email"
              name="email"
              placeholder="abc123@nyu.edu"
              required
              className="w-80 h-10 px-4 mb-8"
            />
            <label htmlFor="password" className="text-white font-light py-2">
              Password<span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              id="password"
              name="password"
              placeholder="••••••••••"
              required
              value={password}
              onChange={handlePasswordChange}
              className={`w-80 h-10 px-4 mb-8 ${
                passwordError && "border-red-500"
              }`}
            />
            {passwordError && (
              <p className="text-red-500 pb-2">{passwordError}</p>
            )}
            <label
              htmlFor="reentered-password"
              className="text-white font-ligh py-2"
            >
              Re-enter Password<span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              id="reentered-password"
              name="reentered-password"
              placeholder="••••••••••"
              required
              className={`w-80 h-10 px-4 mb-8 ${
                reenteredPasswordError && "border-red-500"
              }`}
            />
            {reenteredPasswordError && (
              <p className="text-red-500 pb-2">{reenteredPasswordError}</p>
            )}
            <input
              type="submit"
              className="text-white font-bold w-40 h-10 px-4 border border-white my-16 hover:bg-gray-400 active:bg-white active:text-black"
            />
          </form>
          <h3 className="text-white font-light text-sm text-center">
            Already have an account?{" "}
            <Link href="login" className="underline">
              Log in
            </Link>
          </h3>
        </div>
      </div>
      <div
        // className={`w-1/2 overflow-hidden ${
        //   isBlurred
        //     ? "filter blur-sm scale-105 transition-all duration-10000"
        //     : ""
        // }`}
        className="w-1/2"
        style={{
          backgroundImage: `url('/backgrounds/auth-bg.png')`,
          backgroundSize: "cover",
          backgroundPosition: "right",
        }}
      ></div>
    </div>
  );
}
