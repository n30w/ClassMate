import { Announcement } from "@/lib/types";
import DateBadge from "@/components/badge/DateBadge";

interface props {
  announcements: Announcement[];
}

const AnnouncementDisplay: React.FC<props> = ({ announcements }: props) => {
  {
    /* Checks if announced is null. This won't work by checking if the array is greater than zero, because announced is a promise. */
  }
  return (
    <div className="w-full h-full grid grid-cols-1 grid-rows-3 border-slate-200 border-opacity-10 border-2">
      {announcements ? (
        announcements.map((announcement: Announcement, key: number) => (
          <div className="announcement-item hover:bg-gray-700" key={key}>
            <h2 className="text-white text-3xl mb-1 font-bold">
              {announcement.title}
            </h2>
            <DateBadge date={announcement.date} colorClass={"rose-700"} />
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
