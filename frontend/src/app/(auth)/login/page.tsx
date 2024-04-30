"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";
import Link from "next/link";

export default function Page() {
  // const [isBlurred, setIsBlurred] = useState(false);

  // const handleFormClick = (): void => {
  //   setIsBlurred(true);
  // };

  const [userLogin, setUserLogin] = useState({
    netid: "",
    password: "",
  });
  const [loginError, setLoginError] = useState<string>("");
  const route = useRouter();

  const loginUser = async () => {
    try {
      const res: Response = await fetch("http://localhost:6789/v1/user/login", {
        method: "POST",
        // headers: {
        //   "Content-Type": "application/json",
        //   "Access-Control-Allow-Origin": "*",
        //   "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
        //   "Access-Control-Allow-Headers": "Content-Type, Authorization",
        // },
        body: JSON.stringify({
          netid: userLogin.netid,
          password: userLogin.password,
        }),
      });
      if (res.ok) {
        const data = await res.json();
        localStorage.setItem("token", data.authentication_token.token);
        route.push("/");
        console.log("token: %s", data.authentication_token.token);
      } else {
        console.error("Failed to login user:", res.statusText);
        return [];
      }
    } catch (error) {
      console.error("Error fetching logging user in:", error);
      return [];
    }
  };

  const handleChange = (e: { target: { name: any; value: any } }) => {
    const { name, value } = e.target;
    setUserLogin({
      ...userLogin,
      [name]: value,
    });
  };

  const handleSubmit = async (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const info = await loginUser();
    route.push("");
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
        <div className="flex flex-col my-32">
          <h1 className="text-white font-bold text-3xl pb-16">Log in</h1>
          <form
            action="login.php"
            method="post"
            className="flex flex-col"
            // onClick={handleFormClick}
            onSubmit={handleSubmit}
          >
            <label htmlFor="netid" className="text-white font-light py-2">
              NetID<span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              id="netid"
              name="netid"
              placeholder="abc123"
              required
              className="w-80 h-10 px-4 mb-8"
              onChange={handleChange}
            />
            <label htmlFor="password" className="text-white font-light py-2">
              Password<span className="text-red-500">*</span>
            </label>
            <input
              type="password"
              id="password"
              name="password"
              placeholder="••••••••••"
              required
              className="w-80 h-10 px-4 mb-8"
              onChange={handleChange}
            />
            {loginError && <p className="text-red-500 pb-2">{loginError}</p>}
            <input
              type="submit"
              value="LOG IN"
              className="text-white font-bold w-24 h-10 px-4 border border-white my-16"
            />
          </form>
          <h3 className="text-white font-light text-sm text-center">
            Don&apos;t have an account yet?{" "}
            <Link href="signup" className="underline">
              Sign up
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
