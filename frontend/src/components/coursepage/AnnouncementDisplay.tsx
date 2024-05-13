import { Announcement } from "@/lib/types";
import InfoBadge from "@/components/badge/InfoBadge";
import formattedDate from "@/lib/helpers/formattedDate";
import { useState, useEffect } from "react";

interface props {
  announcements: Announcement[];
}

const AnnouncementDisplay: React.FC<props> = ({ announcements }: props) => {
  const [isTeacher, setIsTeacher] = useState(false);

  useEffect(() => {
    const permissions = localStorage.getItem("permissions");

    if (permissions === "1") {
      setIsTeacher(true);
    }
  });

  const handleDeleteAnnouncement = async (announcementId: string) => {
    try {
      const response = await fetch(
        `http://localhost:6789/v1/course/announcement/${announcementId}/delete`,
        {
          method: "DELETE",
          headers: {
            "Access-Control-Request-Method": "POST",
          },
        }
      );
      if (response.ok) {
        window.location.reload();
        console.log("Announcement deleted successfully");
      } else {
        console.error("Failed to delete announcement");
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

  return (
    <div className="w-full h-full grid grid-cols-1 grid-rows-3 border-slate-200 border-opacity-10 border-2">
      {announcements ? (
        announcements.map((announcement: Announcement, i: number) => (
          <div className="announcement-item hover:bg-gray-700" key={i}>
            <div className="flex justify-between">
              <h2 className="text-white text-3xl mb-1 font-bold">
                {announcement.title}
              </h2>
              {isTeacher && (
                <button
                  onClick={() => handleDeleteAnnouncement(announcement.id)}
                  className="text-white bg-red-500 hover:bg-red-700 active:bg-red-900 font-bold py-2 px-4 rounded mt-2 w-32"
                >
                  Delete
                </button>
              )}
            </div>
            <InfoBadge
              text={formattedDate(announcement.date).toLocaleUpperCase()}
              colorClass={"bg-blue-500"}
            />
            <p className="text-white text-lg font-light">
              {announcement.description}
            </p>
          </div>
        ))
      ) : (
        <>
          <div className={"announcement-item"}>
            <p className="text-hint p-2">New announcements will appear here.</p>
          </div>
          <div className={"announcement-item"}></div>
          <div className={"announcement-item"}></div>
        </>
      )}
    </div>
  );
};

export default AnnouncementDisplay;
