"use client";

import React, { useState, useEffect } from "react";
import AddButton from "@/components/buttons/AddButton";
import { Assignment } from "@/lib/types";

interface props {
  entries: Assignment[];
}

const Assignments: React.FC<props> = (props: props) => {
  const [selectedAssignment, setSelectedAssignment] = useState("");
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  const [assignments, setAssignments] = useState<Assignment[]>([]);

  useEffect(() => {
    fetchAssignments();
  }, []);

  const fetchAssignments = async () => {
    try {
      const response = await fetch("/v1/assignment/read");
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

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedAssignment(event.target.value);
  };

  const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    const files = Array.from(event.dataTransfer.files);
    setUploadedFiles(files);
  };

  const handleFileInputChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    if (event.target.files) {
      const files = Array.from(event.target.files);
      setUploadedFiles(files);
    }
  };

  const handleFileRemove = (index: number) => {
    const newFiles = [...uploadedFiles];
    newFiles.splice(index, 1);
    setUploadedFiles(newFiles);
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

  const readFileAsBase64 = (file: File): Promise<string> => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onload = () => {
        const base64String = reader.result as string;
        // Extract the base64 content from the data URL
        const base64Content = base64String.split(",")[1];
        resolve(base64Content);
      };
      reader.onerror = (error) => reject(error);
      reader.readAsDataURL(file);
    });
  };

  return (
    <div className="w-full">
      <select value={selectedAssignment} onChange={handleSelectChange}>
        <option value="">Choose an assignment</option>
        {assignments.map((assignment, index) => (
          <option key={index} value={index}>
            {assignment.title}
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
      <h2>File Upload</h2>
      <div
        onDrop={handleDrop}
        onDragOver={(event) => event.preventDefault()}
        style={{
          border: "2px dashed #aaa",
          borderRadius: "5px",
          padding: "20px",
          marginTop: "20px",
          width: "550px",
        }}
      >
        <p className="text-white text-l font-bold">File Upload</p>
        <input
          type="file"
          id="file"
          onChange={handleFileInputChange}
          multiple
          style={{ display: "none" }}
        />
        <button
          className="rounded-full bg-white text-black text-sm font-light h-8 p-2 mt-8 flex items-center justify-center"
          onClick={() => postSubmission(uploadedFiles)}
        >
          Upload Files
        </button>
      </div>
      <div>
        <h2 className="text-white text-l mt-8">Uploaded Files:</h2>
        <ul>
          {uploadedFiles.map((file, index) => (
            <li className="text-white text-l mt-4" key={index}>
              {file.name} - {file.size} bytes
              <button
                className="rounded-full bg-white text-black text-sm font-light h-8 p-2 mt-2 flex items-center justify-center"
                onClick={() => handleFileRemove(index)}
              >
                Remove
              </button>
            </li>
          ))}
        </ul>
      </div>
      <button
        className="rounded-full bg-white text-black text-sm font-light h-8 p-2 mt-8 flex items-center justify-center"
        // onClick={handleFileUpload}
      >
        Submit
      </button>
    </div>
  );
};

export default Assignments;
