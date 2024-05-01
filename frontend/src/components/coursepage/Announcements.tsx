"use client";

import React, { useState, useEffect } from "react";
import CreateAnnouncement from "./CreateAnnouncement";
import AddButton from "@/components/buttons/AddButton";
import AnnouncementDisplay from "./AnnouncementDisplay";
import { Discussion } from "@/lib/types";

interface props {
  entries: Discussion[];
  courseId: string;
}

const Announcements: React.FC<props> = (props: props) => {
  const [announcements, setAnnouncements] = useState<Discussion[]>([]);
  const [isCreatingAnnouncement, setIsCreatingAnnouncement] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  const [token, setToken] = useState("");

  const handleCreateAnnouncement = (announcementData: any) => {
    setAnnouncements([...announcements, announcementData]);
  };

  useEffect(() => {
    const token = localStorage.getItem("token");
    setToken(token);
    const permissions = localStorage.getItem("permissions");
    if (permissions === "1") {
      setIsTeacher(true);
    }
    fetchAnnouncements(props.courseId);
  }, []);

  const fetchAnnouncements = async (url: string) => {
    try {
      const response = await fetch(
        "http://localhost:6789/v1/api/announcements/read",
        {
          method: "POST",
          body: JSON.stringify({
            courseid: url,
          }),
        }
      );
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
        {isTeacher && (
          <AddButton
            onClick={() => {
              setIsCreatingAnnouncement(true);
            }}
          />
        )}
      </div>
      <AnnouncementDisplay announcements={announcements} />
      {isCreatingAnnouncement && (
        <CreateAnnouncement
          onClose={() => {
            setIsCreatingAnnouncement(false);
          }}
          onAnnouncementCreate={handleCreateAnnouncement}
          courseId={props.courseId}
          token={token}
        />
      )}
    </div>
  );
};

export default Announcements;
