"use client";

import Announcements from "@/components/coursepage/Announcements";
import Assignments from "@/components/coursepage/Assignments";
import { Course, User } from "@/lib/types";
import AddStudent from "./AddStudent";
import { useEffect, useState } from "react";
import AddButton from "@/components/buttons/AddButton";
import Image from "next/image";
import router from "next/router";

export default function Page({ params }: { params: { id: string } }) {
  const [isAddingStudent, setIsAddingStudent] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  const [data, setData] = useState<Course>({
    assignments: [],
    roster: [],
    id: "",
    name: "",
    description: "",
    professor: "",
    banner: "",
    created_at: "",
    updated_at: "",
    deleted_at: "",
  });
  const [roster, setRoster] = useState<User[]>([]);

  useEffect(() => {
    const permissions = localStorage.getItem("permissions");
    if (permissions === "1") {
      setIsTeacher(true);
    }

    const fetchData = async () => {
      const path = `http://localhost:6789/v1/course/homepage/${params.id}`;
      const response = await fetch(path);
      const { course, roster }: { course: Course; roster: User[] } =
        await response.json();
      return { course, roster };
    };

    fetchData()
      .then(({ course, roster }) => {
        setData(course);
        setRoster(roster);
      })
      .catch(console.error);
  }, [params.id]);

  const url = params.id;

  const handleDeleteCourse = async () => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/delete/${params.id}`,
        {
          method: "DELETE",
        }
      );

      if (response.ok) {
        router.push("/");
      } else {
        console.error("Failed to delete course");
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

  return (
    <>
      <div>
        <div className="relative">
          <div className="bg-slate-900 py-4 px-8 ml-32 mt-32 h-32 w-fit absolute bg-opacity-70 flex flex-col justify-center text-white">
            {data && (
              <>
                <h1 className="text-4xl font-bold pb-2 block text-opacity-100">
                  {data.name}
                </h1>
                <p>COURSE ID: {data.id}</p>
              </>
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

      {/* ANNOUNCEMENT AND ASSIGNMENTS Section*/}
      <div className="grid grid-flow-row grid-cols-3 grid-rows-1 gap-2 mx-20 my-12 mb-20">
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

      {/* ROSTER AND INSTRUCTOR section */}
      <div className="grid grid-flow-row grid-cols-3 grid-rows-1 gap-2 mx-20 my-12 mb-20">
        <div className="flex justify-around">
          {data && (
            <div className="flex flex-col w-full space-y-4">
              <h2 className="text-white font-bold text-2xl">Roster</h2>
              {isTeacher && (
                <AddButton
                  text={"Add Student"}
                  fullWidth={true}
                  onClick={() => {
                    setIsAddingStudent(true);
                  }}
                />
              )}
              <div
                className={
                  "w-full h-full grid grid-flow-row border-slate-200 border-opacity-10 border-2"
                }
              >
                {roster ? (
                  roster.map((user, i) => (
                    <div
                      className="roster-item hover:bg-gray-700 text-white text-md"
                      key={i}
                    >
                      <h4 className="font-bold w-full">{user.full_name}</h4>
                      <p>{user.email}</p>
                    </div>
                  ))
                ) : (
                  <>
                    <div className={"roster-item"}>
                      <p className={"text-hint p-2"}>
                        Students will appear here.
                      </p>
                    </div>
                  </>
                )}
              </div>
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

      {isAddingStudent && (
        <AddStudent
          onClose={() => {
            setIsAddingStudent(false);
          }}
          courseId={url}
        />
      )}
      {isTeacher && (
        <div className="fixed bottom-4 right-4">
          <button
            onClick={handleDeleteCourse}
            className="bg-red-500 hover:bg-red-700 active:bg-red-900 text-white font-bold py-2 px-4 rounded"
          >
            Delete Course
          </button>
        </div>
      )}
    </>
  );
}
