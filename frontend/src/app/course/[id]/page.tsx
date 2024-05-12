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
  const [isAddingStudent, setIsAddingStudent] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  const [data, setData] = useState<Course>({
    id: "",
    name: "",
    description: "",
    professor: "",
    banner: "",
    created_at: "",
    updated_at: "",
    deleted_at: "",
  });

  useEffect(() => {
    const permissions = localStorage.getItem("permissions");
    if (permissions === "1") {
      setIsTeacher(true);
    }

    const fetchData = async () => {
      const path = `http://localhost:6789/v1/course/homepage/${params.id}`;
      const response = await fetch(path);
      const { course }: { course: Course } = await response.json();
      return course;
    };

    fetchData()
      .then((c: Course) => {
        setData(c);
      })
      .catch(console.error);
  }, [params.id]);

  const url = params.id;

  return (
    <div style={{ minHeight: "100vh" }} className="bg-slate-900">
      <Navbar />
      <div>
        <div className="relative">
          <div className="py-4 px-8 ml-32 mt-32 h-32 w-96 absolute  bg-opacity-70 flex flex-col justify-center">
            {data && (
              <h1 className="text-white text-4xl font-bold pb-2 block text-opacity-100">
                {data.name}
              </h1>
            )}
          </div>
        </div>

        {data.banner && (
          <div className={"w-full h-[400px] flex"}>
            <Image
              src={`http://localhost:6789/v1/course/${data.banner}/banner/read`}
              alt="Course Background"
              className="w-full h-full"
              style={{ objectFit: "cover" }}
              width={400}
              height={400}
            />
          </div>
        )}
      </div>

      <div className="grid grid-flow-row grid-cols-3 grid-rows-1 gap-2 mx-20 my-12">
        <div className="flex justify-around col-span-2">
          {data && (
            <div className="flex flex-col w-full space-y-4">
              <h2 className="text-white font-bold text-2xl">Announcements</h2>
              <Announcements courseId={url} />
            </div>
          )}
        </div>
        <div className="flex justify-around">
          {data && (
            <div className="flex flex-col space-y-4">
              <h2 className="text-white font-bold text-2xl">Assignments</h2>
              <Assignments entries={data.assignments!} courseId={url} />
            </div>
          )}
        </div>
      </div>

      <div className="grid grid-flow-row grid-cols-3 grid-rows-1 gap-2 mx-20 my-12">
        <div className="flex justify-around">
          {data && (
            <div className="flex flex-col w-full">
              <h2 className="text-white font-bold text-2xl">Roster</h2>
              <div></div>
            </div>
          )}
        </div>
        <div className="flex justify-around">
          {data && (
            <div className="flex flex-col space-y-4">
              <h2 className="text-white font-bold text-2xl">Instructors</h2>
              <div></div>
            </div>
          )}
        </div>
      </div>

      {isTeacher && (
        <div className="flex justify-end mr-8 mt-4 bg-red-100">
          <AddButton
            text={"Student"}
            onClick={() => {
              setIsAddingStudent(true);
            }}
          />
        </div>
      )}

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
