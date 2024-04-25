"use client";

import React, { useState, useEffect } from "react";
import CreateAnnouncement from "./CreateAnnouncement";
import AddButton from "@/components/buttons/AddButton";
import AnnouncementDisplay from "../announcements/AnnouncementDisplay";
import { Announcement } from "@/lib/types";

interface props {
  entries: Announcement[];
}

const Announcements: React.FC<props> = (props: props) => {
  const [announcements, setAnnouncements] = useState<Announcement[]>([]);
  const [isCreatingAnnouncement, setIsCreatingAnnouncement] = useState(false);

  const handleCreateAnnouncement = (announcementData: any) => {
    setAnnouncements([...announcements, announcementData]);
  };

  useEffect(() => {
    fetchAnnouncements();
  }, []);

  const fetchAnnouncements = async () => {
    try {
      const response = await fetch("/api/announcements");
      if (response.ok) {
        const data = await response.json();
        setAnnouncements(data);
      } else {
        console.error("Failed to fetch announcements:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching announcements:", error);
    }
  };

  return (
    <div className="w-full">
      <div className="flex justify-between border-b-2 border-white mb-4 pb-4">
        <h1 className="text-white font-bold text-2xl">Announcements</h1>
        <AddButton
          onClick={() => {
            setIsCreatingAnnouncement(true);
          }}
        />
      </div>
      <AnnouncementDisplay announcements={announcements} />
      {isCreatingAnnouncement && (
        <CreateAnnouncement
          onClose={() => {
            setIsCreatingAnnouncement(false);
          }}
          onAnnouncementCreate={handleCreateAnnouncement}
        />
      )}
    </div>
  );
};

export default Announcements;
