"use client";

import React, { useState, useEffect } from "react";
import AddButton from "@/components/buttons/AddButton";
import { Assignment } from "@/lib/types";
import CreateAssignment from "./CreateAssignment";
import AssignmentDisplay from "./AssignmentDisplay";

interface props {
  entries: Assignment[];
  courseId: string;
}

const Assignments: React.FC<props> = (props: props) => {
  const [selectedAssignment, setSelectedAssignment] = useState("");
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  const [assignments, setAssignments] = useState<Assignment[]>([]);
  const [isTeacher, setIsTeacher] = useState(false);
  const [isCreatingAssignment, setIsCreatingAssignment] = useState(false);
  const [token, setIsToken] = useState("");
  const [isViewingAssignment, setSetIsViewingAssignment] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setIsToken(token);
    }
    const permissions = localStorage.getItem("permissions");
    if (permissions === "1") {
      setIsTeacher(true);
    }
    fetchAssignments();
  }, [assignments]);

  const fetchAssignments = async () => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/read/${props.courseId}`
      );
      if (response.ok) {
        const data = await response.json();
        setAssignments(data.assignment);
      } else {
        console.error("Failed to fetch assignments:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching assignments:", error);
    }
  };

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedAssignment(event.target.value);
  };

  const postSubmission = async (submissionData: any) => {
    try {
      const formData = new FormData();
      submissionData.forEach((file: File) => {
        formData.append("files", file);
      });
      formData.append("submissiontime", new Date().toISOString());
      formData.append("assignmentid", submissionData.assignmentid);
      formData.append("userid", submissionData.userid);

      const res: Response = await fetch(
        "http://localhost:6789/v1/course/assignment/submission/create",
        {
          method: "POST",
          body: formData,
        }
      );

      if (res.ok) {
      } else {
        console.error("Failed to create submission:", res.statusText);
      }
    } catch (error) {
      console.error("Error creating submission:", error);
    }
  };

  const refreshData = async () => {
    setIsCreatingAssignment(false);
    await fetchAssignments();
  };

  return (
    <div className="w-full">
      <div className="flex">
        {isTeacher && (
          <AddButton
            onClick={() => {
              setIsCreatingAssignment(true);
            }}
          />
        )}
        <select
          className="ml-16 p-2"
          value={selectedAssignment}
          onChange={handleSelectChange}
        >
          <option value="">Choose an assignment</option>
          {assignments &&
            assignments.map((assignment: any, index: number) => (
              <option key={index} value={index}>
                {assignment.name}
              </option>
            ))}
        </select>
      </div>
      {selectedAssignment !== "" &&
        assignments[parseInt(selectedAssignment)] && (
          <>
            <h2 className="text-white text-2xl font-bold mb-2 pt-4">
              {assignments[parseInt(selectedAssignment)].title}
            </h2>
            <p className="text-white text-sm my-4">
              Due Date: {assignments[parseInt(selectedAssignment)].duedate}
            </p>
            <p className="text-white text-lg font-light pb-8">
              {assignments[parseInt(selectedAssignment)].description}
            </p>
          </>
        )}
      {isCreatingAssignment && (
        <CreateAssignment
          onClose={() => {
            refreshData();
          }}
          token={token}
          params={{ id: props.courseId }}
        />
      )}
      {isViewingAssignment && <AssignmentDisplay onClose={() => {}} />}
    </div>
  );
};

export default Assignments;
