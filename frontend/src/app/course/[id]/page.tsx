"use client";

import Announcements from "@/components/coursepage/Announcements";
import Assignments from "@/components/coursepage/Assignments";
import Discussions from "@/components/coursepage/Discussions";
import {
  Announcement,
  Assignment,
  Discussion,
  User,
  Course,
} from "@/lib/types";
import Navbar from "@/components/Navbar";
import AddStudent from "./AddStudent";
import { useState, useEffect } from "react";
import AddButton from "@/components/buttons/AddButton";

// These data names must match what the API returns.
interface HomepageData {
  course_info: {
    name: string;
    teachers: User[];
    assignments: Assignment[];
    discussions: Discussion[];
    announcements: Announcement[];
    roster: User[];
  };
}

// Dynamic route example found here:
// https://nextjs.org/docs/app/building-your-application/routing/dynamic-routes#example
export default function Page({ params }: { params: { id: string } }) {
  const [isAddingStudent, setIsAddingStudent] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  let [data, setData] = useState<HomepageData>();
  // const path = `http://localhost:6789/v1/course/homepage/${params.id}`;
  // const res: HomepageData = fetch(path).then((response) => {
  //   console.log(response.json());
  // });

  useEffect(() => {
    const fetchData = async () => {
      const path = `http://localhost:6789/v1/course/homepage/${params.id}`;
      try {
        const response = await fetch(path);
        if (!response.ok) {
          throw new Error("Failed to fetch data");
        }
        const fetchedData: HomepageData = await response.json();
        setData(fetchedData);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();

    const token = localStorage.getItem("token");
    const permissions = localStorage.getItem("permissions");
    if (permissions === "1") {
      setIsTeacher(true);
    }
  }, [params.id]);

  const url = params.id;

  return (
    <div style={{ backgroundColor: "black", minHeight: "100vh" }}>
      <Navbar />
      <div
        style={{
          backgroundImage: `url('/backgrounds/course-bg.jpg')`,
          backgroundSize: "cover",
          backgroundPosition: "center",
          width: "100%",
          height: "300px",
          paddingTop: "16px",
        }}
      >
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
                {data.course_info.name}
              </h1>
            )}
          </div>
          {/* <div className="flex justify-end">
            <Discussions />
          </div> */}
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
            <Assignments
              courseId={url}
              entries={data.course_info.assignments}
            />
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
