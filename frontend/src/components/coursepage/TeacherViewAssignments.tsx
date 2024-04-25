"use client";

import React, { useEffect, useState } from "react";
import CreateAssignment from "./CreateAssignment";
import AddButton from "@/components/buttons/AddButton";
import { Assignment, User } from "@/lib/types";

interface props {
  onGradeAssignment: (gradeData: any) => void;
}

const TeacherViewAssignment: React.FC<props> = (props: props) => {
  const [selectedAssignment, setSelectedAssignment] = useState("");
  const [selectedStudent, setSelectedStudent] = useState("");
  const [assignments, setAssignments] = useState<Assignment[]>([]);
  const [users, setUsers] = useState<User[]>([]);
  const [isCreatingAssignment, setIsCreatingAssignment] = useState(false);

  useEffect(() => {
    fetchAssignments();
    fetchUsers();
  }, []);

  const [gradeData, setGradeData] = useState({
    id: "",
    grade: "",
    feedback: "",
  });

  const fetchAssignments = async () => {
    try {
      const response = await fetch("/v1/assignment");
      if (response.ok) {
        const data = await response.json();
        setAssignments(data);
      } else {
        console.error("Failed to fetch assignments:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching assignments:", error);
    }
  };

  const fetchUsers = async () => {
    try {
      const response = await fetch("/v1/users");
      if (response.ok) {
        const data = await response.json();
        setUsers(data);
      } else {
        console.error("Failed to fetch users:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching users:", error);
    }
  };

  const postNewGrade = async (gradeData: any) => {
    try {
      const res: Response = await fetch("/v1/assignment/grade", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(gradeData),
      });
      if (res.ok) {
        const newGrade = await res.json();
        newGrade.grade = gradeData.grade;
        newGrade.id = gradeData.id;
        newGrade.feedback.push(gradeData.feedback);
        newGrade.archived = false;
      } else {
        console.error("Failed to grade assignment:", res.statusText);
      }
    } catch (error) {
      console.error("Error grading assignment:", error);
    }
  };

  const handleChange = (e: { target: { name: any; value: any } }) => {
    const { name, value } = e.target;
    setGradeData({
      ...gradeData,
      [name]: value,
    });
  };

  const handleSubmit = async (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const idNum = Date.now().toString();
    const gradeDataWithId = { ...gradeData, id: idNum };

    try {
      const response = await fetch("/v1/assignment/grade", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(gradeDataWithId),
      });

      if (response.ok) {
        console.log("Grade submitted successfully.");
      } else {
        console.error("Failed to submit grade:", response.statusText);
      }
    } catch (error) {
      console.error("Error submitting grade:", error);
    }

    setGradeData({
      id: "",
      grade: "",
      feedback: "",
    });
  };

  const handleCreateAssignment = (assignmentData: any) => {
    setAssignments([...assignments, assignmentData]);
  };

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedAssignment(event.target.value);
  };

  const handleSelectStudentChange = (
    event: React.ChangeEvent<HTMLSelectElement>
  ) => {
    setSelectedStudent(event.target.value);
  };

  return (
    <div className="w-full">
      <div className="flex justify-between border-b-2 border-white mb-4 pb-4">
        <h1 className="text-white font-bold text-2xl">Assignments</h1>
        <AddButton
          onClick={() => {
            setIsCreatingAssignment(true);
          }}
        />
      </div>
      <select value={selectedAssignment} onChange={handleSelectChange}>
        <option value="">Choose an assignment</option>
        {assignments.map((assignment, index) => (
          <option key={index} value={index}>
            {assignment.title}
          </option>
        ))}
      </select>
      <select value={selectedStudent} onChange={handleSelectStudentChange}>
        <option value="">Choose a student</option>
        {users.map((user, index) => (
          <option key={index} value={index}>
            {user.fullname}
          </option>
        ))}
      </select>
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
      <h2>Grade Assignment:</h2>
      <form className="justify-end" onSubmit={handleSubmit}>
        <h1 className="font-bold text-black text-2xl pb-8">Grade Assignment</h1>
        <div className="mb-2">
          <label
            htmlFor="grade"
            className="block text-lg font-medium text-gray-700 py-2"
          >
            Grade:
          </label>
          <input
            type="range"
            id="grade"
            name="grade"
            min="0"
            max="100"
            step="1"
            value={gradeData.grade}
            onChange={handleChange}
            className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            required
          />
        </div>
        <div className="mb-2">
          <label
            htmlFor="feedback"
            className="block text-lg font-medium text-gray-700 py-2"
          >
            Feedback:
          </label>
          <input
            type="text"
            id="feedback"
            name="feedback"
            value={gradeData.feedback}
            onChange={handleChange}
            className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
          />
        </div>
        <button
          type="submit"
          className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Save
        </button>
      </form>
      {isCreatingAssignment && (
        <CreateAssignment
          onClose={() => {
            setIsCreatingAssignment(false);
          }}
          onCourseCreate={handleCreateAssignment}
        />
      )}
    </div>
  );
};

export default TeacherViewAssignment;
