"use client";

import React, { useEffect, useState } from "react";
import { Assignment, Submission } from "@/lib/types";
import CloseButton from "../buttons/CloseButton";

interface props {
  assignmentid: Assignment;
  onClose: () => void;
}

const TeacherViewAssignment: React.FC<props> = ({
  assignmentid,
  onClose,
}: props) => {
  const [assignment, setAssignment] = useState<Assignment>();
  const [submissions, setSubmissions] = useState<Submission[]>([]);
  const [selectedStudent, setSelectedStudent] = useState("");
  const [grade, setGrade] = useState("");
  const [feedback, setFeedback] = useState("");
  const [excelFile, setExcelFile] = useState(null);

  useEffect(() => {
    fetchAssignmentDetails();
    fetchSubmissions();
  }, [assignmentid]);

  const fetchAssignmentDetails = async () => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/read/${assignmentid}`
      );
      if (response.ok) {
        const data = await response.json();
        setAssignment(data.assignment);
        console.log("ASSIGNMENT DETAILS: ", assignment);
      } else {
        console.error(
          "Failed to fetch assignment details:",
          response.statusText
        );
      }
    } catch (error) {
      console.error("Error fetching assignment details:", error);
    }
  };

  const fetchSubmissions = async () => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/submission/${assignmentid}/read`
      );
      if (response.ok) {
        const data = await response.json();
        setSubmissions(data);
        console.log("STUDENTS & SUBMISSIONS: ", submissions);
      } else {
        console.error(
          "Failed to fetch student submissions:",
          response.statusText
        );
      }
    } catch (error) {
      console.error("Error fetching student submissions:", error);
    }
  };

  const handleStudentSelect = (studentId: string) => {
    setSelectedStudent(studentId);
  };

  const handleGradeUpdate = async (grade: any, feedback: any) => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/assignment/submission/${selectedStudent}/update`,
        {
          method: "POST",
          body: JSON.stringify({
            grade: grade,
            feedback: feedback,
          }),
        }
      );
      if (response.ok) {
        console.log("Grade and feedback updated successfully");
      } else {
        console.error(
          "Failed to update grade and feedback:",
          response.statusText
        );
      }
    } catch (error) {
      console.error("Error updating grade and feedback:", error);
    }
  };

  const handleExcelUpload = (file: File) => {
    // Handle uploading of Excel file
  };

  const handleExcelDownload = async () => {
    // Handle downloading of Excel file
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg px-32 py-16 justify-end">
        <CloseButton onClick={onClose} />
        {assignment && (
          <div>
            <h1 className="font-bold text-black text-2xl pb-8">
              {assignment.title}
            </h1>
            <p className="font-bold text-black text-xl pb-8">
              Due Date: {assignment.due_date}
            </p>
            <p className="font-bold text-black text-l pb-8">
              Description: {assignment.description}
            </p>
          </div>
        )}

        <select onChange={(e) => handleStudentSelect(e.target.value)}>
          <option value="">Select Student</option>
          {submissions.map((submission) => (
            <option key={submission.userid} value={submission.userid}>
              {submission.userid}
            </option>
          ))}
        </select>

        {selectedStudent && (
          <div>
            {/* Display student's submitted file and allow download */}
            <p>Student's File: {/* Display student's file here */}</p>
            <button
              onClick={/* Download student's file */}
              className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Download File
            </button>

            <label htmlFor="grade" className="font-bold text-black text-l pb-8">
              Grade:
            </label>
            <input
              type="text"
              id="grade"
              value={grade}
              onChange={(e) => setGrade(e.target.value)}
            />

            <label
              htmlFor="feedback"
              className="font-bold text-black text-l pb-8"
            >
              Feedback:
            </label>
            <textarea
              id="feedback"
              value={feedback}
              onChange={(e) => setFeedback(e.target.value)}
            ></textarea>

            {grade && feedback && (
              <button
                onClick={handleGradeUpdate(grade, feedback)}
                className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                Update Grade & Feedback
              </button>
            )}
          </div>
        )}

        <label htmlFor="excelFile" className="font-bold text-black text-l pb-8">
          Upload Excel File:
        </label>
        <input
          type="file"
          id="excelFile"
          onChange={(e) => handleExcelUpload(e.target.files[0])}
        />

        <button
          onClick={handleExcelDownload}
          className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Download Excel
        </button>
      </div>
    </div>
  );
};

export default TeacherViewAssignment;
