"use client";

import React, { useState } from "react";
import Image from "next/image";
import Link from "next/link";

const page = () => {
  const [isBlurred, setIsBlurred] = useState(false);

  const handleFormClick = (): void => {
    setIsBlurred(true);
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
            onClick={handleFormClick}
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
            <label htmlFor="password" className="text-white font-ligh py-2">
              Password<span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              id="password"
              name="password"
              placeholder="••••••••••"
              required
              className="w-80 h-10 px-4 mb-8"
            />
            <label htmlFor="password" className="text-white font-ligh py-2">
              Re-enter Password<span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              id="password"
              name="password"
              placeholder="••••••••••"
              required
              className="w-80 h-10 px-4 mb-8"
            />
            <input
              type="submit"
              value="SIGN UP"
              className="text-white font-bold w-24 h-10 px-4 border border-white my-16"
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
        className={`w-1/2 overflow-hidden ${
          isBlurred
            ? "filter blur-sm scale-105 transition-all duration-10000"
            : ""
        }`}
        style={{
          backgroundImage: `url('/backgrounds/auth-bg.png')`,
          backgroundSize: "cover",
          backgroundPosition: "right",
        }}
      ></div>
    </div>
  );
};

export default page;
