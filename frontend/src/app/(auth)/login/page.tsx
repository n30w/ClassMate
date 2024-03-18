"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";
import Link from "next/link";
import { userAgentFromString } from "next/server";

export default function Page() {
  // const [isBlurred, setIsBlurred] = useState(false);

  // const handleFormClick = (): void => {
  //   setIsBlurred(true);
  // };

  const [userLogin, setUserLogin] = useState({
    email: "",
    password: "",
  });
  const [loginError, setLoginError] = useState<string>("");
  const route = useRouter();

  const fetchUserInfo = async () => {
    try {
      const res: Response = await fetch("/v1/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userLogin),
      });
      if (res.ok) {
        const userInfo = await res.json();
        return userInfo;
      } else {
        console.error("Failed to fetch user info:", res.statusText);
        return [];
      }
    } catch (error) {
      console.error("Error fetching user info:", error);
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
    const info = await fetchUserInfo();
    for (let i = 0; i < info.length; i++) {
      if (info[i].email === userLogin.email) {
        if (info[i].password === userLogin.password) {
          route.push("");
        } else {
          setLoginError("Wrong password entered");
        }
      }
    }
    setLoginError("Wrong email entered");
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
              onChange={handleChange}
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
