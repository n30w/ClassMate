"use client";
import { Assignment, Submission, User } from "@/lib/types";
import { useEffect, useState } from "react";
import InfoBadge from "@/components/badge/InfoBadge";
import formattedDate from "@/lib/helpers/formattedDate";
import router from "next/router";

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
  const [studentSubmission, setStudentSubmission] = useState<File>();
  const [roster, setRoster] = useState<User[]>([]);
  const [courseId, setCourseId] = useState<string>("");
  const [isTeacher, setIsTeacher] = useState(false);
  const [submittedFile, setSubmittedFile] = useState(false);
  const [grade, setGrade] = useState(0);
  const [feedback, setFeedback] = useState("");
  const [submissionId, setSubmissionId] = useState("");
  const [uploadedExcel, setUploadedExcel] = useState<File>();
  const [studentsSubmission, setStudentsSubmission] = useState<Submission>();

  useEffect(() => {
    const urlPath = window.location.pathname;
    const pathParts = urlPath.split("/");
    const courseIdIndex = pathParts.indexOf("course") + 1;
    const courseId = pathParts[courseIdIndex];
    const permissions = localStorage.getItem("permissions");

    if (permissions === "1") {
      setIsTeacher(true);
    }
    setCourseId(courseId);
  }, []);

  useEffect(() => {
    const t = localStorage.getItem("token");
    if (t) {
      setIsToken(t);
    }

    const fetchAssignment = async (): Promise<Assignment> => {
      const response = await fetch(
        `http://localhost:6789/v1/course/${courseId}/assignment/read`,
        {
          method: "POST",
          body: JSON.stringify({
            assignmentId: params.assignmentId,
            token: token,
          }),
        }
      );
      const { assignment }: { assignment: Assignment } = await response.json();
      return assignment;
    };

    const fetchData = async () => {
      const path = `http://localhost:6789/v1/course/${courseId}/homepage`;
      const response = await fetch(path);
      const { roster }: { roster: User[] } = await response.json();
      return { roster };
    };

    const studentReadSubmission = async () => {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/${params.assignmentId}/submission/read`,
        {
          method: "POST",
          body: JSON.stringify({
            token: token,
          }),
        }
      );
      const { submission }: { submission: Submission } = await response.json();
      return submission;
    };

    fetchData()
      .then(({ roster }) => {
        setRoster(roster);
      })
      .catch(console.error);

    fetchAssignment()
      .then((value: Assignment) => {
        setViewAssignment(value);
      })
      .catch(console.error);

    studentReadSubmission()
      .then((value: Submission) => {
        setStudentsSubmission(value);
        console.log("STUDENT'S SUBMISSION: ", studentsSubmission);
      })
      .catch(console.error);
  }, [courseId]);

  const downloadSubmission = (selectedStudent: string) => {
    const fetchSubmission = async () => {
      const response: any = await fetch(
        `http://localhost:6789/v1/course/assignment/${params.assignmentId}/submission/${selectedStudent}/read`
      );

      const { submission } = await response.json();
      setSubmissionId(submission.id);
      console.log("RESPONSE: ", submission);

      submission.Media.forEach(async (mediaId: string) => {
        try {
          const res: any = await fetch(
            `http://localhost:6789/v1/course/${courseId}/download/${mediaId}`
          );
          const contentDisposition = res.headers.get("Content-Disposition");
          console.log("CONTENT DIS: ", contentDisposition);
          let filename = "file";
          if (contentDisposition) {
            const match = contentDisposition.match(/filename="(.+)"/);
            if (match) {
              filename = match[1];
            }
          }
          const contentType =
            res.headers.get("Content-Type") || "application/octet-stream";
          console.log("CONTENT TYPE: ", contentType);
          const blob = await res.blob();
          const link = document.createElement("a");

          link.href = window.URL.createObjectURL(blob);

          const fileExtension = filename.split(".").pop();

          link.download = `${filename}.pdf`;

          document.body.appendChild(link);
          link.click();

          document.body.removeChild(link);
        } catch (error) {
          console.error("Error downloading media with ID", mediaId, ":", error);
        }
      });
    };

    fetchSubmission()
      .then((value) => {})
      .catch(console.error);
  };

  const handleFileUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files && event.target.files[0];
    if (file) {
      setStudentSubmission(file);
    }
  };

  const handleUploadButtonClick = async () => {
    const subID = await makeSubmission();
    await submitMediaFile(subID);
  };

  const makeSubmission = async () => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/${params.assignmentId}/submission/create`,
        {
          method: "POST",
          body: JSON.stringify({
            token: token,
          }),
        }
      );

      if (response.ok) {
        const { submission } = await response.json();
        return submission.id;
      } else {
        console.error("Failed to upload file");
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

  const submitMediaFile = async (submissionId: string) => {
    if (!studentSubmission) return;
    const formData = new FormData();
    formData.append("files", studentSubmission);
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/submission/${submissionId}/upload`,
        {
          method: "POST",
          body: formData,
        }
      );
      if (response.ok) {
        console.log("Student submission uploaded successfully");
      } else {
        console.error("Failed to upload file");
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

  const handleGradeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newGrade = Number(event.target.value);
    setGrade(newGrade);
  };

  const handleFeedbackChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newFeedback = event.target.value;
    setFeedback(newFeedback);
  };

  const updateSubmission = async () => {
    const response = await fetch(
      `http://localhost:6789/v1/course/assignment/submission/${submissionId}/update`,
      {
        method: "POST",
        body: JSON.stringify({
          grade: grade,
          feedback: feedback,
        }),
      }
    );
  };

  const uploadExcel = async (uploadedExcel: any) => {
    const formData = new FormData();
    formData.append("files", uploadedExcel);
    const response = await fetch(
      `http://localhost:6789/v1/course/${courseId}/assignment/${params.assignmentId}/offline`,
      {
        method: "POST",
        body: formData,
      }
    );
  };

  const downloadExcel = async () => {
    try {
      const res: any = await fetch(
        `http://localhost:6789/v1/course/${courseId}/assignment/${params.assignmentId}/offline`
      );
      const contentDisposition = res.headers.get("Content-Disposition");
      console.log("CONTENT DIS: ", contentDisposition);
      let filename = "file";
      if (contentDisposition) {
        const match = contentDisposition.match(/filename="(.+)"/);
        if (match) {
          filename = match[1];
        }
      }
      const contentType =
        res.headers.get("Content-Type") || "application/octet-stream";
      console.log("CONTENT TYPE: ", contentType);
      const blob = await res.blob();
      const link = document.createElement("a");

      link.href = window.URL.createObjectURL(blob);

      const fileExtension = filename.split(".").pop();

      link.download = `${filename}.xlsx`;

      document.body.appendChild(link);
      link.click();

      document.body.removeChild(link);
    } catch (error) {
      console.log("Error downloading excel file: ", error);
    }
  };

  const handleExcelFileUpload = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files && event.target.files[0];
    if (file) {
      if (
        file.type ===
        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
      ) {
        setUploadedExcel(file);
        uploadExcel(uploadedExcel);
      }
    }
  };

  const handleDeleteAssignment = async () => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/${params.assignmentId}/delete`,
        {
          method: "DELETE",
        }
      );

      if (response.ok) {
        router.push(`/course/${courseId}`);
        console.log("Assignment deleted successfully");
      } else {
        console.error("Failed to delete assignment");
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

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
              <h2>Due Date</h2>
              <InfoBadge
                text={formattedDate(
                  viewAssignment.due_date
                ).toLocaleUpperCase()}
              />
            </div>
            {isTeacher && (
              <div>
                <div className={"misc-item h-32"}>
                  <h2>Grade</h2>
                  <input
                    type="number"
                    id="grade"
                    multiple
                    onChange={handleGradeChange}
                    required
                    className="text-black mb-4 mr-8 rounded-lg pl-2"
                  />
                </div>
                <div className={"misc-item h-32"}>
                  <h2>Feedback</h2>
                  <input
                    type="longtext"
                    id="feedback"
                    multiple
                    onChange={handleFeedbackChange}
                    required
                    className="text-black mb-4 rounded-lg pl-2"
                  />
                </div>
                <button
                  onClick={() => {
                    updateSubmission();
                  }}
                  className="bg-white text-black font-bold py-2 px-4 rounded mb-16 hover:bg-gray-300 active:bg-gray-500 mt-8"
                >
                  Submit grade and feedback
                </button>
                <button
                  onClick={handleDeleteAssignment}
                  className="bg-red-500 text-white font-bold py-2 px-4 rounded mt-4 hover:bg-red-700"
                >
                  Delete Assignment
                </button>
              </div>
            )}
            {!isTeacher && (
              <div>
                <div className={"misc-item h-32"}>
                  <h2>Grade</h2>
                  <InfoBadge
                    text={formattedDate(
                      viewAssignment.due_date
                    ).toLocaleUpperCase()}
                  />
                </div>
                <div className={"misc-item h-32"}>
                  <h2>Feedback</h2>
                  <InfoBadge
                    text={formattedDate(
                      viewAssignment.due_date
                    ).toLocaleUpperCase()}
                  />
                </div>
              </div>
            )}
          </div>
          {roster && !isTeacher && !submittedFile && (
            <>
              <div className="flex flex-col">
                <label
                  htmlFor="fileUpload"
                  className="text-white text-2xl mb-8"
                >
                  Upload your submission:
                </label>
                <input
                  type="file"
                  id="fileUpload"
                  multiple
                  onChange={handleFileUpload}
                  required
                />
                <button
                  onClick={() => {
                    handleUploadButtonClick();
                    setSubmittedFile(true);
                  }}
                  className="bg-white text-black font-bold py-2 px-4 rounded"
                >
                  Upload File
                </button>
              </div>
            </>
          )}
          {roster && isTeacher && (
            <>
              {roster.map((user, i) => (
                <div
                  className="roster-item hover:bg-gray-700 text-white text-md h-fit"
                  key={i}
                  onClick={() => downloadSubmission(user.id)}
                >
                  <h4 className="font-bold w-full">{user.full_name}</h4>
                  <p>{user.id}</p>
                </div>
              ))}
              {!roster && isTeacher && (
                <div className={"roster-item"}>
                  <p className={"text-hint p-2"}>Students will appear here.</p>
                </div>
              )}
            </>
          )}
          {submittedFile && (
            <div className={"roster-item"}>
              <p className={"text-hint p-2"}>
                Submitted assignment successfully!
              </p>
            </div>
          )}
          {isTeacher && (
            <div>
              <label
                htmlFor="fileUpload"
                className="text-white text-xl mb-16 mr-4"
              >
                Upload Excel File:
              </label>
              <input
                type="file"
                id="fileUpload"
                accept=".xlsx"
                onChange={handleExcelFileUpload}
                required
                className="mb-16"
              />
              <button
                onClick={() => {
                  downloadExcel();
                }}
                className="bg-white text-black font-bold py-2 px-4 rounded hover:bg-gray-300 active:bg-gray-500"
              >
                Download Excel File
              </button>
            </div>
          )}
        </div>
      </div>
    </>
  );
}
