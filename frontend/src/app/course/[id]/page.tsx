"use client";

import Announcements from "@/components/coursepage/Announcements";
import Assignments from "@/components/coursepage/Assignments";
import { Course } from "@/lib/types";
import Navbar from "@/components/Navbar";
import AddStudent from "./AddStudent";
import { useState, useEffect } from "react";
import AddButton from "@/components/buttons/AddButton";
import Image from "next/image";

export default function Page({ params }: { params: { id: string } }) {
  const initialCourse: Course = {
    id: "",
    name: "",
    description: "",
    professor: "",
    banner: "",
  };

  const [isAddingStudent, setIsAddingStudent] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  const [data, setData] = useState<Course>(initialCourse);

  useEffect(() => {
    const fetchData = async () => {
      const path = `http://localhost:6789/v1/course/homepage/${params.id}`;
      try {
        const response = await fetch(path);
        if (!response.ok) {
          throw new Error("Failed to fetch data");
        }
        const { course }: { course: Course } = await response.json();
        setData(course);
        console.log(data);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();

    const permissions = localStorage.getItem("permissions");
    if (permissions === "1") {
      setIsTeacher(true);
    }
  }, [params.id]);

  const url = params.id;

  return (
    <div style={{ backgroundColor: "black", minHeight: "100vh" }}>
      <Navbar />
      <div>
        {data && (
          <Image
            src={`http://localhost:6789/v1/course/${data.banner}/banner/read`}
            alt="Course Background"
            width={400}
            height={400}
            objectFit="cover"
          />
        )}
        {isTeacher && (
          <div className="flex justify-end mr-8">
            <AddButton
              text={"Add Student"}
              onClick={() => {
                setIsAddingStudent(true);
              }}
            />
          </div>
        )}
        <div className="relative">
          <div className="py-4 px-8 ml-32 mt-32 h-32 w-96 absolute bg-black bg-opacity-70 flex flex-col justify-center">
            {data && (
              <h1 className="text-white text-3xl font-bold pb-2 block text-opacity-100">
                {data.name}
              </h1>
            )}
          </div>
        </div>
      </div>
      <div className="flex justify-around p-16">
        {data && (
          <div className="flex flex-col w-96">
            <Announcements courseId={url} />
          </div>
        )}
        {data && (
          <div className="flex flex-col">
            <Assignments courseId={url} entries={data.assignments!} />
          </div>
        )}
      </div>
      {isAddingStudent && (
        <AddStudent
          onClose={() => {
            setIsAddingStudent(false);
          }}
          courseId={url}
        />
      )}
    </div>
  );
}
