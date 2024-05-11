"use client";

import React, { useEffect, useState } from "react";
import CloseButton from "@/components/buttons/CloseButton";
import { Assignment } from "@/lib/types";

interface props {
  onClose: () => void;
  assignment_id: string;
}

const AssignmentDisplay: React.FC<props> = (props: props) => {
  const [assignment, setAssignment] = useState<Record<string, any> | null>(
    null
  );

  useEffect(() => {
    fetchAssignment();
  }, []);

  const fetchAssignment = async () => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/read/${props.assignment_id}`
      );
      if (response.ok) {
        const data = await response.json();
        console.log("SINGLE ASSIGNMENT INFO: ", data);
        setAssignment(data.assignment);
      } else {
        console.error("Failed to fetch assignments:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching assignments:", error);
    }
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    //Submission logic
    props.onClose();
  };

  const isPastDueDate =
    assignment && new Date(assignment.due_date) < new Date();

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg px-32 py-16 justify-end">
        <CloseButton onClick={props.onClose} />
        <form className="justify-end" onSubmit={handleSubmit}>
          <h1 className="font-bold text-black text-2xl pb-8">
            {assignment ? assignment.post.Title : ""}
          </h1>
          <h1 className="font-bold text-black text-xl pb-8">
            Due Date:{" "}
            {assignment
              ? new Date(assignment.due_date).toLocaleDateString()
              : ""}
          </h1>
          <h1 className="font-bold text-black text-l pb-8">
            Description: {assignment ? assignment.post.Description : ""}
          </h1>
          {!isPastDueDate ? (
            <div className="mb-4">
              <label
                htmlFor="file"
                className="block text-lg font-medium text-gray-700 py-2"
              >
                Drag and drop your file here:
              </label>
              <input
                type="file"
                id="file"
                name="file"
                className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                onChange={() => {}}
              />
            </div>
          ) : (
            <h1 className="font-bold text-black text-l pb-8">
              You can no longer submit this assignment.
            </h1>
          )}
          {!isPastDueDate ? (
            <button
              type="submit"
              className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Submit
            </button>
          ) : (
            <button
              type="submit"
              className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Close
            </button>
          )}
        </form>
      </div>
    </div>
  );
};

export default AssignmentDisplay;
