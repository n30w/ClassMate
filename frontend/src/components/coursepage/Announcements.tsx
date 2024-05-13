"use client";

import React, { useState, useEffect } from "react";
import CreateAnnouncement from "./CreateAnnouncement";
import AddButton from "@/components/buttons/AddButton";
import AnnouncementDisplay from "./AnnouncementDisplay";
import { useRouter, usePathname } from "next/navigation";
import { Announcement } from "@/lib/types";

interface props {
  courseId: string;
}

const Announcements: React.FC<props> = ({ courseId }: props) => {
  const router = useRouter();
  const pathName = usePathname();

  // Function is a variation of: https://www.joshwcomeau.com/nextjs/refreshing-server-side-props/
  // and https://nextjs.org/docs/app/api-reference/functions/use-pathname
  // and https://github.com/vercel/next.js/discussions/62146
  const refreshData = () => {
    router.push(pathName);
    window.location.reload();
  };

  const [isCreatingAnnouncement, setIsCreatingAnnouncement] = useState(false);
  const [isTeacher, setIsTeacher] = useState(false);
  const [token, setToken] = useState("");

  async function fetchAnnouncements() {
    const response = await fetch(
      `http://localhost:6789/v1/course/${courseId}/announcement/read`
    );
    const { announcements } = await response.json();
    setAnnouncements(announcements);
  }

  const [announced, setAnnouncements] = useState<Announcement[]>([]);

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
      {isTeacher && (
        <AddButton
          fullWidth={true}
          text="Create Announcement"
          onClick={() => {
            setIsCreatingAnnouncement(true);
          }}
        />
      )}

      <AnnouncementDisplay announcements={announced} />

      {isCreatingAnnouncement && (
        <CreateAnnouncement
          onClose={() => {
            setIsCreatingAnnouncement(false);
            refreshData();
          }}
          token={token}
          params={{ id: courseId }}
        />
      )}
    </div>
  );
};

export default Announcements;
