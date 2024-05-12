"use client";
import { Assignment } from "@/lib/types";
import { useEffect, useState } from "react";

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
            // token: token,
          }),
        },
      );
      const { assignment }: { assignment: Assignment } = await response.json();
      return assignment;
    };

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
      <h1>{viewAssignment.title}</h1>
    </>
  );
}
