"use client";

import React, { useEffect, useState } from "react";
import AddButton from "@/components/buttons/AddButton";
import { Assignment } from "@/lib/types";
import CreateAssignment from "./CreateAssignment";
import AssignmentDisplay from "./AssignmentDisplay";
import InfoBadge from "@/components/badge/InfoBadge";
import truncateString from "@/lib/helpers/truncateString";
import Router from "next/client";
import { useRouter } from "next/navigation";
import Link from "next/link";
import formattedDate from "@/lib/helpers/formattedDate";

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
    new Map<string, Assignment>(),
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
      const init: RequestInit = {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      };
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/read/${courseId}`,
      );
      const { assignments }: { assignments: Assignment[] } =
        await response.json();
      setAssignments(assignments);
      return assignments;
    };

    fetchAssignments()
      .then((a: Assignment[]) => {
        if (a) {
          // Assign map's keys and values based on fetched assignments.
          const newMap: Map<string, Assignment> = new Map<string, Assignment>();
          for (let i = 0; i < a.length; i++) {
            const el = a[i];
            newMap.set(el.title, el);
          }
          setAssignmentMap(newMap);
          console.log(newMap);
        }
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

  const router = useRouter();

  return (
    <div className="w-full h-full">
      {isTeacher && (
        <AddButton
          fullWidth={true}
          text="Create Assignment"
          onClick={() => {
            setIsCreatingAssignment(true);
          }}
        />
      )}
      <div className="w-full grid grid-cols-1 grid-rows-3 border-2 border-slate-300 border-opacity-10">
        {assignments ? (
          assignments.map((assignment: Assignment, i: number) => (
            <Link href={`/course/${courseId}/assignments/${assignment.id}`}>
              <div className="assignment-item hover:bg-gray-700" key={i}>
                <h5 className="mb-2 text-lg text-white">{assignment.title}</h5>
                <InfoBadge
                  text={formattedDate(assignment.due_date).toLocaleUpperCase()}
                />
                {assignment.description && (
                  <p className="font-normal tracking-wide text-gray-400">
                    {truncateString(assignment.description, 20)}
                  </p>
                )}
              </div>
            </Link>
          ))
        ) : (
          <>
            <div className={"assignment-item"}>
              <p className={"text-hint"}> New assignments will appear here.</p>
            </div>
            <div className={"assignment-item"}></div>
          </>
        )}
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
      {/*{isViewingAssignment && isTeacher && (*/}
      {/*  <TeacherViewAssignment*/}
      {/*    onClose={() => {*/}
      {/*      setSetIsViewingAssignment(false);*/}
      {/*    }}*/}
      {/*    assignmentid={selectedAssignment}*/}
      {/*  />*/}
      {/*)}*/}
    </div>
  );
};

export default Assignments;
