"use client";

import React, { useState } from "react";
import addStudentToClass from "./actions";
import CloseButton from "@/components/buttons/CloseButton";

interface AddStudentProps {
  onClose: () => void;
  courseId: string;
}

const AddStudent: React.FC<AddStudentProps> = (props) => {
  const perms = localStorage.getItem("permissions");
  const [netId, setNetId] = useState("");

  const postNewStudent = async (studentid: string) => {
    try {
      const res: Response = await fetch(
        "http://localhost:6789/v1/course/addstudent",
        {
          method: "POST",
          body: JSON.stringify({
            netid: studentid,
            courseid: props.courseId,
          }),
        }
      );
      if (res.ok) {
        const response = await res.json();
        console.log(response);
      } else {
        console.error("Failed to add student to the course:", res.statusText);
      }
    } catch (error) {
      console.error("Error adding student to the course:", error);
    }
  };

  const handleAddStudent = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    console.log("submitted student netid");
    postNewStudent(netId);
    setNetId("");
    props.onClose();
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg px-32 py-16 justify-end">
        <CloseButton onClick={props.onClose} />
        <form className="justify-end" onSubmit={handleAddStudent}>
          <h1 className="font-bold text-black text-2xl pb-8">Add Student</h1>
          <div className="mb-2">
            <label
              htmlFor="banner"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Student NetId:
            </label>
            <input
              type="text"
              name="netid"
              value={netId}
              onChange={(e) => setNetId(e.target.value)}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Add
          </button>
        </form>
      </div>
    </div>
  );
};

export default AddStudent;
