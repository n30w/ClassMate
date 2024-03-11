"use client";

import React, { useState } from "react";
import CreateAnnouncement from "./CreateAnnouncement";

interface Announcements {
  id: string;
  title: string;
  date: string;
  description: string;
}

const Announcements: React.FC = () => {
  const [announcements, setAnnouncements] = useState<Announcements[]>([]);
  const [isCreatingAnnouncement, setIsCreatingAnnouncement] = useState(false);

  const handleCreateAnnouncement = (announcementData: any) => {
    setAnnouncements([...announcements, announcementData]);
  };

  const closeMakeAnnouncement = () => {
    setIsCreatingAnnouncement(false);
  };

  const handleMakeAnnouncement = () => {
    setIsCreatingAnnouncement(true);
  };

  const announcementsDisplay =
    announcements.length > 0 ? (
      announcements.map((announcement, key) => (
        <div key={key}>
          <h2 className="text-white text-2xl mb-1">{announcement.title}</h2>
          <h3 className="text-white text-sm mb-2">{announcement.date}</h3>
          <p className="text-white text-sm font-light border-b-2 border-white mb-4 pb-4">
            {announcement.description}
          </p>
        </div>
      ))
    ) : (
      <p className="text-white">No announcements yet.</p>
    );

  return (
    <div className="w-full">
      <div className="flex justify-between border-b-2 border-white mb-4 pb-4">
        <h1 className="text-white font-bold text-2xl">Announcements</h1>
        <button
          className="rounded-full bg-white text-black text-sm font-light py-1 px-2 mt-2 flex items-center justify-center"
          onClick={handleMakeAnnouncement}
        >
          +
        </button>
      </div>
      {announcementsDisplay}
      {isCreatingAnnouncement && (
        <CreateAnnouncement
          onClose={closeMakeAnnouncement}
          onCourseCreate={handleCreateAnnouncement}
        />
      )}
    </div>
  );
};

export default Announcements;
