"use client";

import Image from "next/image";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import CreateCourse from "@/components/homepage/CreateCourse";
import AddButton from "@/components/buttons/AddButton";
import { Course } from "@/lib/types";
import CourseItem from "@/components/homepage/Courses";

export default function Home() {
  const [isCreatingCourse, setIsCreatingCourse] = useState(false);
  const [courses, setCourses] = useState<Course[]>([]);
  const [courseBanners, setCourseBanners] = useState<Map<string, any>>(
    new Map()
  );
  const [navbarActive, setNavbarActive] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  const router = useRouter();

  const currentTerm = "Spring 2024";

  const handleCreateCourse = (courseData: any) => {
    setCourses([...courses, courseData]);
  };

  useEffect(() => {
    const fetchData = async () => {
      const token = localStorage.getItem("token");
      const permissions = localStorage.getItem("permissions");
      if (permissions === "1") {
        setIsTeacher(true);
      }
      if (token) {
        try {
          const fetchedCourses = await fetchCourses(token);
          console.log("FETCHED COURSES: ", fetchedCourses);
          const keys = Object.keys(fetchedCourses.courses);
          const coursesWithBanners = await Promise.all(
            keys.map(async (key) => {
              const course = fetchedCourses.courses[key];
              const banner = await fetchBanner(course.banner);
              console.log("BANNER FOR COURSE ID", course.id, ":", banner);
              setCourseBanners((prevState) =>
                new Map(prevState).set(course.id, banner)
              );
              return course;
            })
          );
          console.log("COURSES WITH BANNERS: ", courseBanners);
          setCourses(coursesWithBanners);
        } catch (error) {
          console.error("Error fetching courses:", error);
        }
      }
    };
    fetchData();
  }, []);

  const fetchCourses = async (tok: string) => {
    try {
      const res: Response = await fetch("http://localhost:6789/v1/home", {
        method: "POST",
        body: JSON.stringify({
          token: tok,
        }),
      });
      if (res.ok) {
        const data = await res.json();
        console.log(data);
        return data;
      } else {
        console.error("Failed to fetch courses:", res.statusText);
        return [];
      }
    } catch (error) {
      console.error("Error fetching courses:", error);
      return [];
    }
  };

  const fetchBanner = async (bannerId: string) => {
    try {
      const res: Response = await fetch(
        `http://localhost:6789/v1/course/${bannerId}/banner/read`
      );
      if (res.ok) {
        const data = await res.json();
        console.log(data);
        return data;
      } else {
        console.error("Failed to fetch courses:", res.statusText);
        return [];
      }
    } catch (error) {
      console.error("Error fetching courses:", error);
      return [];
    }
  };

  const handleIconClick = () => {
    setNavbarActive(!navbarActive);
  };

  return (
    <div style={{ backgroundColor: "black", minHeight: "100vh" }}>
      <nav
        style={{
          backgroundImage: `url('/backgrounds/dashboard-bg.jpeg')`,
          backgroundSize: "cover",
          backgroundPosition: "center",
        }}
      >
        <div className="relative">
          <div className="absolute inset-0 bg-black opacity-70"></div>
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
      <div className="bg-black bg-cover bg-no-repeat">
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
          {courses.map((course, i) => (
            <CourseItem
              key={i}
              data={course}
              banner={courseBanners.get(course.id)}
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
