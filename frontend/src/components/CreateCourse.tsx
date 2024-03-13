"use client";

import React, { useState } from "react";
import {AddIcon} from "@sanity/icons";

interface props {
  onClose: () => void;
  onCourseCreate: (courseData: any) => void;
}

const CreateCourse: React.FC<props> = (props: props) => {
  const [courseData, setCourseData] = useState({
    id: "",
    title: "",
    professor: "",
    location: "",
  });

  const postNewCourse = async (courseData: any) => {
    try {
      const res: Response = await fetch("/v1/course/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(courseData),
      })
      if (res.ok) {
        const newCourse = await res.json();
        newCourse.name = courseData.title;
        newCourse.id = courseData.id;
        newCourse.teachers.push(courseData.professor);
        newCourse.archived = false;
      } else {
        console.error("Failed to create course:", res.statusText);
      }
    } catch (error) {
      console.error("Error creating course:", error);
    }
  };

  const handleChange = (e: { target: { name: any; value: any } }) => {
    const { name, value } = e.target;
    setCourseData({
      ...courseData,
      [name]: value,
    });
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const idNum = Date.now().toString();
    setCourseData({
      ...courseData,
      id: idNum,
    });
    props.onCourseCreate({ ...courseData, id: idNum });
    postNewCourse(courseData);
    props.onClose();
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg px-32 py-16 justify-end">
        <button
          className="absolute top-0 right-0 m-2 text-black text-lg font-bold cursor-pointer"
          onClick={props.onClose}
        >
          <AddIcon />
        </button>
        <form className="justify-end" onSubmit={handleSubmit}>
          <h1 className="font-bold text-black text-2xl pb-8">
            Create New Course
          </h1>
          <div className="mb-2">
            <label
              htmlFor="title"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Course Name:
            </label>
            <input
              type="text"
              id="title"
              name="title"
              value={courseData.title}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            />
          </div>
          <div className="mb-2">
            <label
              htmlFor="teacher"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Professor:
            </label>
            <input
              type="text"
              id="teacher"
              name="professor"
              value={courseData.professor}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            />
          </div>
          <div className="mb-2">
            <label
              htmlFor="location"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Location:
            </label>
            <input
              type="text"
              id="location"
              name="location"
              value={courseData.location}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            />
          </div>
          <button
            type="submit"
            className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Create
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreateCourse;
