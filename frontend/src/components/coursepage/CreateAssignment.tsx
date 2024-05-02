"use client";

import React, { useState } from "react";
import CloseButton from "@/components/buttons/CloseButton";

interface props {
  onClose: () => void;
  onCourseCreate: (assignmentData: any) => void;
}

const CreateAssignment: React.FC<props> = (props: props) => {
  const [assignmentData, setAssignmentData] = useState({
    id: "",
    title: "",
    dueDate: "",
    description: "",
  });

  const postNewAssignment = async (assignmentData: any) => {
    try {
      const res: Response = await fetch(
        "http://localhost:6789/v1/course/assignment/create",
        {
          method: "POST",
          body: JSON.stringify(assignmentData),
        }
      );
      if (res.ok) {
      } else {
        console.error("Failed to create assignment:", res.statusText);
      }
    } catch (error) {
      console.error("Error creating assignment:", error);
    }
  };

  const handleChange = (e: { target: { name: any; value: any } }) => {
    const { name, value } = e.target;
    setAssignmentData({
      ...assignmentData,
      [name]: value,
    });
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const idNum = Date.now().toString();
    setAssignmentData({
      ...assignmentData,
      id: idNum,
    });
    props.onCourseCreate({ ...assignmentData, id: idNum });
    postNewAssignment(assignmentData);
    props.onClose();
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg px-32 py-16 justify-end">
        <CloseButton onClick={props.onClose} />
        <form className="justify-end" onSubmit={handleSubmit}>
          <h1 className="font-bold text-black text-2xl pb-8">
            Create Assignment
          </h1>
          <div className="mb-2">
            <label
              htmlFor="title"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Assignment Title:
            </label>
            <input
              type="text"
              id="title"
              name="title"
              value={assignmentData.title}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            />
          </div>
          <div className="mb-2">
            <label
              htmlFor="duedate"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Due Date:
            </label>
            <input
              type="long text"
              id="duedate"
              name="duedate"
              value={assignmentData.dueDate}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            />
          </div>
          <div className="mb-2">
            <label
              htmlFor="description"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Description:
            </label>
            <input
              type="text"
              id="description"
              name="description"
              value={assignmentData.description}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            />
          </div>
          <button
            type="submit"
            className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Create
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreateAssignment;
