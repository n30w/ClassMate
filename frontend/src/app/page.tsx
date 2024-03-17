"use client";

import Image from "next/image";
import Courses from "@/components/Courses";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import CreateCourse from "@/components/CreateCourse";
import AddButton from "@/components/buttons/AddButton";

interface Course {
  id: string;
  title: string;
  professor: string;
  location: string;
}

const CourseDisplay: React.FC<{ courses: any[] }> = ({ courses }) => {
  const router = useRouter();
  return (
    <>
      {courses.map((course) => (
        <Courses
          key={course.id}
          courseName={course.title}
          professor={course.professor}
          loc={course.location}
          onClick={() => router.push(`/course/${course.id}`)}
        />
      ))}
    </>
  );
};

export default function Home() {
  const [isCreatingCourse, setIsCreatingCourse] = useState(false);
  const [courses, setCourses] = useState<Course[]>([]);

  const handleCreateCourse = (courseData: any) => {
    setCourses([...courses, courseData]);
  };

  const fetchCourses = async () => {
    try {
      const res: Response = await fetch("/v1/course/read", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (res.ok) {
        const courses = await res.json();
        return courses;
      } else {
        console.error("Failed to fetch courses:", res.statusText);
        return [];
      }
    } catch (error) {
      console.error("Error fetching courses:", error);
      return [];
    }
  };

  useEffect(() => {
    const getCourses = async () => {
      const fetchedCourses = await fetchCourses();
      setCourses(fetchedCourses);
    };

    getCourses();
  }, []);

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
          <div className="py-8 px-32">
            <div className="flex items-center gap-4">
              <Image
                src="/backgrounds/NYU-logo.png"
                width="150"
                height="39"
                alt="NYU Logo"
                className="z-10"
              />
              <Image
                src="/backgrounds/darkspace.png"
                width="200"
                height="39"
                alt="Darkspace Logo"
                className="z-10"
              />
            </div>
          </div>
        </div>
      </nav>
      <div className="bg-black bg-cover bg-no-repeat">
        <div className="flex items-center justify-between py-8 px-32">
          <h1 className="font-bold text-4xl text-white">Spring 2024</h1>
          <AddButton
            text={"New Course"}
            onClick={() => {
              setIsCreatingCourse(true);
            }}
          />
        </div>
        <CourseDisplay courses={courses} />
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
