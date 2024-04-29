"use client";

import Image from "next/image";
import Link from "next/link";
import React, { useState } from "react";
import validatePassword from "@/lib/helpers/passwordValidator";
import FormInput from "./FormInput";
import { useRouter } from "next/navigation";

const SignUpForm = () => {
  const [userData, setUserData] = useState({
    email: "",
    fullname: "",
    password: "",
    netid: "",
    membership: 0,
  });
  const [passwordError, setPasswordError] = useState("");
  // const [reenteredPassword, setReenteredPassword] = useState("");
  // const [reenteredPasswordError, setReenteredPasswordError] = useState("");

  const router = useRouter();

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const isValidPassword = validatePassword(userData.password);
    // const doPasswordsMatch = userData.password === reenteredPassword;

    // console.log(userData.password);
    // console.log(reenteredPassword);

    if (!isValidPassword) {
      setPasswordError(
        "Password must have at least one letter, one number, one special character, and at least 8 characters long."
      );
      return;
    }

    // if (!doPasswordsMatch) {
    //   setReenteredPasswordError("Passwords do not match.");
    //   return;
    // }
    // Call the function to post new user data
    postNewUser(userData);
  };

  const handleChange = (e: { target: { name: any; value: any } }) => {
    const { name, value } = e.target;
    setUserData({
      ...userData,
      [name]: value,
    });
    setPasswordError("");
    // setReenteredPassword(value);
    // setReenteredPasswordError("");
  };

  const postNewUser = async (userData: any) => {
    try {
      const res: Response = await fetch(
        "http://localhost:6789/v1/user/create",
        {
          method: "POST",
          body: JSON.stringify({
            email: userData.email,
            fullname: userData.fullname,
            password: userData.password,
            netid: userData.netid,
            membership: userData.membership,
          }),
        }
      );
      if (res.status !== 400) {
        router.push("/login");
      } else {
        setPasswordError(res.statusText);
        console.error("Failed to create user:", res.statusText);
      }
    } catch (error) {
      console.error("Error creating user:", error);
    }
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
            <FormInput
              label="NetId"
              type="text"
              name="netid"
              placeholder="abc123"
              value={userData.netid}
              onChange={handleChange}
            />
            <FormInput
              label="Fullname"
              type="text"
              name="fullname"
              placeholder="John Smith"
              value={userData.fullname}
              onChange={handleChange}
            />
            <FormInput
              label="Email"
              type="text"
              name="email"
              placeholder="abc123@nyu.edu"
              value={userData.email}
              onChange={handleChange}
            />
            <div className="flex justify-between items-center">
              <span className="text-white mr-4">Membership:</span>
              <div>
                <button
                  type="button"
                  className="bg-white px-4 py-2 rounded focus:outline-none mr-16 mb-8"
                  onClick={() => setUserData({ ...userData, membership: 1 })}
                >
                  Teacher
                </button>
                <button
                  type="button"
                  className="bg-white px-4 py-2 rounded focus:outline-none"
                  onClick={() => setUserData({ ...userData, membership: 0 })}
                >
                  Student
                </button>
              </div>
            </div>
            <FormInput
              label="Password"
              type="password"
              name="password"
              placeholder="Enter password"
              value={userData.password}
              onChange={handleChange}
              errorMessage={passwordError}
            />
            {/* <FormInput
              label="Re-enter Password"
              type="password"
              name="reenteredPassword"
              placeholder="Re-enter password"
              value={reenteredPassword}
              onChange={(e: {
                target: { value: React.SetStateAction<string> };
              }) => setReenteredPassword(e.target.value)}
              errorMessage={reenteredPasswordError}
            />*/}
            <input
              type="submit"
              className="text-white font-bold w-40 h-10 px-4 border border-white my-16 hover:bg-gray-400 active:bg-white active:text-black"
              data-testid="submitButton"
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
};

export default SignUpForm;
