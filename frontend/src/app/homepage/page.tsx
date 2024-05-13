"use client";

import Image from "next/image";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import CreateCourse from "@/components/homepage/CreateCourse";
import AddButton from "@/components/buttons/AddButton";
import { Course } from "@/lib/types";
import CourseItem from "@/components/homepage/Courses";

export default function Home() {
  const initCourses: Course[] = [];
  const [isCreatingCourse, setIsCreatingCourse] = useState(false);
  const [courseArray, setCourseArray] = useState<Course[]>(initCourses);
  const [navbarActive, setNavbarActive] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  const router = useRouter();

  const currentTerm = "Spring 2024";

  const handleCreateCourse = (courseData: any) => {
    setCourseArray([...courseArray, courseData]);
  };

  useEffect(() => {
    const token = localStorage.getItem("token")!;
    const permissions = localStorage.getItem("permissions");

    console.log(token);
    if (permissions === "1") {
      setIsTeacher(true);
    }
    const fetchCourses = async (tok: string) => {
      const route = "http://localhost:6789/v1/home";
      const res: Response = await fetch(route, {
        method: "POST",
        body: JSON.stringify({
          token: tok,
        }),
      });

      const { courses }: { courses: Course[] } = await res.json();
      setCourseArray(courses);
      console.log(courseArray);
    };

    fetchCourses(token).catch(console.error);
  }, []);

  const handleIconClick = () => {
    setNavbarActive(!navbarActive);
  };

  return (
    <div style={{ minHeight: "100vh" }} className="bg-slate-900">
      <nav
        style={{
          backgroundImage: `url('/backgrounds/dashboard-bg.jpeg')`,
          backgroundSize: "cover",
          backgroundPosition: "center",
        }}
      >
        <div className="relative">
          <div className="absolute inset-0 opacity-70"></div>
          <div className="py-8 px-32 flex justify-between items-center relative z-10">
            <div className="flex items-center gap-4">
              <Image
                src="/backgrounds/NYU-logo.png"
                width="150"
                height="39"
                alt="NYU Logo"
              />
              <Image
                src="/backgrounds/darkspace.png"
                width="200"
                height="39"
                alt="Darkspace Logo"
              />
            </div>
            <div className="flex items-center gap-4">
              {navbarActive && (
                <div className="bg-white h-16 w-40 p-2 flex justify-center items-center rounded-md gap-4">
                  <p
                    className="text-black border-black hover:text-gray-500"
                    onClick={() => router.push("/login")}
                  >
                    Login
                  </p>
                  <p
                    className="text-black border-black hover:text-gray-500"
                    onClick={() => router.push("/signup")}
                  >
                    Sign up
                  </p>
                </div>
              )}
              <Image
                src="/backgrounds/profile-icon.png"
                width="50"
                height="39"
                alt="Profile Icon"
                onClick={handleIconClick}
              />
            </div>
          </div>
        </div>
      </nav>
      <div className="bg-cover bg-no-repeat">
        (
        <div className="flex items-center justify-between py-8 px-32">
          <h1 className="font-bold text-4xl text-white">{currentTerm}</h1>
          {isTeacher && (
            <AddButton
              text={"New Course"}
              onClick={() => {
                setIsCreatingCourse(true);
              }}
            />
          )}
        </div>
        )
        <div className="grid grid-cols-3 gap-4 mr-16">
          {courseArray.map((course, i) => (
            <CourseItem
              key={i}
              data={course}
              onClick={() => {
                router.push(`/course/${course.id}`);
              }}
            />
          ))}
        </div>
        {isCreatingCourse && (
          <CreateCourse
            onClose={() => {
              setIsCreatingCourse(false);
            }}
            onCourseCreate={handleCreateCourse}
          />
        )}
      </div>
    </div>
  );
}
