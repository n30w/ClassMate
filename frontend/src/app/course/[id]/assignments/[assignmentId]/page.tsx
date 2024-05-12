"use client";
import { Assignment, Submission } from "@/lib/types";
import { useEffect, useState } from "react";
import InfoBadge from "@/components/badge/InfoBadge";
import formattedDate from "@/lib/helpers/formattedDate";

export default function Page({ params }: { params: { assignmentId: string } }) {
  const [viewAssignment, setViewAssignment] = useState<Assignment>({
    created_at: "",
    deleted_at: "",
    description: "",
    due_date: "",
    id: "",
    title: "",
    updated_at: "",
  });
  const [token, setIsToken] = useState("");
  useEffect(() => {
    const t = localStorage.getItem("token");
    if (t) {
      setIsToken(t);
    }

    const fetchAssignment = async (): Promise<Assignment> => {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/read/${params.assignmentId}`,
        {
          method: "POST",
          body: JSON.stringify({
            assignmentId: params.assignmentId,
            token: token,
          }),
        },
      );
      const { assignment }: { assignment: Assignment } = await response.json();
      return assignment;
    };

    const fetchSubmissions = async (): Promise<Submission> => {};

    fetchAssignment()
      .then((value: Assignment) => {
        setViewAssignment(value);
      })
      .catch(console.error);
  });

  // PDF viewer?

  // Query assignment data from the database.

  return (
    <>
      <div className={"container mx-20 my-12 mb-20"}>
        <div className="flex-col gap-2 text-white pb-8">
          <h1 className={"text-white text-5xl font-bold"}>
            {viewAssignment.title}
          </h1>
        </div>
        <div className="grid grid-flow-row grid-cols-2 grid-rows-1 gap-2 text-white">
          <div className={"flex flex-col"}>
            <div className={"misc-item"}>
              <h2>Description</h2>
              {viewAssignment.description ? (
                <p className={"text-lg"}>{viewAssignment.description}</p>
              ) : (
                <p className={"text-hint"}>No description provided.</p>
              )}
            </div>
            <div className={"misc-item"}>
              <h2>UUID</h2>
              <InfoBadge text={viewAssignment.id} colorClass={"bg-green-700"} />
            </div>
            <div className={"misc-item"}>
              <h2>Due Date</h2>
              <InfoBadge
                text={formattedDate(
                  viewAssignment.due_date,
                ).toLocaleUpperCase()}
              />
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
