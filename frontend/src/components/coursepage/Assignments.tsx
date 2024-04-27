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

  const handleFileUpload = async () => {
    if (uploadedFiles.length === 0) {
      console.error("No files selected for upload.");
      return;
    }

    const formData = new FormData();
    uploadedFiles.forEach((file) => {
      formData.append("files", file);
    });

    try {
      const response = await fetch("/v1/upload", {
        method: "POST",
        body: formData,
      });

      if (response.ok) {
        console.log("Files uploaded successfully.");
      } else {
        console.error("Failed to upload files:", response.statusText);
      }
    } catch (error) {
      console.error("Error uploading files:", error);
    }
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
          onChange={handleFileInputChange}
          multiple
          style={{ display: "none" }}
        />
        <button
          className="rounded-full bg-white text-black text-sm font-light h-8 p-2 mt-8 flex items-center justify-center"
          onClick={() =>
            (
              document.querySelector(
                'input[type="file"]'
              ) as HTMLInputElement | null
            )?.click()
          }
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
        onClick={handleFileUpload}
      >
        Submit
      </button>
    </div>
  );
};

export default Assignments;
