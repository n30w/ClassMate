"use client";

import React, { useState, useEffect } from "react";
import CreateAnnouncement from "./CreateAnnouncement";
import AddButton from "@/components/buttons/AddButton";

interface Announcements {
  id: string;
  title: string;
  date: string;
  description: string;
}

const AnnouncementDisplay: React.FC<{ announcements: any[] }> = ({
  announcements,
}) => {
  return announcements.length > 0 ? (
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
};

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

  const fetchAnnouncements = async () => {
    try {
      const res: Response = await fetch("/v1/course/announcement/read", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (res.ok) {
        const announcement = await res.json();
        return announcement;
      } else {
        console.error("Failed to fetch announcement:", res.statusText);
        return [];
      }
    } catch (error) {
      console.error("Error fetching announcement:", error);
      return [];
    }
  };

  useEffect(() => {
    const getAnnouncement = async () => {
      const fetchedAnnouncements = await fetchAnnouncements();
      setAnnouncements(fetchedAnnouncements);
    };

    getAnnouncement();
  }, []);

  return (
    <div className="w-full">
      <div className="flex justify-between border-b-2 border-white mb-4 pb-4">
        <h1 className="text-white font-bold text-2xl">Announcements</h1>
        <AddButton onClick={handleMakeAnnouncement} />
      </div>
      <AnnouncementDisplay announcements={announcements} />
      {isCreatingAnnouncement && (
        <CreateAnnouncement
          onClose={closeMakeAnnouncement}
          onAnnouncementCreate={handleCreateAnnouncement}
        />
      )}
    </div>
  );
};

export default Announcements;
