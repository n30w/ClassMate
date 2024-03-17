"use client";

import React, { useState } from "react";
import CloseButton from "@/components/buttons/CloseButton";

interface props {
  onClose: () => void;
  onAnnouncementCreate: (announcementData: any) => void;
}

const CreateAnnouncement: React.FC<props> = (props: props) => {
  const currentDate = new Date();

  const formattedDate = `${currentDate
    .toLocaleDateString("en-US", {
      month: "2-digit",
      day: "2-digit",
      year: "numeric",
    })
    .replace(/\//g, "-")} ${currentDate.toLocaleTimeString("en-US", {
    hour: "2-digit",
    minute: "2-digit",
  })}`;

  const [announcementData, setAnnouncementData] = useState({
    id: "",
    title: "",
    date: formattedDate,
    description: "",
  });

  const postNewAnnouncement = async (announcementData: any) => {
    try {
      const res: Response = await fetch("/v1/course/announcement/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(announcementData),
      });
      if (res.ok) {
        const newAnnouncement = await res.json();
        newAnnouncement.name = announcementData.title;
        newAnnouncement.id = announcementData.id;
        newAnnouncement.description = announcementData.description;
        newAnnouncement.date = announcementData.date;
      } else {
        console.error("Failed to create announcement:", res.statusText);
      }
    } catch (error) {
      console.error("Error creating announcement:", error);
    }
  };

  const handleChange = (e: { target: { name: any; value: any } }) => {
    const { name, value } = e.target;
    setAnnouncementData({
      ...announcementData,
      [name]: value,
    });
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const idNum = Date.now().toString();
    setAnnouncementData({
      ...announcementData,
      id: idNum,
    });
    props.onAnnouncementCreate({ ...announcementData, id: idNum });
    postNewAnnouncement(announcementData);
    props.onClose();
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg px-32 py-16 justify-end">
        <CloseButton onClick={props.onClose} />
        <form className="justify-end" onSubmit={handleSubmit}>
          <h1 className="font-bold text-black text-2xl pb-8">
            Create Announcement
          </h1>
          <div className="mb-2">
            <label
              htmlFor="title"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Announcement Title:
            </label>
            <input
              type="text"
              id="title"
              name="title"
              value={announcementData.title}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            />
          </div>
          <div className="mb-2">
            <label
              htmlFor="description"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Description:
            </label>
            <input
              type="text"
              id="description"
              name="description"
              value={announcementData.description}
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

export default CreateAnnouncement;
