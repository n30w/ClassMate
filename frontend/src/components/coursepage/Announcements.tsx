"use client";

import React, { useState, useEffect } from "react";
import CreateAnnouncement from "./CreateAnnouncement";
import AddButton from "@/components/buttons/AddButton";
import AnnouncementDisplay from "./AnnouncementDisplay";
import axios from "axios";
import { revalidatePath } from "next/cache";

interface props {
  courseId: string;
}

const Announcements: React.FC<props> = (props: props) => {
  const [isCreatingAnnouncement, setIsCreatingAnnouncement] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  const [token, setToken] = useState("");

  async function fetchAnnouncements() {
    const response = await fetch(
      `http://localhost:6789/v1/course/announcement/read/${props.courseId}`
    );
    const { announcements } = await response.json();
    setAnnouncements(announcements);
  }

  const [announced, setAnnouncements] = useState<any>();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }

    const permissions = localStorage.getItem("permissions");
    if (permissions === "1") {
      setIsTeacher(true);
    }

    fetchAnnouncements();
  }, []);

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
      <AnnouncementDisplay announcements={announced} />

      {isCreatingAnnouncement && (
        <CreateAnnouncement
          onClose={() => {
            setIsCreatingAnnouncement(false);
          }}
          token={token}
          params={{ id: props.courseId }}
        />
      )}
    </div>
  );
};

export default Announcements;
