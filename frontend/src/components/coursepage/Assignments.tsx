"use client";

import React, { useState, useEffect } from "react";
import AddButton from "@/components/buttons/AddButton";
import { Assignment } from "@/lib/types";
import CreateAssignment from "./CreateAssignment";
import AssignmentDisplay from "./AssignmentDisplay";
import formattedDate from "@/lib/helpers/formattedDate";
import TeacherViewAssignment from "./TeacherViewAssignments";

interface props {
  entries: Assignment[];
  courseId: string;
}

const Assignments: React.FC<props> = ({ entries, courseId }: props) => {
  const [selectedAssignment, setSelectedAssignment] = useState({
    id: "",
    title: "",
    due_date: "",
    description: "",
    created_at: "",
    updated_at: "",
    deleted_at: "",
  });

  const [assignments, setAssignments] = useState<Assignment[]>(entries);
  /**
   * assignmentMap maps assignment object values to their
   * name pairs.
   */
  const [assignmentMap, setAssignmentMap] = useState<Map<string, Assignment>>(
    new Map<string, Assignment>()
  );

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

    const fetchAssignments = async (): Promise<Assignment[]> => {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/read/${courseId}`
      );
      const { assignments }: { assignments: Assignment[] } =
        await response.json();
      setAssignments(assignments);
      return assignments;
    };

    fetchAssignments()
      .then((a: Assignment[]) => {
        // Assign map's keys and values based on fetched assignments.
        const newMap: Map<string, Assignment> = new Map<string, Assignment>();
        for (let i = 0; i < a.length; i++) {
          const el = a[i];
          newMap.set(el.title, el);
        }
        setAssignmentMap(newMap);
        console.log(newMap);
      })
      .catch(console.error);
  }, []);

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    // Get the assignment by its UUID.
    const value: Assignment = assignmentMap.get(event.target.value)!;
    setSelectedAssignment(value);
    setSetIsViewingAssignment(true);
  };

  const refreshData = async () => {
    setIsCreatingAssignment(false);
  };

  return (
    <div className="w-full bg-red-50">
      {isTeacher && (
        <AddButton
          fullWidth={true}
          text="Create Assignment"
          onClick={() => {
            setIsCreatingAssignment(true);
          }}
        />
      )}
      <div className="flex">
        <ul className="w-full">
          {assignments &&
            assignments.map((assignment: Assignment, i: number) => (
              <li
                className="flex flex-col max-w-sm p-6 h-46 border shadow bg-gray-900 border-gray-700 hover:bg-gray-700"
                key={i}
              >
                <h5 className="mb-2 text-lg text-white">{assignment.title}</h5>
                <div className="mb-2 rounded-xl  bg-yellow-700 w-fit">
                  <h6 className="text-xs tracking-wide text-white px-2 py-1">
                    {formattedDate(assignment.due_date).toLocaleUpperCase()}
                  </h6>
                </div>
                {/* <p className="font-normal tracking-wide text-gray-400">
                  {truncateString(assignment.description, 50)}
                </p> */}
              </li>
            ))}
        </ul>
      </div>
      {isCreatingAssignment && (
        <CreateAssignment
          onClose={() => {
            refreshData();
          }}
          token={token}
          params={{ id: courseId }}
        />
      )}
      {isViewingAssignment && !isTeacher && (
        <AssignmentDisplay
          onClose={() => {
            setSetIsViewingAssignment(false);
          }}
          assignment={selectedAssignment}
        />
      )}
      {isViewingAssignment && isTeacher && (
        <TeacherViewAssignment
          onClose={() => {
            setSetIsViewingAssignment(false);
          }}
          assignmentid={selectedAssignment}
        />
      )}
    </div>
  );
};

export default Assignments;
